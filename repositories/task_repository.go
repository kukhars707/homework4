package repositories

import (
	"database/sql"
	"fmt"
	model "github.com/kukhars707/homework4/models"
)

type TaskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) TaskRepository {
	return TaskRepository{db: db}
}

func (r *TaskRepository) GetTasks() (*[]model.Task, error)  {
	fmt.Println("Not Implement")
}

func (r *TaskRepository) GetTask(taskID uint) (*model.Task, error)  {
	fmt.Println("Not Implement")
}

func (r *TaskRepository) CreateTasks() error  {
	fmt.Println("Not Implement")
}

func (r *TaskRepository) EditTasks(taskID uint) error  {
	fmt.Println("Not Implement")
}

func (r *TaskRepository) RemoveTasks(taskID uint) error  {
	fmt.Println("Not Implement")
}