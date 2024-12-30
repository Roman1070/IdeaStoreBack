package main

import (
	common "idea-store-auth/cmd"
	appBoards "idea-store-auth/internal/app/boards"
	"idea-store-auth/internal/config"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := config.MustLoad()

	log := common.SetupLogger(cfg.Env)

	boardsApp := appBoards.New(log, cfg.GRPC.BoardsMS.Port, cfg.BoardsStoragePath)

	go func() {
		boardsApp.GRPCServer.MustRun()
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop

	boardsApp.GRPCServer.Stop()
	log.Info("Gracefully stopped")
}