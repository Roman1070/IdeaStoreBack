package main

import (
	common "idea-store-auth/cmd"
	appComments "idea-store-auth/internal/app/comments"
	"idea-store-auth/internal/config"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := config.MustLoad()

	log := common.SetupLogger(cfg.Env)

	commentsApp := appComments.New(log, cfg.GRPC.CommentsMS.Port, cfg.CommentsStoragePath)

	go func() {
		commentsApp.GRPCServer.MustRun()
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop

	commentsApp.GRPCServer.Stop()
	log.Info("Gracefully stopped")
}
