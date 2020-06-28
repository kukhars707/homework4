package services

import (
	"fmt"
	model "github.com/kukhars707/homework4/models"
	"github.com/kukhars707/homework4/repositories"
)

type ColumnService struct {
	repository repositories.ColumnRepository
}

func NewColumnService(repository repositories.ColumnRepository) ColumnService {
	return ColumnService{repository: repository}
}

func (s *ColumnService) GetColumns(projectID string) (*[]model.Column, error) {
	return s.repository.GetColumns(projectID)
}

func (s *ColumnService) GetColumn(projectID string, columnID string) (*model.Column, error) {
	return s.repository.GetColumn(projectID, columnID)
}

func (s *ColumnService) CreateColumn(projectID string, name string) (*model.Column, error) {
	column := &model.Column{
		ProjectID: projectID,
		Name: name,
	}

	if err := s.repository.CreateColumn(column); err != nil {
		return column, err
	}

	return column, nil
}

func (s *ColumnService) EditColumn(projectID string, columnID string, name string) (*model.Column, error) {
	column := &model.Column{
		ID: columnID,
		ProjectID: projectID,
		Name: name,
	}

	if err := s.repository.EditColumn(column); err != nil {
		return column, err
	}

	return column, nil
}

func (s *ColumnService) RemoveColumn(projectID string, columnID string) (*model.Column, error) {
	return s.repository.RemoveColumn(projectID, columnID)
}