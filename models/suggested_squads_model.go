package models

type SuggestedSquadsModel struct {
	SquadID      int64  `json:"squad_id"`
	Name         string `json:"name"`
	SquadProfile string `json:"squad_profile"`
	Level        string `json:"level"`
	Engagement   string `json:"engagement"`
	Category     string `json:"category"`
	IsPrivate    int64  `json:"is_private"`
	IsPaid       int64  `json:"is_paid"`
	Admin        int64  `json:"admin"`
	Members      int64  `json:"members"`
}
