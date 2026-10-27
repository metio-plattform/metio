// SPDX-License-Identifier: Apache-2.0

package servers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/metio-plattform/metio/internal/db"
)

func UpdateServerStatus(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	serverID := vars["id"]
	if serverID == "" {
		writeJSONError(w, "server id is required", http.StatusBadRequest)
		return
	}

	var status db.Status
	if err := json.NewDecoder(r.Body).Decode(&status); err != nil {
		writeJSONError(w, "invalid request body", http.StatusBadRequest)
		return
	}

	dbConn, _, err := GetDBConnection(ctx)
	if err != nil {
		log.Printf("Error creating db connection: %v", err)
		writeJSONError(w, "failed to connect to database", http.StatusInternalServerError)
		return
	}

	serverConfig, err := dbConn.GetServerConfig(ctx, serverID)
	if err != nil {
		writeJSONError(w, "server not found", http.StatusNotFound)
		return
	}

	if err := dbConn.UpdateStatus(ctx, serverConfig.Name, status); err != nil {
		log.Printf("Error updating status: %v", err)
		writeJSONError(w, "failed to update status", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
