package handlers

import (
	"desk-mess-service/models"
	"net/http"
)

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, http.StatusMethodNotAllowed, "method_not_allowed", "Only GET method is allowed")
		return
	}
	response := models.HealthResponse{
		Status:  "ok",
		Service: "desk-mess-service",
	}

	writeJSON(w, http.StatusOK, response)
}
