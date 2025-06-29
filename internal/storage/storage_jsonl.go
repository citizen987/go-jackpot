package storage

import (
	"encoding/json"
	"log"
	"os"
)

type storageJsonl struct {
	filename string
}

func NewStorageJsonl(filename string) *storageJsonl {
	return &storageJsonl{filename: filename}
}

func (s *storageJsonl) Save(data JackpotLog) {
	file, err := s.openFile()
	if err != nil {
		log.Println("Error oppening file:", err)
		return
	}
	defer file.Close()

	s.saveOnFile(data, file)
}

func (s *storageJsonl) Close() {
}

func (s *storageJsonl) openFile() (*os.File, error) {
	return os.OpenFile(s.filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
}

func (s *storageJsonl) saveOnFile(data JackpotLog, file *os.File) {
	jsonLine, err := json.Marshal(data)
	if err != nil {
		log.Println("Error marshalling JSON:", err)
		return
	}

	_, err = file.Write(append(jsonLine, '\n'))
	if err != nil {
		log.Println("Error writing to file:", err)
	}
}
