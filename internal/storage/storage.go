package storage

type JackpotLog struct {
	Timestamp string `json:"timestamp"`
	Bet       int    `json:"bet"`
	IsWon     bool   `json:"is_won"`
}

type Storage interface {
	Save(data JackpotLog)
	Close()
}
