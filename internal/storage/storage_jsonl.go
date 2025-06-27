package storage

import (
	"encoding/json"
	"log"
	"os"
)

type storageJsonl struct {
	filename string
}

func NewStorage(filename string) Storage {
	return &storageJsonl{filename: filename}
}

func (s *storageJsonl) Save(data JackpotLog) {
	file, err := os.OpenFile(s.filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println("Error oppening file:", err)
		return
	}
	defer file.Close()

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
