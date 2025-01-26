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
	SendMessage(ctx context.Context, message models.Message) (int64, error)
	GetMessages(ctx context.Context, senderId, recieverId int64) ([]*models.Message, error)
	CreateChat(ctx context.Context, user1, user2 int64) (*emptypb.Empty, error)
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

func (s *serverAPI) SendMessage(ctx context.Context, req *chatsv1.SendMessageRequest) (*chatsv1.SendMessageResponse, error) {
	slog.Info("grpc started SendMessage")

	resp, err := s.chats.SendMessage(ctx, models.Message{
		SenderId:           req.Data.SenderId,
		RecieverId:         req.Data.RecieverId,
		Filename:           req.Data.FileName,
		Text:               req.Data.Text,
		CreationDate:       req.Data.SendingDate,
		CheckChatExistance: req.Data.CheckChatExistance,
	})
	if err != nil {
		slog.Error("grpc error SendMessage")
		return nil, fmt.Errorf("grpc SendMessage error :%v", err.Error())
	}
	return &chatsv1.SendMessageResponse{
		Id: resp,
	}, nil
}
func (s *serverAPI) GetMessages(ctx context.Context, req *chatsv1.GetMessagesRequest) (*chatsv1.GetMessagesResponse, error) {
	slog.Info("grpc started GetMessages")

	resp, err := s.chats.GetMessages(ctx, req.FirstId, req.SecondId)

	if err != nil {
		slog.Error("grpc error GetMessages")
		return nil, fmt.Errorf("grpc GetMessages error :%v", err.Error())
	}
	var messages []*chatsv1.MessageData
	for _, m := range resp {
		messages = append(messages, &chatsv1.MessageData{
			Id:                 m.ID,
			SenderId:           m.SenderId,
			RecieverId:         m.RecieverId,
			FileName:           m.Filename,
			Text:               m.Text,
			SendingDate:        m.CreationDate,
			CheckChatExistance: m.CheckChatExistance,
		})
	}
	return &chatsv1.GetMessagesResponse{
		Messages: messages,
	}, nil
}

func (s *serverAPI) CreateChat(ctx context.Context, req *chatsv1.CreateChatRequest) (*emptypb.Empty, error) {
	slog.Info("grpc satarted CreateChat")
	_, err := s.chats.CreateChat(ctx, req.FirstId, req.SecondId)

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
	var chats []*chatsv1.ChatUserData
	for _, chat := range resp {

		if chat.FirstData.UserId == req.UserId {
			chats = append(chats, &chatsv1.ChatUserData{
				Id:     chat.SecondData.UserId,
				Name:   chat.SecondData.Username,
				Avatar: chat.SecondData.Avatar,
			})
		} else {
			chats = append(chats, &chatsv1.ChatUserData{
				Id:     chat.FirstData.UserId,
				Name:   chat.FirstData.Username,
				Avatar: chat.FirstData.Avatar,
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
