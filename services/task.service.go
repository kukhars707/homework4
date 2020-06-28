package services

import (
	"fmt"
	model "github.com/kukhars707/homework4/models"
	"github.com/kukhars707/homework4/repositories"
)

type TaskService struct {
	repository repositories.TaskRepository
}

func NewTaskService(repository repositories.TaskRepository) TaskService {
	return TaskService{repository: repository}
}

func (s *TaskService) GetTasks(projectID string, columnID string) (*[]model.Task, error) {
	return s.repository.GetTasks(projectID, columnID)
}

func (s *TaskService) GetTask(projectID string, columnID string, taskID string) (*model.Task, error) {
	return s.repository.GetTask(projectID, columnID, taskID)
}

func (s *TaskService) CreateTask(projectID string, columnID string, name string, description string) (*model.Task, error) {
	task := &model.Task{
		Name: name,
		Description: description,
		ProjectID: projectID,
		ColumnID: columnID,
	}

	if err := s.repository.CreateTasks(task); err != nil {
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

	if err := s.repository.EditTasks(task); err != nil {
		return task, err
	}

	return task, nil
}

func (s *TaskService) RemoveTasks(projectID string, columnID string, taskId string) (*model.Task, error) {
	return s.repository.RemoveTasks(projectID, columnID, taskId)
}