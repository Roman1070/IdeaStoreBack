package postgre

import (
	"context"
	"fmt"
	profilesv1 "idea-store-auth/gen/go/profiles"
	"idea-store-auth/internal/domain/models"
	"idea-store-auth/internal/utils"
	"log/slog"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Storage) CheckChatExistance(ctx context.Context, firstId, secondId int64) (bool, error) {
	const query = `
		SELECT COUNT(*) 
		FROM chats 
		WHERE (first_id = $1 AND second_id = $2) 
		OR (first_id = $2 AND second_id = $1);
	`

	rowsCount := 0
	err := s.db.QueryRow(ctx, query, firstId, secondId).Scan(&rowsCount)
	if err != nil {
		slog.Error("storage error CheckChatExistance: " + err.Error())
		return false, fmt.Errorf("storage error CheckChatExistance: %v", err.Error())
	}

	return rowsCount > 0, nil
}

func (s *Storage) SendMessage(ctx context.Context, message models.Message) (int64, error) {
	slog.Info("storage started SendMessage")

	const query = `
		INSERT INTO messages(sender_id, reciever_id, file_name, content, send_date, sending_date_seconds,idea_id) 
		VALUES($1,$2,$3,$4,$5,$6,$7)
		RETURNING id;
	`

	dateInSeconds, err := utils.DateTimeToSecondsForDb(message.CreationDate)
	if err != nil {
		slog.Error("storage error SendMessage: " + err.Error())
		return emptyValue, fmt.Errorf("storage error SendMessage: %v", err.Error())
	}

	var lastInsertId int64
	err = s.db.QueryRow(ctx, query, message.SenderId, message.RecieverId, message.Filename,
		message.Text, message.CreationDate, dateInSeconds, message.IdeaId).Scan(&lastInsertId)

	if err != nil {
		slog.Error("storage SendMessage error: " + err.Error())
		return emptyValue, fmt.Errorf("storage SendMessage error: %v", err.Error())
	}

	return lastInsertId, nil
}

func (s *Storage) GetMessages(ctx context.Context, firstId, secondId int64) ([]*models.Message, error) {
	slog.Info("storage started GetMessages")

	const query = `
		SELECT id,sender_id, reciever_id, file_name, content, send_date, idea_id 
		FROM messages 
		WHERE (sender_id=$1 AND reciever_id = $2) 
		OR (sender_id=$2 AND reciever_id = $1) 
		ORDER BY sending_date_seconds;
	`

	rows, err := s.db.Query(ctx, query, firstId, secondId)
	if err != nil {
		slog.Error("storage error GetMessages: " + err.Error())
		return nil, fmt.Errorf("storage error GetMessages: %v", err.Error())
	}

	result := make([]*models.Message, 0, 20)
	for rows.Next() {
		message := models.Message{}
		err = rows.Scan(&message.ID, &message.SenderId, &message.RecieverId, &message.Filename, &message.Text, &message.CreationDate, &message.IdeaId)
		if err != nil {
			slog.Error("storage error GetMessages: " + err.Error())
			return nil, fmt.Errorf("storage error GetMessages: %v", err.Error())
		}

		result = append(result, &message)
	}

	return result, nil
}

func (s *Storage) CreateChat(ctx context.Context, firstId, secondId int64) (*emptypb.Empty, error) {
	slog.Info("storage started CreateChat")

	const query = `
		INSERT INTO chats(first_id,second_id) 
		VALUES($1,$2);
	`
	_, err := s.db.Exec(ctx, query, firstId, secondId)
	if err != nil {
		slog.Error("storage error CreateChat: " + err.Error())
		return nil, fmt.Errorf("storage error CreateChat: %v", err.Error())
	}

	return nil, nil
}

func (s *Storage) GetUsersChats(ctx context.Context, userId int64) ([]*models.ChatData, error) {
	slog.Info("storage started GetUsersChats, id = " + fmt.Sprint(userId))

	const query = `
		SELECT id, first_id, second_id 
		FROM chats 
		WHERE first_id = $1 
		OR second_id = $1;
	`

	rows, err := s.db.Query(ctx, query, userId)
	if err != nil {
		slog.Error("storage error GetUsersChats: " + err.Error())
		return nil, fmt.Errorf("storage error GetUsersChats: %v", err.Error())
	}

	var result []*models.ChatData
	for rows.Next() {
		var chat models.ChatData
		err = rows.Scan(&chat.ID, &chat.FirstData.UserId, &chat.SecondData.UserId)
		if err != nil {
			slog.Error("storage error GetUsersChats: " + err.Error())
			return nil, fmt.Errorf("storage error GetUsersChats: %v", err.Error())
		}

		if userId == chat.SecondData.UserId {
			profile, err := s.profilesClient.GetProfileLight(ctx, &profilesv1.GetProfileLightRequest{
				UserId: chat.FirstData.UserId,
			})
			if err != nil {
				slog.Error("storage error GetUsersChats: " + err.Error())
				return nil, fmt.Errorf("storage error GetUsersChats: %v", err.Error())
			}

			chat.FirstData.Avatar = profile.Avatar
			chat.FirstData.Username = profile.Name
		} else {
			profile, err := s.profilesClient.GetProfileLight(ctx, &profilesv1.GetProfileLightRequest{
				UserId: chat.SecondData.UserId,
			})
			if err != nil {
				slog.Error("storage error GetUsersChats: " + err.Error())
				return nil, fmt.Errorf("storage error GetUsersChats: %v", err.Error())
			}

			chat.SecondData.Avatar = profile.Avatar
			chat.SecondData.Username = profile.Name
		}

		result = append(result, &chat)
	}

	return result, nil
}

func (s *Storage) DeleteChat(ctx context.Context, chatId int64) (*emptypb.Empty, error) {
	return nil, nil
}
