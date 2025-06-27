package storage

type JackpotLog struct {
	Timestamp string `json:"timestamp"`
	Bet       int    `json:"bet"`
	IsWon     bool   `json:"boolean"`
}

type Storage interface {
	Save(data JackpotLog)
}
