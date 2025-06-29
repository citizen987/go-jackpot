package storage

import (
	"os"
	"testing"
)

func TestStorageJsonlAsync_Save(t *testing.T) {
	tmpfile := testStorageJsonl_createFile(t)
	defer os.Remove(tmpfile.Name())
	entry := testStorageJsonl_getLog()
	storage := NewStorageJsonlAsync(tmpfile.Name())

	storage.Save(entry)
	storage.Close()

	testStorageJsonl_checkSave(t, tmpfile, entry)
}
