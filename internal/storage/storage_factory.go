package storage

type StorageType string

const (
	SotrageTypeSync  StorageType = "sync"
	StorageTypeAsync StorageType = "async"
)

func NewStorage(storageType StorageType, filename string) Storage {
	switch storageType {
	case SotrageTypeSync:
		return NewStorageJsonl(filename)
	case StorageTypeAsync:
		return NewStorageJsonlAsync(filename)
	default:
		panic("Unknown storage type: " + string(storageType))
	}
}
