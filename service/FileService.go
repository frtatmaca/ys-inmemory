package service

import (
	"github.com/frtatmaca/ys-inmemory/core"
	"github.com/frtatmaca/ys-inmemory/models"
)

type FileService struct {
	repositoryAdapter core.RepositoryAdapter
}

func NewFileService(repositoryAdapter core.RepositoryAdapter) FileService {
	return FileService{repositoryAdapter}
}

func (p *FileService) ReadFile() map[string]string {
	repo, _ := p.repositoryAdapter.NewFileRepository(models.GOB_REPOSITORY)
	return repo.Read()
}

func (p *FileService) WriteFile(keyValStore map[string]string) bool {
	repo, _ := p.repositoryAdapter.NewFileRepository(models.GOB_REPOSITORY)
	return repo.WriteFile(keyValStore)
}
