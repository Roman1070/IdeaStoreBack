package storage

import "errors"

var (
	ErrUserExists   = errors.New("User already exists")
	ErrUserNotFound = errors.New("User not found")
	ErrAppNotFound  = errors.New("App not found")
	ErrIdeaNotFound  = errors.New("Idea not found")
	ErrBoardNotFound  = errors.New("Board not found")
)
