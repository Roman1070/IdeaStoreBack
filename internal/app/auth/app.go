package app

import (
	grpcApp "idea-store-auth/internal/app/grpc/auth"
	"idea-store-auth/internal/services/auth"
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
	tokenTTL time.Duration,
) *App {
	authStorage, err := sqlite.New(authStoragePath)
	if err != nil {
		panic(err)
	}

	authService := auth.New(log, authStorage, authStorage, authStorage, tokenTTL)

	grpcApp := grpcApp.New(log, authService, grpcPort)

	return &App{
		GRPCServer: grpcApp,
	}
}
