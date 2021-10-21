package repository

var KeyValStore = map[string]string{}

type Repository interface {
	Read() (map[string]string)
	WriteFile(map[string]string) (bool)
}