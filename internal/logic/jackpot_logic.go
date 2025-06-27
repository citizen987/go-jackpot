package logic

import (
	"fmt"
	"go-jackpot/internal/config"
	"math/rand"
	"time"
)

type JackpotLogic interface {
	Bet(bet int) bool
}

type jackpotLogicImpl struct {
	rng    *rand.Rand
	minBet int
	maxBet int
}

func NewJackpotLogic(cfg config.Config) JackpotLogic {
	src := rand.NewSource(time.Now().UnixNano())
	return &jackpotLogicImpl{
		rng:    rand.New(src),
		minBet: cfg.MinBet,
		maxBet: cfg.MaxBet,
	}
}

func (jl *jackpotLogicImpl) Bet(bet int) bool {
	if bet <= 0 {
		panic("Bet must be greater than 0")
	}
	if bet < jl.minBet || bet > jl.maxBet {
		panic(fmt.Sprintf("Bet must be between %d and %d", jl.minBet, jl.maxBet))
	}
	return jl.rng.Intn(100)/bet == 0
}
