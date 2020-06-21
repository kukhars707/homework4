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
	fmt.Println("Not implement")
}

func (r *ProjectRepository) GetProject(projectID uint) (*model.Project, error) {
	fmt.Println("Not implement")
}

func (r *ProjectRepository) CreateProject() error {
	fmt.Println("Not implement")
}

func (r *ProjectRepository) EditProjects(projectID uint) error {
	fmt.Println("Not implement")
}

func (r *ProjectRepository) RemoveProjects(projectID uint) error {
	fmt.Println("Not implement")
}