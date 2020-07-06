package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/kukhars707/homework4/services"
	"net/http"
)

type ProjectHandler struct {
	service services.ProjectService
}

func NewProjectHandler(service services.ProjectService) ProjectHandler {
	return ProjectHandler{service: service}
}

func (h ProjectHandler) GetProjects(w http.ResponseWriter, r *http.Request) {
	projects, err := h.service.GetProjects()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(projects)
	if err != nil {
		panic(err)
		return
	}
}

func (h ProjectHandler) GetProject(w http.ResponseWriter, r *http.Request) {
	projectID := mux.Vars(r)["id"]
	project, err := h.service.GetProject(projectID)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(project)
	if err != nil {
		panic(err)
		return
	}
}

func (h ProjectHandler) CreateProject(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	description := r.FormValue("description")

	project, err := h.service.CreateProject(name, description)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(project)
}

func (h ProjectHandler) EditProject(w http.ResponseWriter, r *http.Request) {
	projectID := mux.Vars(r)["id"]
	name := r.FormValue("name")
	description := r.FormValue("description")

	project, err := h.service.EditProject(projectID, name, description)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(project)
}

func (h ProjectHandler) RemoveProject(w http.ResponseWriter, r *http.Request) {
	projectID := mux.Vars(r)["id"]
	err := h.service.RemoveProject(projectID)

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
