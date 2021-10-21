package core

import (
	"github.com/frtatmaca/ys-inmemory/models"
	"github.com/frtatmaca/ys-inmemory/repository"
)

type RepositoryAdapter struct{}

func NewRepositoryAdapter() RepositoryAdapter {
	return RepositoryAdapter{}
}

func (s *RepositoryAdapter) NewFileRepository(repoType string) (repo repository.Repository, err error) {
	switch repoType {
	case models.GOB_REPOSITORY:
		repo = repository.NewFileGobRepository()
	case models.JSON_REPOSITORY:
		repo = repository.NewFileJsonRepository()
	default:
		repo = repository.NewFileGobRepository()
	}

	return
}
