package main

import (
	"desk-mess-service/handlers"
	"log"
	"net/http"
)

func main() {
	mux:= http.NewServeMux()
	mux.HandleFunc("/health", handlers.HealthHandler)
	mux.HandleFunc("/desk/status", handlers.DeskStatusHandler)
	mux.HandleFunc("/desk/items", handlers.DeskItemsHandler)
	mux.HandleFunc("/", handlers.NotFoundHandler)
	log.Println("Desk Mess Service is running on port 9090")

	err := http.ListenAndServe(":9090", mux)
	if err != nil {
		log.Fatal(err)
	}
}
