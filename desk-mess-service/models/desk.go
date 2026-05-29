package models

type DeskItem struct {
	Name       string  `json:"name"`
	Count      int     `json:"count"`
	MessWeight float64 `json:"mess_weight"`
}

type DeskStatusResponse struct {
	DeskID       string     `json:"desk_id"`
	Items        []DeskItem `json:"items"`
	EntropyScore float64    `json:"entropy_score"`
	Status       string     `json:"status"`
}
