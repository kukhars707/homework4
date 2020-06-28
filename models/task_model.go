package model

type Task struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ProjectID 	string `json:"projectId"`
	ColumnID	string `json:"columnId"`
}
