package service

import (
	"go-jackpot/internal/logic"
	"go-jackpot/internal/storage"
	"time"
)

type JackpotService interface {
	CheckWon(rq JackpotRequest) JackpotResponse
}

type JackpotResponse struct {
	IsWon bool `json:"is_won"`
}

type JackpotRequest struct {
	Bet int `json:"bet"`
}

type jackpotServiceImpl struct {
	jackpotLogic logic.JackpotLogic
	storage      storage.Storage
}

func NewJackpotService(jackpotLogic logic.JackpotLogic, storage storage.Storage) JackpotService {
	return &jackpotServiceImpl{
		jackpotLogic: jackpotLogic,
		storage:      storage,
	}
}

func (s *jackpotServiceImpl) CheckWon(rq JackpotRequest) JackpotResponse {
	isWon := s.jackpotLogic.Bet(rq.Bet)

	entry := storage.JackpotLog{
		Timestamp: time.Now().Format(time.RFC3339),
		Bet:       rq.Bet,
		IsWon:     isWon,
	}
	s.storage.Save(entry)

	return JackpotResponse{IsWon: isWon}
}
