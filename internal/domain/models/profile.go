package models

type Profile struct {
	ID          int64
	Email       string
	AvatarImage string
	Name        string
	Description string
	Link        string
	Boards      []int64
	SavedIdeas  []int64
}