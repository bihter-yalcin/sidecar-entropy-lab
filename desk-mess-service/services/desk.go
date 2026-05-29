package services

import (
	"desk-mess-service/models"
	"math"
)

func GetDeskItems() []models.DeskItem {
	return []models.DeskItem{
		{
			Name:       "coffee_cups",
			Count:      2,
			MessWeight: 0.10,
		},
		{
			Name:       "open_tabs",
			Count:      43,
			MessWeight: 0.01,
		},
		{
			Name:       "sticky_notes",
			Count:      8,
			MessWeight: 0.04,
		},
		{
			Name:       "loose_cables",
			Count:      5,
			MessWeight: 0.05,
		},
		{
			Name:       "unread_messages",
			Count:      12,
			MessWeight: 0.02,
		},
	}
}

func CalculateEntropyScore(items []models.DeskItem) float64 {
	var score float64

	for _, item := range items {
		score += float64(item.Count) * item.MessWeight
	}

	score = math.Min(score, 1.0)

	return math.Round(score*100) / 100
}

func DetermineDeskStatus(entropyScore float64) string {
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

func GetDeskStatus() models.DeskStatusResponse {
	items := GetDeskItems()
	entropyScore := CalculateEntropyScore(items)

	return models.DeskStatusResponse{
		DeskID:       "home-office-desk",
		Items:        items,
		EntropyScore: entropyScore,
		Status:       DetermineDeskStatus(entropyScore),
	}
}
