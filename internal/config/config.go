package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	MinBet int `json:"min_bet"`
	MaxBet int `json:"max_bet"`
}

func LoadConfig(path string) Config {
	f, err := os.Open(path)
	if err != nil {
		panic(fmt.Sprintf("error opening config file '%s': %v", path, err))
	}
	defer f.Close()

	var cfg Config
	if err := json.NewDecoder(f).Decode(&cfg); err != nil {
		panic(fmt.Sprintf("error parsing config file '%s': %v", path, err))
	}
	return cfg
}
