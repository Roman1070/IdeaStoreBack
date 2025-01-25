package main

import (
	common "idea-store-auth/cmd"
	appChats "idea-store-auth/internal/app/chats"
	"idea-store-auth/internal/config"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := config.MustLoad()

	log := common.SetupLogger(cfg.Env)

	chatsApp := appChats.New(log, cfg.GRPC.ChatsMS.Port, cfg.ChatsStoragePath)

	go func() {
		chatsApp.GRPCServer.MustRun()
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop

	chatsApp.GRPCServer.Stop()
	log.Info("Gracefully stopped")
}
