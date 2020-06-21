package repositories

import (
	"database/sql"
	"fmt"
	model "github.com/kukhars707/homework4/models"
)

type ColumnRepository struct {
	db *sql.DB
}

func NewColumnRepository(db *sql.DB) ColumnRepository {
	return ColumnRepository{db: db}
}

func (r *ColumnRepository) GetColumns() (*[]model.Column, error) {
	fmt.Println("Not implement")
}

func (r *ColumnRepository) GetColumn(columnID uint) (*model.Column, error) {
	fmt.Println("Not implement")
}

func (r *ColumnRepository) CreateColumn() error {
	fmt.Println("Not implement")
}

func (r *ColumnRepository) EditColumns(columnID uint) error {
	fmt.Println("Not implement")
}

func (r *ColumnRepository) RemoveColumns(columnID uint) error {
	fmt.Println("Not implement")
}