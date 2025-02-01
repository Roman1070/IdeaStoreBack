package app

import (
	common "idea-store-auth/internal/app"
	grpcApp "idea-store-auth/internal/app/grpc/ideas"
	"idea-store-auth/internal/services/ideas"
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

	ideasService := ideas.New(log, storage)

	grpcApp := grpcApp.New(log, ideasService, grpcPort)

	return &App{
		GRPCServer: grpcApp,
	}
}
