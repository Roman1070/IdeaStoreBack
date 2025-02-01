package app

import (
	common "idea-store-auth/internal/app"
	grpcApp "idea-store-auth/internal/app/grpc/auth"
	"idea-store-auth/internal/services/auth"
	sqlite "idea-store-auth/internal/storage/postgre"
	"log/slog"
	"time"
)

type App struct {
	GRPCServer *common.App
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
