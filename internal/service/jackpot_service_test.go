package service

import (
	"go-jackpot/internal/storage"
	"testing"
)

type mockJackpotLogic struct {
	betCalledWith int
	betResult     bool
}

func (m *mockJackpotLogic) Bet(bet int) bool {
	m.betCalledWith = bet
	return m.betResult
}

type mockStorage struct {
	savedLog *storage.JackpotLog
}

func (m *mockStorage) Save(data storage.JackpotLog) {
	m.savedLog = &data
}

func (m *mockStorage) Close() {
}

func TestCheckWon_CallsLogicAndStorage(t *testing.T) {
	logicMock := &mockJackpotLogic{betResult: true}
	storageMock := &mockStorage{}
	service := NewJackpotService(logicMock, storageMock)
	req := JackpotRequest{Bet: 10}
	resp := service.CheckWon(req)

	if !resp.IsWon {
		t.Errorf("Expected IsWon to be true, got %v", resp.IsWon)
	}
	if logicMock.betCalledWith != req.Bet {
		t.Errorf("Expected Bet to be called with %d, got %d", req.Bet, logicMock.betCalledWith)
	}
	if storageMock.savedLog == nil {
		t.Errorf("Expected storage to save a log entry, but it was nil")
	}
	if storageMock.savedLog.Bet != req.Bet || storageMock.savedLog.IsWon != resp.IsWon {
		t.Errorf("Expected saved log to match request and response, got Bet: %d, IsWon: %v",
			storageMock.savedLog.Bet, storageMock.savedLog.IsWon)
	}
}
