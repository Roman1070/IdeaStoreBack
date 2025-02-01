package common

import (
	"fmt"
	"idea-store-auth/internal/config"
	"idea-store-auth/internal/lib/logger/handlers/slogpretty"
	"log/slog"
	"os"
)

const (
	EnvLocal = "local"
	EnvDev   = "dev"
	EnvProd  = "prod"
)

func SetupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case EnvLocal:
		log = SetupPrettySlog()
	case EnvDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case EnvProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}

func SetupPrettySlog() *slog.Logger {
	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}

func GrpcBoardsAddress(cfg *config.Config) string {
	return fmt.Sprintf("boards_go:%v", cfg.GRPC.BoardsMS.Port)
}
func GrpcAuthAddress(cfg *config.Config) string {
	return fmt.Sprintf("auth_go:%v", cfg.GRPC.AuthMS.Port)
}
func GrpcIdeasAddress(cfg *config.Config) string {
	return fmt.Sprintf("ideas_go:%v", cfg.GRPC.IdeasMS.Port)
}
func GrpcProfilesAddress(cfg *config.Config) string {
	return fmt.Sprintf("profiles_go:%v", cfg.GRPC.ProfilesMS.Port)
}
func GrpcCommentsAddress(cfg *config.Config) string {
	return fmt.Sprintf("comments_go:%v", cfg.GRPC.CommentsMS.Port)
}
func GrpcChatsAddress(cfg *config.Config) string {
	return fmt.Sprintf("chats_go:%v", cfg.GRPC.ChatsMS.Port)
}
