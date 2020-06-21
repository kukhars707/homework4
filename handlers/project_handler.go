package handlers

import (
	"fmt"
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
	fmt.Println("get project not implement")
}

func (h ProjectHandler) GetProject(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get project not implement")
}

func (h ProjectHandler) CreateProject(w http.ResponseWriter, r *http.Request) {
	fmt.Println("create project not implement")
}

func (h ProjectHandler) EditProject(w http.ResponseWriter, r *http.Request) {
	fmt.Println("edit project not implement")
}

func (h ProjectHandler) RemoveProject(w http.ResponseWriter, r *http.Request) {
	fmt.Println("delete project not implement")
}
