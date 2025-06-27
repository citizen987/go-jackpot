package logic

import (
	"math/rand"
	"time"
)

type JackpotLogic interface {
	Bet(bet int) bool
}

type jackpotLogicImpl struct {
	rng *rand.Rand
}

func NewJackpotLogic() JackpotLogic {
	src := rand.NewSource(time.Now().UnixNano())
	return &jackpotLogicImpl{
		rng: rand.New(src),
	}
}

func (jl *jackpotLogicImpl) Bet(bet int) bool {
	if bet <= 0 {
		panic("Bet must be grater than 0")
	}
	isWon := jl.rng.Intn(100)/bet == 0
	return isWon
}
