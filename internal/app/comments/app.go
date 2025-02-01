package app

import (
	common "idea-store-auth/internal/app"
	grpcApp "idea-store-auth/internal/app/grpc/comments"
	"idea-store-auth/internal/services/comments"
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
	storage, err := postgre.New()
	if err != nil {
		panic(err)
	}

	commentsService := comments.New(log, storage)

	grpcApp := grpcApp.New(log, commentsService, grpcPort)

	return &App{
		GRPCServer: grpcApp,
	}
}
