package handlers

import (
	"desk-mess-service/services"
	"encoding/json"
	"net/http"
)

func DeskStatusHandler(w http.ResponseWriter, r *http.Request) {
	response := services.GetDeskStatus()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func DeskItemsHandler(w http.ResponseWriter, r *http.Request) {
	items := services.GetDeskItems()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}
