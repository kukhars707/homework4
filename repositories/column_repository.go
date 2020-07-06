package repositories

import (
	"database/sql"
	model "github.com/kukhars707/homework4/models"
)

type ColumnRepository struct {
	db *sql.DB
}

func NewColumnRepository(db *sql.DB) ColumnRepository {
	return ColumnRepository{db: db}
}

func (r *ColumnRepository) GetColumns(projectId string) (*[]model.Column, error) {
	var c model.Column
	sqlStatement := `SELECT * FROM column WHERE projectID = $1;`
	result, err := r.db.Query(sqlStatement, projectId)

	if err != nil {
		panic(err.Error())
	}

	var cSlice []model.Column

	for result.Next() {
		result.Scan(&c.ID, &c.Name)

		cSlice = append(cSlice, c)
	}

	return &cSlice, nil
}

func (r *ColumnRepository) CreateColumn(column *model.Column) error {
	sqlStatement := `INSERT INTO column (projectID, name) VALUES ($1, $2);`

	_, err := r.db.Exec(sqlStatement, column.ProjectID, column.Name)

	if err != nil {
		panic(err.Error())
	}

	return err
}

func (r *ColumnRepository) EditColumn(column *model.Column) error {
	sqlStatement := `UPDATE column SET name = $2, projectID = $3 WHERE id = $1;`

	_, err := r.db.Exec(sqlStatement, column.ID, column.Name, column.ProjectID)

	if err != nil {
		panic(err.Error())
	}

	return err
}

func (r *ColumnRepository) RemoveColumn(columnID string) error {
	sqlStatement := `DELETE column WHERE id = $1;`

	_, err := r.db.Exec(sqlStatement, columnID)

	if err != nil {
		panic(err.Error())
	}

	return err
}