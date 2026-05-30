package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type HealthResponse struct {
	Status  string `json:"status"`
	Service string `json:"service"`
	Role    string `json:"role"`
}

func sidecarHealthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]string{
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

	writeJSON(w, http.StatusOK, response)
}

func createReverseProxy(target string) http.Handler {
	targetURL, err := url.Parse(target)
	if err != nil {
		log.Fatal(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	return proxy
}

func writeJSON(w http.ResponseWriter, statusCode int, response any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}

func main() {
	mux := http.NewServeMux()

	deskMessProxy := createReverseProxy("http://localhost:9090")

	mux.HandleFunc("/sidecar/health", sidecarHealthHandler)
	mux.Handle("/", deskMessProxy)

	log.Println("Entropy Sidecar is running on port 8080")
	log.Println("Proxying requests to Desk Mess Service on port 9090")

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}