package logic

import (
	"testing"
)

func TestExec100(t *testing.T) {
	logic := NewJackpotLogic()
	for i := 0; i < 100; i++ {
		logic.Bet(1)
	}
}

func TestBetZero(t *testing.T) {
	jl := NewJackpotLogic()
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic for bet of 0, but did not panic")
		}
	}()
	jl.Bet(0)
}
