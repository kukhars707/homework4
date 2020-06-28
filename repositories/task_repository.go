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

func (r *TaskRepository) GetTasks(projectID string, columnID string) (*[]model.Task, error)  {
	fmt.Println("Not Implement")
}

func (r *TaskRepository) GetTask(projectID string, columnID string, taskID string) (*model.Task, error)  {
	fmt.Println("Not Implement")
}

func (r *TaskRepository) CreateTasks(task *model.Task) error  {
	fmt.Println("Not Implement")
}

func (r *TaskRepository) EditTasks(task *model.Task) error  {
	fmt.Println("Not Implement")
}

func (r *TaskRepository) RemoveTasks(projectID string, columnID string, taskId string) error  {
	fmt.Println("Not Implement")
}