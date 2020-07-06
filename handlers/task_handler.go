package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/kukhars707/homework4/services"
	"net/http"
)

type TaskHandler struct {
	service services.TaskService
}

func NewTaskHandler(service services.TaskService) TaskHandler {
	return TaskHandler{service: service}
}

func (h TaskHandler) GetTasks(w http.ResponseWriter, r *http.Request)  {
	projectID := mux.Vars(r)["projectId"]

	tasks, err := h.service.GetTasks(projectID)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(tasks)
	if err != nil {
		panic(err)
		return
	}
}

func (h TaskHandler) GetTask(w http.ResponseWriter, r *http.Request)  {
	taskID := mux.Vars(r)["taskId"]

	task, err := h.service.GetTask(taskID)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(task)
	if err != nil {
		panic(err)
		return
	}
}

func (h TaskHandler) CreateTasks(w http.ResponseWriter, r *http.Request)  {
	projectID := mux.Vars(r)["projectId"]
	columnID := mux.Vars(r)["columnId"]
	name := r.FormValue("name")
	description := r.FormValue("description")

	task, err := h.service.CreateTask(projectID, columnID, name, description)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(task)
}

func (h TaskHandler) EditTasks(w http.ResponseWriter, r *http.Request)  {
	projectID := mux.Vars(r)["projectId"]
	columnID := mux.Vars(r)["columnId"]
	taskID := mux.Vars(r)["taskId"]
	name := r.FormValue("name")
	description := r.FormValue("description")

	task, err := h.service.EditTask(projectID, columnID, taskID, name, description)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(task)
}

func (h TaskHandler) RemoveTasks(w http.ResponseWriter, r *http.Request)  {
	taskID := mux.Vars(r)["taskId"]

	err := h.service.RemoveTask(taskID)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	if err != nil {
		panic(err)
		return
	}
}
