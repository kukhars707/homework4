package repositories

import (
	"database/sql"
	"fmt"
	model "github.com/kukhars707/homework4/models"
)

type ProjectRepository struct {
	db *sql.DB
}

func NewProjectRepository(db *sql.DB) ProjectRepository {
	return ProjectRepository{db: db}
}

func (r *ProjectRepository) GetProjects() (*[]model.Project, error) {
	project := &[]model.Project{{ID: "1", Name: "test", Description: "desc"}, {ID: "2", Name: "test 2", Description: "desc 2"}}

	return project, nil
}

func (r *ProjectRepository) GetProject(projectID string) (*model.Project, error) {
	fmt.Println(projectID)
	project := &model.Project{
		ID: "1", Name: "test", Description: "desc",
	}

	return project, nil
}

func (r *ProjectRepository) CreateProject(project *model.Project) error {
	return "err"
}

func (r *ProjectRepository) EditProject(project *model.Project) error {
	return "err"
}

func (r *ProjectRepository) RemoveProject(projectID string) error {
	return "err"
}