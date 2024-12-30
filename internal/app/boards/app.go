package app

import (
	common "idea-store-auth/internal/app"
	grpcApp "idea-store-auth/internal/app/grpc/boards"
	"idea-store-auth/internal/services/boards"
	"idea-store-auth/internal/storage/sqlite"
	"log/slog"
)

type App struct {
	GRPCServer *common.App
}

func New(
	log *slog.Logger,
	grpcPort int,
	storagePath string,
) *App {
	boardsStorage, err := sqlite.New(storagePath)
	if err != nil {
		panic(err)
	}

	boardsService := boards.New(log, boardsStorage)

	grpcApp := grpcApp.New(log, boardsService, grpcPort)

	return &App{
		GRPCServer: grpcApp,
	}
}
