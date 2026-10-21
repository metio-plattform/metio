package servers

import (
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
	"github.com/metio-plattform/metio/internal/services"
)

// SearchModpacks serves GET /api/modpacks/search. It responds 200 with the
// Modrinth results as a JSON array; the service never returns an error, so an
// unreachable Modrinth yields 200 with empty or stale results rather than a
// failure the browser could not recover from.
func SearchModpacks(w http.ResponseWriter, r *http.Request) {
	query, mcVersion := searchParams(r.URL)

	packs := []services.ModrinthPack{}
	if query != "" {
		packs = SearchModrinthPacks(r.Context(), query, mcVersion)
	}

	writeJSONResponse(w, http.StatusOK, packs)
}

// GetModpack serves GET /api/modpacks/{id}. It answers 200 with the project
// details so a server's installed pack can be shown by name rather than by its
// opaque project id. Like search, it never fails the browser: an unreachable
// Modrinth or an unknown project degrades to the id standing in for the name.
func GetModpack(w http.ResponseWriter, r *http.Request) {
	modpackID := mux.Vars(r)["id"]

	pack, ok := GetModrinthPack(r.Context(), modpackID)
	if !ok {
		pack = modrinthPackFallback(modpackID)
	}

	writeJSONResponse(w, http.StatusOK, pack)
}

// modrinthPackFallback is the degraded project description served when the
// catalogue cannot resolve a pack: the stored id doubles as its label.
func modrinthPackFallback(modpackID string) services.ModrinthPack {
	return services.ModrinthPack{
		ID:           modpackID,
		Slug:         modpackID,
		Name:         modpackID,
		GameVersions: []string{},
	}
}

// ListModpackVersions serves GET /api/modpacks/{id}/versions. It responds 200
// with the pack's versions as a JSON array, degraded to empty when Modrinth is
// unreachable.
func ListModpackVersions(w http.ResponseWriter, r *http.Request) {
	modpackID := mux.Vars(r)["id"]

	writeJSONResponse(w, http.StatusOK, ListModrinthPackVersions(r.Context(), modpackID))
}

// searchParams extracts the query and optional Minecraft version filter.
func searchParams(u *url.URL) (query, mcVersion string) {
	q := u.Query()
	return q.Get("query"), q.Get("mcVersion")
}
