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

func (s *TaskService) GetTasks() ([]model.Task, error) {
	fmt.Println("Not implement")
}

func (s *TaskService) GetTask(taskID uint) (model.Task, error) {
	fmt.Println("Not implement")
}

func (s *TaskService) CreateTasks() error {
	fmt.Println("Not implement")
}

func (s *TaskService) EditTask(taskID uint) error {
	fmt.Println("Not implement")
}

func (s *TaskService) RemoveTasks() error {
	fmt.Println("Not implement")
}