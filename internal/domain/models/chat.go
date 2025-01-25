package models

type ChatData struct {
	ID       int64
	FirstId  int64
	SecondId int64
}
type ChatUserData struct {
	UserId   int64
	Username string
	Avatar   string
}
