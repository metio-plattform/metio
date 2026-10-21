package servers

import (
	"encoding/json"
	"testing"

	"github.com/metio-plattform/metio/internal/db"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func modpackConfigPtr(mc db.ModpackConfig) *db.ModpackConfig { return &mc }

func modpackConfigPPtr(mc db.ModpackConfig) **db.ModpackConfig {
	p := &mc
	return &p
}

// modpackRemovalPPtr is the "remove the pack" state crud.go expects: a present
// outer pointer whose inner pointer is nil.
func modpackRemovalPPtr() **db.ModpackConfig {
	var p *db.ModpackConfig
	return &p
}

func TestUpdateServerRequest_UnmarshalModpackTriState(t *testing.T) {
	tests := []struct {
		name     string
		body     string
		wantName string
		wantNil  bool // no Modpack pointer at all: leave unchanged
		wantNone bool // pointer to nil: remove the pack
	}{
		{
			name:     "absent leaves the pack unchanged",
			body:     `{"name":"renamed"}`,
			wantName: "renamed",
			wantNil:  true,
		},
		{
			name:     "null removes the pack",
			body:     `{"modpack":null,"minecraftVersion":"1.21.11"}`,
			wantNone: true,
		},
		{
			name: "object sets the pack",
			body: `{"modpack":{"platform":"modrinth","projectId":"abC123","versionId":"v1"}}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var req UpdateServerRequest
			require.NoError(t, json.Unmarshal([]byte(tt.body), &req))

			if tt.wantName != "" {
				require.NotNil(t, req.Name, "other fields must still decode through the custom unmarshaller")
				assert.Equal(t, tt.wantName, *req.Name)
			}

			switch {
			case tt.wantNil:
				assert.Nil(t, req.Modpack, "an absent field must not be treated as a change")
			case tt.wantNone:
				require.NotNil(t, req.Modpack, "a null field must be present")
				assert.Nil(t, *req.Modpack, "a null field must decode to an inner nil pointer")
			default:
				require.NotNil(t, req.Modpack)
				require.NotNil(t, *req.Modpack)
				assert.Equal(t, "modrinth", (*req.Modpack).Platform)
				assert.Equal(t, "abC123", (*req.Modpack).ProjectID)
				assert.Equal(t, "v1", (*req.Modpack).VersionID)
			}
		})
	}
}

func TestUpdateServerRequest_UnmarshalRejectsMalformedModpack(t *testing.T) {
	var req UpdateServerRequest
	err := json.Unmarshal([]byte(`{"modpack":{"platform":42}}`), &req)

	assert.Error(t, err)
}

func TestClassifyUpdate(t *testing.T) {
	strPtr := func(s string) *string { return &s }

	tests := []struct {
		name       string
		req        UpdateServerRequest
		wantUpdate UpdateType
	}{
		{
			name:       "no changes is in-place",
			req:        UpdateServerRequest{},
			wantUpdate: UpdateTypeInPlace,
		},
		{
			name: "machine type change is a resize",
			req: UpdateServerRequest{
				MachineType: strPtr("n2-standard-4"),
			},
			wantUpdate: UpdateTypeResize,
		},
		{
			name: "version change is a recreate",
			req: UpdateServerRequest{
				MinecraftVersion: strPtr("1.21.1"),
			},
			wantUpdate: UpdateTypeRecreate,
		},
		{
			name: "version change wins over machine type",
			req: UpdateServerRequest{
				MachineType:      strPtr("n2-standard-4"),
				MinecraftVersion: strPtr("1.21.1"),
			},
			wantUpdate: UpdateTypeRecreate,
		},
		{
			name: "setting a pack is a recreate",
			req: UpdateServerRequest{
				Modpack: modpackConfigPPtr(db.ModpackConfig{Platform: "modrinth", ProjectID: "abC123"}),
			},
			wantUpdate: UpdateTypeRecreate,
		},
		{
			name: "removing a pack is a recreate",
			req: UpdateServerRequest{
				Modpack: modpackRemovalPPtr(),
			},
			wantUpdate: UpdateTypeRecreate,
		},
		{
			name: "pack change wins over version",
			req: UpdateServerRequest{
				MinecraftVersion: strPtr("1.21.1"),
				Modpack:          modpackConfigPPtr(db.ModpackConfig{Platform: "modrinth", ProjectID: "abC123"}),
			},
			wantUpdate: UpdateTypeRecreate,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, int(tt.wantUpdate), classifyUpdate(tt.req, nil))
		})
	}
}
