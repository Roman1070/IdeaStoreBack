package chats

import (
	"context"
	"fmt"
	"idea-store-auth/internal/domain/models"
	"idea-store-auth/internal/grpc/chats"
	"log/slog"

	"google.golang.org/protobuf/types/known/emptypb"
)

type Chats struct {
	log *slog.Logger
	Api chats.Chats
}

func New(log *slog.Logger, api chats.Chats) *Chats {
	return &Chats{
		log: log,
		Api: api,
	}
}
func (c *Chats) SendMessage(ctx context.Context, message models.Message) (*emptypb.Empty, error) {
	c.log.Info("service started SendMessage")

	_, err := c.Api.SendMessage(ctx, models.Message{
		SenderId:           message.SenderId,
		RecieverId:         message.RecieverId,
		Filename:           message.Filename,
		Text:               message.Text,
		CreationDate:       message.CreationDate,
		CheckChatExistance: message.CheckChatExistance,
	})
	if err != nil {
		c.log.Error("service error SendMessage: " + err.Error())
		return nil, fmt.Errorf("serivce error SendMessage: %v", err.Error())
	}

	return nil, nil
}

func (c *Chats) GetMessages(ctx context.Context, senderId, recieverId int64) ([]*models.Message, error) {
	c.log.Info("service started GetMessages")

	resp, err := c.Api.GetMessages(ctx, senderId, recieverId)

	if err != nil {
		c.log.Error("service error GetMessages: " + err.Error())
		return nil, fmt.Errorf("serivce error GetMessages: %v", err.Error())
	}

	return resp, nil
}

func (c *Chats) CreateChat(ctx context.Context, user1, user2 int64) (*emptypb.Empty, error) {
	c.log.Info("service started CreateChat")
	_, err := c.Api.CreateChat(ctx, user1, user2)

	if err != nil {
		c.log.Error("service error CreateChat: " + err.Error())
		return nil, fmt.Errorf("serivce error CreateChat: %v", err.Error())
	}

	return nil, nil
}

func (c *Chats) GetUsersChats(ctx context.Context, userId int64) ([]*models.ChatData, error) {
	c.log.Info("service started GetUsersChats")

	resp, err := c.Api.GetUsersChats(ctx, userId)

	if err != nil {
		c.log.Error("service error GetUsersChats: " + err.Error())
		return nil, fmt.Errorf("serivce error GetUsersChats: %v", err.Error())
	}

	return resp, nil
}

func (c *Chats) DeleteChat(ctx context.Context, chatId int64) (*emptypb.Empty, error) {
	c.log.Info("service started DeleteChat")
	_, err := c.Api.DeleteChat(ctx, chatId)

	if err != nil {
		c.log.Error("service error DeleteChat: " + err.Error())
		return nil, fmt.Errorf("serivce error DeleteChat: %v", err.Error())
	}

	return nil, nil
}
