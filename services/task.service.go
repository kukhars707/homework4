package services

import (
	model "github.com/kukhars707/homework4/models"
	"github.com/kukhars707/homework4/repositories"
)

type TaskService struct {
	repository repositories.TaskRepository
}

func NewTaskService(repository repositories.TaskRepository) TaskService {
	return TaskService{repository: repository}
}

func (s *TaskService) GetTasks(projectID string) (*[]model.Task, error) {
	return s.repository.GetTasks(projectID)
}

func (s *TaskService) GetTask(taskID string) (*model.Task, error) {
	return s.repository.GetTask(taskID)
}

func (s *TaskService) CreateTask(projectID string, columnID string, name string, description string) (*model.Task, error) {
	task := &model.Task{
		Name: name,
		Description: description,
		ProjectID: projectID,
		ColumnID: columnID,
	}

	if err := s.repository.CreateTask(task); err != nil {
		return task, err
	}

	return task, nil
}

func (s *TaskService) EditTask(projectID string, columnID string, taskId string, name string, description string) (*model.Task, error) {
	task := &model.Task{
		ID: taskId,
		Name: name,
		Description: description,
		ProjectID: projectID,
		ColumnID: columnID,
	}

	if err := s.repository.EditTask(task); err != nil {
		return task, err
	}

	return task, nil
}

func (s *TaskService) RemoveTask(taskId string) error {
	return s.repository.RemoveTask(taskId)
}