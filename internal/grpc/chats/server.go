package chats

import (
	"context"
	chatsv1 "idea-store-auth/gen/go/chats"
	"idea-store-auth/internal/domain/models"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Chats interface {
	SendMessage(ctx context.Context, senderId, recieverId int64, fileName, text, creationDate string) (*emptypb.Empty, error)
	GetMessages(ctx context.Context, senderId, recieverId int64) (*[]models.Message, error)
	CreateChat(ctx context.Context, user1, user2 models.ChatUserData) (*emptypb.Empty, error)
	GetUsersChats(ctx context.Context, userId int64) (*[]models.ChatData, error)
	DeleteChar(ctx context.Context, chatId int64) (*emptypb.Empty, error)
}

type serverAPI struct {
	chatsv1.UnimplementedChatsServer
	chats Chats
}

func Register(gRPC *grpc.Server, chats Chats) {
	chatsv1.RegisterChatsServer(gRPC, &serverAPI{chats: chats})
}
