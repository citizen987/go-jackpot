package storage

import (
	"bufio"
	"encoding/json"
	"os"
	"testing"
)

func testStorageJsonl_createFile(t *testing.T) *os.File {
	tmpfile, err := os.CreateTemp("", "jackpot-temp.jsonl")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	return tmpfile
}

func testStorageJsonl_getLog() JackpotLog {
	return JackpotLog{
		Timestamp: "2025-06-27T12:00:00Z",
		Bet:       10,
		IsWon:     true,
	}
}

func testStorageJsonl_checkSave(t *testing.T, tmpfile *os.File, entry JackpotLog) {
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

func TestStorageJsonl_Save(t *testing.T) {
	tmpfile := testStorageJsonl_createFile(t)
	defer os.Remove(tmpfile.Name())
	entry := testStorageJsonl_getLog()
	storage := NewStorageJsonl(tmpfile.Name())

	storage.Save(entry)
	storage.Close()

	testStorageJsonl_checkSave(t, tmpfile, entry)
}
