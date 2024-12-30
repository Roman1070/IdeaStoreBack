package profiles

import (
	common "idea-store-auth/internal/app"
	grpcApp "idea-store-auth/internal/app/grpc/profiles"
	"idea-store-auth/internal/services/profiles"
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
	storage, err := sqlite.New(storagePath)
	if err != nil {
		panic(err)
	}

	profilesService := profiles.New(log, storage)

	grpcApp := grpcApp.New(log, profilesService, grpcPort)

	return &App{
		GRPCServer: grpcApp,
	}
}
