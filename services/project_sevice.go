package services

import (
	model "github.com/kukhars707/homework4/models"
	"github.com/kukhars707/homework4/repositories"
)

type ProjectService struct {
	repository repositories.ProjectRepository
}

func NewProjectService(repository repositories.ProjectRepository) ProjectService {
	return ProjectService{repository: repository}
}

func (s *ProjectService) GetProjects() (*[]model.Project, error) {
	return s.repository.GetProjects()
}

func (s *ProjectService) GetProject(projectID string) (*model.Project, error) {
	return s.repository.GetProject(projectID)
}

func (s *ProjectService) CreateProject(name string, description string) (*model.Project, error) {
	project := &model.Project{
		Name: name,
		Description: description,
	}

	if err := s.repository.CreateProject(project); err != nil {
		return project, err
	}

	return project, nil
}

func (s *ProjectService) EditProject(projectID string, name string, description string) (*model.Project, error) {
	project := &model.Project{
		ID: projectID,
		Name: name,
		Description: description,
	}

	if err := s.repository.EditProject(project); err != nil {
		return project, err
	}

	return project, nil
}

func (s *ProjectService) RemoveProject(projectID string) (*model.Project, error) {
	return s.repository.RemoveProject(projectID)
}
