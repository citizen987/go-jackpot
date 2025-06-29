package storage

import (
	"log"
	"time"
)

type storageJsonlAsnyc struct {
	storageJsonl
	logs chan JackpotLog
	quit chan struct{}
}

func NewStorageJsonlAsync(filename string) *storageJsonlAsnyc {
	storage := &storageJsonlAsnyc{
		storageJsonl: storageJsonl{filename: filename},
		logs:         make(chan JackpotLog, 100),
		quit:         make(chan struct{}),
	}
	go storage.processWriter()
	return storage
}

func (s *storageJsonlAsnyc) Save(data JackpotLog) {
	s.logs <- data
}

func (s *storageJsonlAsnyc) Close() {
	s.flush()
	close(s.quit)
}

func (s *storageJsonlAsnyc) processWriter() {
	file, err := s.openFile()
	if err != nil {
		log.Println("Error oppening file:", err)
		return
	}
	defer file.Close()

	for {
		select {
		case data := <-s.logs:
			s.saveOnFile(data, file)
		case <-s.quit:
			log.Println("Storage writter shuting down.")
			return
		}
	}
}

func (s *storageJsonlAsnyc) flush() {
	for {
		if len(s.logs) == 0 {
			return
		}
		time.Sleep(10 * time.Millisecond)
	}
}
