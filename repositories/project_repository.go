package repositories

import (
	"database/sql"
	model "github.com/kukhars707/homework4/models"
)

type ProjectRepository struct {
	db *sql.DB
}

func NewProjectRepository(db *sql.DB) ProjectRepository {
	return ProjectRepository{db: db}
}

func (r *ProjectRepository) GetProjects() (*[]model.Project, error) {
	var p model.Project
	selectStatement := `SELECT * from project;`
	result, err := r.db.Query(selectStatement)

	if err != nil {
		panic(err.Error())
	}

	var pSlice []model.Project
	for result.Next() {
		result.Scan(&p.ID, &p.Name, &p.Description)
		pSlice = append(pSlice, p)
	}

	return &pSlice, nil
}

func (r *ProjectRepository) GetProject(projectID string) (*model.Project, error) {
	selectStatement := `SELECT * FROM project WHERE id = $1;`
	result, err := r.db.Query(selectStatement, projectID)

	if err != nil {
		panic(err.Error())
	}

	var p model.Project

	for result.Next() {
		result.Scan(&p.ID, &p.Name, &p.Description)
	}

	return &p, nil
}

func (r *ProjectRepository) CreateProject(project *model.Project) error {
	insertStatement := `
		INSERT INTO project (name, description) 
		VALUES ($1, $2);`
	_, err := r.db.Exec(insertStatement, project.Name, project.Description)
	if err != nil {
		panic(err)
	}

	return err
}

func (r *ProjectRepository) EditProject(project *model.Project) error {
	sqlStatement := `UPDATE project SET name = $2, description = $3 WHERE id = $1;`

	_, err := r.db.Exec(sqlStatement, project.ID, project.Name, project.Description)

	if err != nil {
		panic(err.Error())
	}

	return err
}

func (r *ProjectRepository) RemoveProject(projectID string) error {
	sqlStatement := `DELETE FROM project WHERE id = $1;`

	_, err := r.db.Exec(sqlStatement, projectID)

	if err != nil {
		panic(err.Error())
	}

	return err
}