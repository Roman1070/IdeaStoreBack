package app

import (
	grpcApp "idea-store-auth/internal/app/grpc"
	"idea-store-auth/internal/services/auth"
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
	authStoragePath string,
	ideasStoragePath string,
	tokenTTL time.Duration,
) *App {
	authStorage, err := sqlite.New(authStoragePath)
	if err != nil {
		panic(err)
	}
	ideasStorage, err := sqlite.New(ideasStoragePath)
	if err != nil {
		panic(err)
	}

	authService := auth.New(log, authStorage, authStorage, authStorage, tokenTTL)
	ideasService := ideas.New(log,ideasStorage)

	grpcApp := grpcApp.New(log, authService, ideasService.Api, grpcPort)

	return &App{
		GRPCServer: grpcApp,
	}
}
