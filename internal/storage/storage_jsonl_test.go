package storage

import (
	"bufio"
	"encoding/json"
	"os"
	"testing"
)

func TestStorageJsonl_Save(t *testing.T) {
	tmpfile, err := os.CreateTemp("", "jackpot-temp.jsonl")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpfile.Name())

	storage := NewStorage(tmpfile.Name())
	entry := JackpotLog{
		Timestamp: "2025-06-27T12:00:00Z",
		Bet:       10,
		IsWon:     true,
	}
	storage.Save(entry)

	file, err := os.Open(tmpfile.Name())
	if err != nil {
		t.Fatalf("Failed to open temp file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if !scanner.Scan() {
		t.Fatal("No data written to file")
	}

	line := scanner.Text()
	var result JackpotLog
	if err := json.Unmarshal([]byte(line), &result); err != nil {
		t.Fatalf("Failed to unmarshal JSON: %v", err)
	}
	if result != entry {
		t.Errorf("Expected %v, got %v", entry, result)
	}
	if scanner.Scan() {
		t.Error("Expected only one line in the file, but found more")
	}
}
