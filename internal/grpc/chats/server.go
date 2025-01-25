package chats

import (
	"context"
	"fmt"
	chatsv1 "idea-store-auth/gen/go/chats"
	"idea-store-auth/internal/domain/models"
	"log/slog"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Chats interface {
	SendMessage(ctx context.Context, senderId, recieverId int64, fileName, text, creationDate string) (*emptypb.Empty, error)
	GetMessages(ctx context.Context, senderId, recieverId int64) ([]*models.Message, error)
	CreateChat(ctx context.Context, user1, user2 models.ChatUserData) (*emptypb.Empty, error)
	GetUsersChats(ctx context.Context, userId int64) ([]*models.ChatData, error)
	DeleteChat(ctx context.Context, chatId int64) (*emptypb.Empty, error)
}

type serverAPI struct {
	chatsv1.UnimplementedChatsServer
	chats Chats
}

func Register(gRPC *grpc.Server, chats Chats) {
	chatsv1.RegisterChatsServer(gRPC, &serverAPI{chats: chats})
}

func (s *serverAPI) SendMessage(ctx context.Context, req *chatsv1.SendMessageRequest) (*emptypb.Empty, error) {
	slog.Info("grpc started SendMessage")

	_, err := s.chats.SendMessage(ctx, req.Data.SenderId, req.Data.RecieverId, req.Data.FileName, req.Data.Text, req.Data.SendingDate)
	if err != nil {
		slog.Error("grpc error SendMessage")
		return nil, fmt.Errorf("grpc SendMessage error :%v", err.Error())
	}
	return nil, nil
}
func (s *serverAPI) GetMessages(ctx context.Context, req *chatsv1.GetMessagesRequest) (*chatsv1.GetMessagesResponse, error) {
	slog.Info("grpc started GetMessages")

	resp, err := s.chats.GetMessages(ctx, req.SenderId, req.RecieverId)

	if err != nil {
		slog.Error("grpc error GetMessages")
		return nil, fmt.Errorf("grpc GetMessages error :%v", err.Error())
	}
	var messages []*chatsv1.MessageData
	for _, m := range resp {
		messages = append(messages, &chatsv1.MessageData{
			SenderId:    m.SenderId,
			RecieverId:  m.RecieverId,
			FileName:    m.Filename,
			Text:        m.Text,
			SendingDate: m.Text,
		})
	}
	return &chatsv1.GetMessagesResponse{
		Messages: messages,
	}, nil
}

func (s *serverAPI) CreateChat(ctx context.Context, req *chatsv1.CreateChatRequest) (*emptypb.Empty, error) {
	slog.Info("grpc satarted CreateChat")
	firstData := models.ChatUserData{
		UserId:   req.FirstData.Id,
		Username: req.FirstData.Name,
		Avatar:   req.FirstData.Avatar,
	}
	secondData := models.ChatUserData{
		UserId:   req.SecondData.Id,
		Username: req.SecondData.Name,
		Avatar:   req.SecondData.Avatar,
	}
	_, err := s.chats.CreateChat(ctx, firstData, secondData)

	if err != nil {
		slog.Error("grpc error CreateChat")
		return nil, fmt.Errorf("grpc CreateChat error :%v", err.Error())
	}
	return nil, nil
}

func (s *serverAPI) GetUsersChats(ctx context.Context, req *chatsv1.GetUsersChatsRequest) (*chatsv1.GetUsersChatsResponse, error) {
	slog.Info("grpc started GetUsersChats")

	resp, err := s.chats.GetUsersChats(ctx, req.UserId)
	if err != nil {
		slog.Error("grpc error GetUsersChats")
		return nil, fmt.Errorf("grpc GetUsersChats error :%v", err.Error())
	}
	var chats []*chatsv1.ChatData
	for _, chat := range resp {
		if chat.User1.UserId == req.UserId {
			chats = append(chats, &chatsv1.ChatData{
				Id:     chat.User2.UserId,
				Name:   chat.User2.Username,
				Avatar: chat.User2.Avatar,
			})
		} else {
			chats = append(chats, &chatsv1.ChatData{
				Id:     chat.User1.UserId,
				Name:   chat.User1.Username,
				Avatar: chat.User1.Avatar,
			})
		}
	}
	return &chatsv1.GetUsersChatsResponse{
		Chats: chats,
	}, nil
}

func (s *serverAPI) DeleteChat(ctx context.Context, req *chatsv1.DeleteChatRequest) (*emptypb.Empty, error) {
	slog.Info("grpc started to DeleteChat")

	_, err := s.chats.DeleteChat(ctx, req.ChatId)
	if err != nil {
		slog.Error("grpc error DeleteChat")
		return nil, fmt.Errorf("grpc DeleteChat error :%v", err.Error())
	}
	return nil, nil
}
