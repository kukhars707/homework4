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

func (s *ColumnService) GetColumns() ([]model.Column, error) {
	fmt.Println("Not implement")
}

func (s *ColumnService) GetColumn(columnID uint) (model.Column, error) {
	fmt.Println("Not implement")
}

func (s *ColumnService) CreateColumns() error {
	fmt.Println("Not implement")
}

func (s *ColumnService) EditColumns(columnID uint) error {
	fmt.Println("Not implement")
}

func (s *ColumnService) RemoveColumns(columnID uint) error {
	fmt.Println("Not implement")
}