package sqlite

import (
	"context"
	"fmt"
	"idea-store-auth/internal/domain/models"
	"log/slog"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Storage) SendMessage(ctx context.Context, message models.Message) (*emptypb.Empty, error) {
	slog.Info("storage started SendMessage")

	stmt, err := s.db.Prepare("INSERT INTO messages(sender_id, reciever_id, file_name, content, send_date) VALUES(?,?,?,?,?)")

	if err != nil {
		slog.Error("storage error SendMessage: " + err.Error())
		return nil, fmt.Errorf("storage error SendMessage: %v", err.Error())
	}

	_, err = stmt.ExecContext(ctx, message.SenderId, message.RecieverId, message.Filename, message.Text, message.CreationDate)

	if err != nil {
		slog.Error("storage error SendMessage: " + err.Error())
		return nil, fmt.Errorf("storage error SendMessage: %v", err.Error())
	}
	return nil, nil
}

func (s *Storage) GetMessages(ctx context.Context, senderId, recieverId int64) ([]*models.Message, error) {

	slog.Info("storage started GetMessages")

	stmt, err := s.db.Prepare("SELECT id, file_name, content, send_date FROM messages WHERE sender_id=? AND reciever_id = ?")

	if err != nil {
		slog.Error("storage error GetMessages: " + err.Error())
		return nil, fmt.Errorf("storage error GetMessages: %v", err.Error())
	}

	rows, err := stmt.QueryContext(ctx, senderId, recieverId)

	if err != nil {
		slog.Error("storage error GetMessages: " + err.Error())
		return nil, fmt.Errorf("storage error GetMessages: %v", err.Error())
	}
	var result []*models.Message
	for rows.Next() {
		message := models.Message{}
		err = rows.Scan(&message.ID, &message.Filename, &message.Text, &message.CreationDate)

		if err != nil {
			slog.Error("storage error GetMessages: " + err.Error())
			return nil, fmt.Errorf("storage error GetMessages: %v", err.Error())
		}
		message.RecieverId = recieverId
		message.SenderId = senderId
		result = append(result, &message)
	}

	return result, nil
}

func (s *Storage) CreateChat(ctx context.Context, firstId, secondId int64) (*emptypb.Empty, error) {
	return nil, nil
}

func (s *Storage) GetUsersChats(ctx context.Context, userId int64) ([]*models.ChatData, error) {
	return nil, nil
}

func (s *Storage) DeleteChat(ctx context.Context, chatId int64) (*emptypb.Empty, error) {
	return nil, nil
}
