package models

type ExpertListModel struct {
	ID           int64  `json:"id"`
	First_name   string `json:"Name"`
	Category_ids string `json:"category_ids"`
}