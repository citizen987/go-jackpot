package server

import (
	"context"
	"encoding/json"
	"go-jackpot/internal/service"
	"log"
	"net/http"
	"time"
)

type Server struct {
	service    service.JackpotService
	httpServer *http.Server
}

func NewServer(service service.JackpotService) *Server {
	mux := http.NewServeMux()
	server := &Server{
		service: service,
		httpServer: &http.Server{
			Addr:    ":8080",
			Handler: mux,
		},
	}
	mux.HandleFunc("/jackpot-draw", server.handleJackpotDraw)
	return server
}

func (s *Server) Start() {
	log.Println("Server is running on http://localhost:8080")
	err := s.httpServer.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		panic("Error starting server: " + err.Error())
	}
}

func (s *Server) StartAsync() {
	go func() {
		s.Start()
	}()
}

func (s *Server) Stop() {
	log.Print("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.httpServer.Shutdown(ctx); err != nil {
		log.Printf("Error shutting down server: %v", err)
	}
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
