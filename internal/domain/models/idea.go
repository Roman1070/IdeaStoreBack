package models

type Idea struct {
	ID          int64  `json:"idea_id"`
	Image       string `json:"image"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Link        string `json:"link"`
	Tags        string `json:"tags"`
	UserID      int64  `json:"user_id"`
}