package main

import (
	cfg "go-jackpot/internal/config"
	lgc "go-jackpot/internal/logic"
	svc "go-jackpot/internal/server"
	srv "go-jackpot/internal/service"
	stg "go-jackpot/internal/storage"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	server, storage := initApp()
	runServer(server)
	waitForShutdown()
	shutdown(server, storage)
}

func initApp() (*svc.Server, stg.Storage) {
	config := cfg.LoadConfig("config/config.json")
	logic := lgc.NewJackpotLogic(config)
	storage := stg.NewStorage(stg.StorageTypeAsync, "jackpot-log.jsonl")
	service := srv.NewJackpotService(logic, storage)
	server := svc.NewServer(service)
	return server, storage
}

func runServer(server *svc.Server) {
	server.StartAsync()
}

func waitForShutdown() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop
}

func shutdown(server *svc.Server, storage stg.Storage) {
	server.Stop()
	storage.Close()
}
