package app

import (
	common "idea-store-auth/internal/app"
	grpcApp "idea-store-auth/internal/app/grpc/chats"
	"idea-store-auth/internal/services/chats"
	"idea-store-auth/internal/storage/postgre"
	"log/slog"
)

type App struct {
	GRPCServer *common.App
}

func New(
	log *slog.Logger,
	grpcPort int,
) *App {
	chatsStorage, err := postgre.New()
	if err != nil {
		panic(err)
	}

	chatsService := chats.New(log, chatsStorage)

	grpcApp := grpcApp.New(log, chatsService, grpcPort)

	return &App{
		GRPCServer: grpcApp,
	}
}
