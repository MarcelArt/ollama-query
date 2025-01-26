package services

import "github.com/MarcelArt/ollama-query/repositories"

type ITableService interface {
	GetTables() ([]string, error)
}

type TableService struct {
	repo repositories.ITableRepo
}

func NewTableService(repo repositories.ITableRepo) *TableService {
	return &TableService{
		repo: repo,
	}
}

func (s *TableService) GetTables() ([]string, error) {
	return s.repo.GetTables()
}
