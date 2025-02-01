package main

import (
	common "idea-store-auth/cmd"
	appProfiles "idea-store-auth/internal/app/profiles"
	"idea-store-auth/internal/config"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := config.MustLoad()

	log := common.SetupLogger(cfg.Env)

	profilesApp := appProfiles.New(log, cfg.GRPC.ProfilesMS.Port)

	go func() {
		profilesApp.GRPCServer.MustRun()
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop

	profilesApp.GRPCServer.Stop()
	log.Info("Gracefully stopped")
}
