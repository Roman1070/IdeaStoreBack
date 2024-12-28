package app

import (
	grpcApp "idea-store-auth/internal/app/grpc/ideas"
	"idea-store-auth/internal/services/ideas"
	"idea-store-auth/internal/storage/sqlite"
	"log/slog"
	"time"
)

type App struct {
	GRPCServer *grpcApp.App
}

func New(
	log *slog.Logger,
	grpcPort int,
	storagePath string,
	tokenTTL time.Duration,
) *App {
	ideasStorage, err := sqlite.New(storagePath)
	if err != nil {
		panic(err)
	}

	ideasService := ideas.New(log, ideasStorage)

	grpcApp := grpcApp.New(log, ideasService, grpcPort)

	return &App{
		GRPCServer: grpcApp,
	}
}
