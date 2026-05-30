package handlers

import "net/http"

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	writeError(w, http.StatusNotFound, "not_found", "The requested endpoint was not found")
}