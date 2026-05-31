package handlers

import "net/http"

type HealthResponse struct {
	Status  string `json:"status"`
	Service string `json:"service"`
	Role    string `json:"role"`
}

func SidecarHealthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		WriteJSON(w, http.StatusMethodNotAllowed, map[string]string{
			"error":   "method_not_allowed",
			"message": "Only GET method is allowed",
		})
		return
	}

	response := HealthResponse{
		Status:  "ok",
		Service: "entropy-sidecar",
		Role:    "operational_companion",
	}

	WriteJSON(w, http.StatusOK, response)
}