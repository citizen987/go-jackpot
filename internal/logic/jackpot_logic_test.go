package logic

import (
	"go-jackpot/internal/config"
	"testing"
)

func NewMockConfig() config.Config {
	return config.Config{
		MinBet: 1,
		MaxBet: 100,
	}
}

func TestExec100(t *testing.T) {
	cfg := NewMockConfig()
	logic := NewJackpotLogic(cfg)
	for i := 0; i < 100; i++ {
		logic.Bet(1)
	}
}

func TestBetZero(t *testing.T) {
	cfg := NewMockConfig()
	jl := NewJackpotLogic(cfg)
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic for bet of 0, but did not panic")
		}
	}()
	jl.Bet(0)
}
