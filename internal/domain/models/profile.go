package models

import profilesv1 "idea-store-auth/gen/go/profiles"

type Profile struct {
	ID          int64
	Email       string
	AvatarImage string
	Name        string
	Description string
	Link        string
	Boards      []int64
	SavedIdeas  []*profilesv1.IdeaBoardPair
}