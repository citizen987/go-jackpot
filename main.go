package main

import (
	cfg "go-jackpot/internal/config"
	lgc "go-jackpot/internal/logic"
	svc "go-jackpot/internal/server"
	srv "go-jackpot/internal/service"
	stg "go-jackpot/internal/storage"
)

func main() {
	config := cfg.LoadConfig("config/config.json")
	logic := lgc.NewJackpotLogic(config)
	storage := stg.NewStorage("jackpot-log.jsonl")
	service := srv.NewJackpotService(logic, storage)
	server := svc.NewServer(service)
	server.Start()
}
