package repository

import (
	"encoding/json"
	"io/ioutil"
)

type FileJsonRepository struct{}

func NewFileJsonRepository() *FileJsonRepository {
	return &FileJsonRepository{}
}

func (r *FileJsonRepository) Read() (map[string]string) {
	content, err := ioutil.ReadFile("test2.gob")
	if err != nil {
		panic(err)
	}

	var data map[string]interface{}
	err = json.Unmarshal([]byte(content), &data)

	if err != nil {
		panic(err)
	}

	return make(map[string]string, 0)
}

func (r *FileJsonRepository) WriteFile(keyValStore map[string]string) (bool) {

	return true
}
