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

func (r *ColumnRepository) GetColumns(projectID string) (*[]model.Column, error) {
	column := &[]model.Column{{ID: "1", Name: "name"}, {ID: "2", Name: "name2"}}

	return column, nil
}

func (r *ColumnRepository) GetColumn(projectID string, columnID string) (*model.Column, error) {
	fmt.Println(columnID)
	column := &model.Column{ID: "1", Name: "name"}

	return column, nil
}

func (r *ColumnRepository) CreateColumn(column *model.Column) error {
	fmt.Println("Not implement")
}

func (r *ColumnRepository) EditColumn(column *model.Column) error {
	fmt.Println("Not implement")
}

func (r *ColumnRepository) RemoveColumn(projectID string, columnID string) error {
	fmt.Println("Not implement")
}