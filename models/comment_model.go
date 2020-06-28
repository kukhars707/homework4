package model

type Comment struct {
	ID   		string `json:"id"`
	Text 		string `json:"text"`
	ProjectID	string `json:"projectId"`
	ColumnID	string `json:"columnId"`
	TaskID		string `json:"taskId"`
}
