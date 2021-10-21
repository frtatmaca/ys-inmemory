package repository

import (
	"encoding/gob"
	"log"
	"os"
)

type FileGobRepository struct{}

func NewFileGobRepository() *FileGobRepository {
	return &FileGobRepository{}
}

func (r *FileGobRepository) Read() (map[string]string) {
	keyValStore := make(map[string]string, 0)

	fileName := "test.gob"

	if _, err := os.Stat(fileName); err == nil {
		f, err := os.Open(fileName)
		defer f.Close()

		if err != nil {
			log.Panicf("failed reading data from file: %s", err)
		}

		d := gob.NewDecoder(f)
		err = d.Decode(&keyValStore)

		if err != nil {
			panic(err)
		}
	} else if os.IsNotExist(err) {
		_, err := os.Create(fileName)

		if err != nil {
			log.Fatalf("failed creating file: %s", err)
		}
	}

	return keyValStore
}

func (r *FileGobRepository) WriteFile(keyValStore map[string]string) (bool) {
	fileName := "test.gob"

	os.Remove(fileName)

	f, _ := os.Create(fileName)

	enc := gob.NewEncoder(f)
	enc.Encode(keyValStore)

	f.Close()
	return true
}
