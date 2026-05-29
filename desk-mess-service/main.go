package main

import (
	"encoding/json"
	"log"
	"math"
	"net/http"
)

type HealthResponse struct {
	Status  string `json:"status"`
	Service string `json:"service"`
}

type DeskStatusResponse struct {
	DeskID         string  `json:"desk_id"`
	CoffeeCups     int     `json:"coffee_cups"`
	OpenTabs       int     `json:"open_tabs"`
	StickyNotes    int     `json:"sticky_notes"`
	LooseCables    int     `json:"loose_cables"`
	UnreadMessages int     `json:"unread_messages"`
	EntropyScore   float64 `json:"entropy_score"`
	Status         string  `json:"status"`
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	response := HealthResponse{
		Status:  "ok",
		Service: "desk-mess-service",
	}

	writeJSON(w, response)
}

func deskStatusHandler(w http.ResponseWriter, r *http.Request) {
	coffeeCups := 2
	openTabs := 43
	stickyNotes := 8
	looseCables := 5
	unreadMessages := 12

	entropyScore := calculateEntropyScore(
		coffeeCups,
		openTabs,
		stickyNotes,
		looseCables,
		unreadMessages,
	)

	response := DeskStatusResponse{
		DeskID:         "home-office-desk",
		CoffeeCups:     coffeeCups,
		OpenTabs:       openTabs,
		StickyNotes:    stickyNotes,
		LooseCables:    looseCables,
		UnreadMessages: unreadMessages,
		EntropyScore:   entropyScore,
		Status:         determineDeskStatus(entropyScore),
	}

	writeJSON(w, response)
}

func calculateEntropyScore(
	coffeeCups int,
	openTabs int,
	stickyNotes int,
	looseCables int,
	unreadMessages int,
) float64 {
	score := float64(coffeeCups)*0.10 +
		float64(openTabs)*0.01 +
		float64(stickyNotes)*0.04 +
		float64(looseCables)*0.05 +
		float64(unreadMessages)*0.02

	score = math.Min(score, 1.0)

	return math.Round(score*100) / 100
}

func determineDeskStatus(entropyScore float64) string {
	if entropyScore <= 0.30 {
		return "calm_workspace"
	}

	if entropyScore <= 0.60 {
		return "creative_mess"
	}

	if entropyScore <= 0.85 {
		return "focus_under_attack"
	}

	return "entropy_has_won"
}

func writeJSON(w http.ResponseWriter, response any) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/desk/status", deskStatusHandler)

	log.Println("Desk Mess Service is running on port 9090")

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal(err)
	}
}
