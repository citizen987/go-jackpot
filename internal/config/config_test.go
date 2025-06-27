package config

import (
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

func LoadConfigFromRoot() Config {
	_, b, _, _ := runtime.Caller(0)
	base := filepath.Dir(b)
	projectRoot := filepath.Join(base, "..", "..")
	configPath := filepath.Join(projectRoot, "config", "config.json")
	println(configPath)
	return LoadConfig(configPath)
}

func TestLoadConfig_RealFile(t *testing.T) {
	cfg := LoadConfigFromRoot()
	if cfg.MinBet <= 0 {
		t.Errorf("expected MinBet > 0, got %d", cfg.MinBet)
	}
	if cfg.MaxBet < cfg.MinBet {
		t.Errorf("expected MaxBet >= MinBet, got MaxBet=%d, MinBet=%d", cfg.MaxBet, cfg.MinBet)
	}
}

func TestLoadConfig_Success(t *testing.T) {
	tmpFile, err := os.CreateTemp("", "config_test_*.json")
	if err != nil {
		t.Fatalf("Error creating temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	configAux := Config{
		MinBet: 1,
		MaxBet: 50,
	}
	configAuxBytes, err := json.Marshal(configAux)
	if err != nil {
		t.Fatalf("Error on marshal aux config: %v", err)
	}
	content := string(configAuxBytes)
	if _, err := tmpFile.WriteString(content); err != nil {
		t.Fatalf("Error writing temp file: %v", err)
	}
	tmpFile.Close()

	cfg := LoadConfig(tmpFile.Name())
	if cfg != configAux {
		t.Errorf("Expected %d got %d", configAux, cfg)
	}
}
