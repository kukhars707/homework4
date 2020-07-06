package repositories

import (
	"database/sql"
	model "github.com/kukhars707/homework4/models"
)

type TaskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) TaskRepository {
	return TaskRepository{db: db}
}

func (r *TaskRepository) GetTasks(projectID string) (*[]model.Task, error)  {
	var t model.Task
	sqlStatement := `SELECT * FROM task WHERE projectID = $1`
	result, err := r.db.Query(sqlStatement, projectID)

	if err != nil {
		panic(err.Error())
	}

	var tSlice []model.Task
	for result.Next() {
		result.Scan(&t.ID, &t.Name, &t.Description, &t.ProjectID, &t.ColumnID)

		tSlice = append(tSlice, t)
	}

	return &tSlice, nil

}

func (r *TaskRepository) GetTask(taskID string) (*model.Task, error)  {
	selectStatement := `SELECT * FROM task WHERE id = $1;`
	result, err := r.db.Query(selectStatement, taskID)

	if err != nil {
		panic(err.Error())
	}

	var t model.Task

	for result.Next() {
		result.Scan(&t.ID, &t.Name, &t.Description, &t.ColumnID, &t.ProjectID)
	}

	return &t, nil
}

func (r *TaskRepository) CreateTask(task *model.Task) error  {
	sqlStatement := `INSERT INTO task (name, description, projectID, columnID) VALUES ($1, $2, $3, $4);`

	_, err := r.db.Exec(sqlStatement, task.Name, task.Description, task.ProjectID, task.ColumnID)

	if err != nil {
		panic(err.Error())
	}

	return err
}

func (r *TaskRepository) EditTask(task *model.Task) error  {
	sqlStatement := `UPDATE task SET name = $1, description = $2, projectID = $3, columnID = $4 WHERE id = $5;`

	_, err := r.db.Exec(sqlStatement, task.Name, task.Description, task.ProjectID, task.ColumnID, task.ID)

	if err != nil {
		panic(err.Error())
	}

	return err
}

func (r *TaskRepository) RemoveTask(taskId string) error  {
	sqlStatement := `DELETE FROM task WHERE id = $1;`

	_, err := r.db.Exec(sqlStatement, taskId)

	if err != nil {
		panic(err.Error())
	}

	return err
}