package main

import (
	"desk-mess-service/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/health", handlers.HealthHandler)
	http.HandleFunc("/desk/status", handlers.DeskStatusHandler)
	http.HandleFunc("/desk/items", handlers.DeskItemsHandler)

	log.Println("Desk Mess Service is running on port 9090")

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal(err)
	}
}
