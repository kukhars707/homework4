package services

import (
	"fmt"
	model "github.com/kukhars707/homework4/models"
	"github.com/kukhars707/homework4/repositories"
)

type ProjectService struct {
	repository repositories.ProjectRepository
}

func NewProjectService(repository repositories.ProjectRepository) ProjectService {
	return ProjectService{repository: repository}
}

func (s *ProjectService) GetProjects() ([]model.Project, error) {
	fmt.Println("Not implement")
}

func (s *ProjectService) GetProject(projectID uint) (model.Project, error) {
	fmt.Println("Not implement")
}

func (s *ProjectService) CreateProject() error {
	fmt.Println("Not implement")
}

func (s *ProjectService) EditProject(projectID uint) error {
	fmt.Println("Not implement")
}

func (s *ProjectService) RemoveProject(projectID uint) error {
	fmt.Println("Not implement")
}
