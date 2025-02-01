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

func (s *Storage) SendMessage(ctx context.Context, message models.Message) (int64, error) {
	slog.Info("storage started SendMessage")
	dateInSeconds, err := utils.DateTimeToSecondsForDb(message.CreationDate)
	if err != nil {
		slog.Error("storage error SendMessage: " + err.Error())
		return emptyValue, fmt.Errorf("storage error SendMessage: %v", err.Error())
	}

	if message.CheckChatExistance {
		stmt, err := s.db.Prepare("SELECT COUNT(*) FROM chats WHERE (first_id = ? AND second_id = ?) OR (first_id = ? AND second_id = ?)")
		if err != nil {
			slog.Error("storage error SendMessage: " + err.Error())
			return emptyValue, fmt.Errorf("storage error SendMessage: %v", err.Error())
		}
		row := stmt.QueryRowContext(ctx, message.RecieverId, message.SenderId, message.SenderId, message.RecieverId)
		rowsCount := 0
		err = row.Scan(&rowsCount)

		if err != nil {
			slog.Error("storage error SendMessage: " + err.Error())
			return emptyValue, fmt.Errorf("storage error SendMessage: %v", err.Error())
		}

		if rowsCount == 0 {
			_, err = s.CreateChat(ctx, message.SenderId, message.RecieverId)
			if err != nil {
				slog.Error("storage error SendMessage: " + err.Error())
				return emptyValue, fmt.Errorf("storage error SendMessage: %v", err.Error())
			}
		}
	}

	stmt, err := s.db.Prepare(
		"INSERT INTO messages(sender_id, reciever_id, file_name, content, send_date, sending_date_seconds,idea_id) VALUES(?,?,?,?,?,?,?)")

	if err != nil {
		slog.Error("storage error SendMessage: " + err.Error())
		return emptyValue, fmt.Errorf("storage error SendMessage: %v", err.Error())
	}

	res, err := stmt.ExecContext(ctx, message.SenderId, message.RecieverId, message.Filename,
		message.Text, message.CreationDate, dateInSeconds, message.IdeaId)

	if err != nil {
		slog.Error("storage error SendMessage: " + err.Error())
		return emptyValue, fmt.Errorf("storage error SendMessage: %v", err.Error())
	}
	id, err := res.LastInsertId()
	if err != nil {
		slog.Error("storage error SendMessage: " + err.Error())
		return emptyValue, fmt.Errorf("storage error SendMessage: %v", err.Error())
	}
	return id, nil
}

func (s *Storage) GetMessages(ctx context.Context, firstId, secondId int64) ([]*models.Message, error) {

	slog.Info("storage started GetMessages")

	stmt, err := s.db.Prepare("SELECT id,sender_id, reciever_id, file_name, content, send_date, idea_id FROM messages WHERE (sender_id=? AND reciever_id = ?) OR (sender_id=? AND reciever_id = ?) ORDER BY sending_date_seconds")

	if err != nil {
		slog.Error("storage error GetMessages: " + err.Error())
		return nil, fmt.Errorf("storage error GetMessages: %v", err.Error())
	}

	rows, err := stmt.QueryContext(ctx, firstId, secondId, secondId, firstId)

	if err != nil {
		slog.Error("storage error GetMessages: " + err.Error())
		return nil, fmt.Errorf("storage error GetMessages: %v", err.Error())
	}
	var result []*models.Message
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

	stmt, err := s.db.Prepare("INSERT INTO chats(first_id,second_id) VALUES(?,?)")
	if err != nil {
		slog.Error("storage error CreateChat: " + err.Error())
		return nil, fmt.Errorf("storage error CreateChat: %v", err.Error())
	}
	_, err = stmt.ExecContext(ctx, firstId, secondId)
	if err != nil {
		slog.Error("storage error CreateChat: " + err.Error())
		return nil, fmt.Errorf("storage error CreateChat: %v", err.Error())
	}
	return nil, nil
}

func (s *Storage) GetUsersChats(ctx context.Context, userId int64) ([]*models.ChatData, error) {
	slog.Info("storage started GetUsersChats, id = " + fmt.Sprint(userId))

	stmt, err := s.db.Prepare("SELECT id, first_id, second_id FROM chats WHERE first_id = ? OR second_id = ?")

	if err != nil {
		slog.Error("storage error GetUsersChats: " + err.Error())
		return nil, fmt.Errorf("storage error GetUsersChats: %v", err.Error())
	}

	rows, err := stmt.QueryContext(ctx, userId, userId)

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
