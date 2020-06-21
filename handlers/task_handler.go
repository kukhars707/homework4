package handlers

import (
	"fmt"
	"github.com/kukhars707/homework4/services"
	"net/http"
)

type TaskHandler struct {
	service services.TaskService
}

func (h TaskHandler) NewTaskHandler(service services.TaskService) TaskHandler {
	return TaskHandler{service: service}
}

func (h TaskHandler) GetTasks(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("Not implement")
}

func (h TaskHandler) GetTask(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("Not implement")
}

func (h TaskHandler) CreateTasks(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("Not implement")
}

func (h TaskHandler) EditTasks(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("Not implement")
}

func (h TaskHandler) RemoveTasks(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("Not implement")
}
