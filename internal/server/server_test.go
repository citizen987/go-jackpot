package server

import (
	"bytes"
	"encoding/json"
	"go-jackpot/internal/service"
	"net/http"
	"net/http/httptest"
	"testing"
)

type mockService struct {
	mockResult service.JackpotResponse
}

func (m *mockService) CheckWon(req service.JackpotRequest) service.JackpotResponse {
	return m.mockResult
}

func TestHandleJackpotDraw_OK(t *testing.T) {
	mock := &mockService{
		mockResult: service.JackpotResponse{IsWon: true},
	}
	s := &Server{service: mock}

	body := service.JackpotRequest{Bet: 1}
	jsonBody, _ := json.Marshal(body)

	req := httptest.NewRequest(http.MethodPost, "/jackpot-draw", bytes.NewReader(jsonBody))
	w := httptest.NewRecorder()

	s.handleJackpotDraw(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	var result service.JackpotResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		t.Fatalf("Error decoding response: %v", err)
	}

	if !result.IsWon {
		t.Errorf("Expected IsWon to be true, got %v", result.IsWon)
	}
}

func TestHandleJackpotDraw_InvalidMethod(t *testing.T) {
	mock := &mockService{}
	s := &Server{service: mock}

	req := httptest.NewRequest(http.MethodGet, "/jackpot-draw", nil)
	w := httptest.NewRecorder()

	s.handleJackpotDraw(w, req)

	if w.Result().StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("Expected status code %d, got %d", http.StatusMethodNotAllowed, w.Result().StatusCode)
	}
}

func TestHandleJackpotDraw_InvalidJson(t *testing.T) {
	mock := &mockService{}
	s := &Server{service: mock}

	body := bytes.NewBufferString("not-json")
	req := httptest.NewRequest(http.MethodPost, "/jackpot-draw", body)
	w := httptest.NewRecorder()

	s.handleJackpotDraw(w, req)

	if w.Result().StatusCode != http.StatusBadRequest {
		t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, w.Result().StatusCode)
	}
}
