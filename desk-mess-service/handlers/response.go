package handlers

import (
	"desk-mess-service/models"
	"encoding/json"
	"net/http"
)

func writeJSON(w http.ResponseWriter, statusCode int, response any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}

func writeError(w http.ResponseWriter, statusCode int, errorCode string, message string) {
	response := models.ErrorResponse{
		Error:   errorCode,
		Message: message,
	}

	writeJSON(w, statusCode, response)
}
