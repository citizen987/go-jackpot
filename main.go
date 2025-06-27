package main

import (
	"go-jackpot/internal/logic"
	"go-jackpot/internal/server"
	"go-jackpot/internal/service"
	"go-jackpot/internal/storage"
)

func main() {
	logic := logic.NewJackpotLogic()
	storage := storage.NewStorage("jackpot-log.jsonl")
	service := service.NewJackpotService(logic, storage)
	server := server.NewServer(service)
	server.Start()
}
