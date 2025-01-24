package models

type ChatData struct {
	ID    int64
	User1 ChatUserData
	User2 ChatUserData
}
type ChatUserData struct {
	UserId   int64
	Username string
	Avatar   string
}
