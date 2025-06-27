package server

import (
	"encoding/json"
	"go-jackpot/internal/service"
	"log"
	"net/http"
)

type Server struct {
	service service.JackpotService
}

func NewServer(service service.JackpotService) *Server {
	return &Server{service: service}
}

func (s *Server) Start() {
	http.HandleFunc("/jackpot-draw", s.handleJackpotDraw)
	log.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func (s *Server) handleJackpotDraw(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("Recovered from panic: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
	}()

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req service.JackpotRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	result := s.service.CheckWon(req)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
