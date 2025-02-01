package main

import (
	common "idea-store-auth/cmd"
	appIdeas "idea-store-auth/internal/app/ideas"
	"idea-store-auth/internal/config"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := config.MustLoad()

	log := common.SetupLogger(cfg.Env)

	ideasApp := appIdeas.New(log, cfg.GRPC.IdeasMS.Port)

	go func() {
		ideasApp.GRPCServer.MustRun()
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop

	ideasApp.GRPCServer.Stop()
	log.Info("Gracefully stopped")
}
