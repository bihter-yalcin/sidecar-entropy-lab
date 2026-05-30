package handlers

import (
	"desk-mess-service/services"
	"net/http"
)

func DeskStatusHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {

		writeError(w, http.StatusMethodNotAllowed, "method_not_allowed", "Only GET method is allowed")
		return
	}
	response := services.GetDeskStatus()

	writeJSON(w, http.StatusOK, response)
}

func DeskItemsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, http.StatusMethodNotAllowed, "method_not_allowed", "Only GET method is allowed")
		return
	}
	items := services.GetDeskItems()

	writeJSON(w, http.StatusOK, items)
}
