package router

import (
	"github.com/gorilla/mux"
	"github.com/kukhars707/homework4/handlers"
)

type Router struct {
	Route *mux.Router
}

func NewRouter() Router {
	return Router{Route: mux.NewRouter()}
}

func (r *Router) CreateRoutes(projectHandler handlers.ProjectHandler) {
	r.Route.HandleFunc("/project", projectHandler.GetProjects).Methods("GET")
	r.Route.HandleFunc("/project/{id}", projectHandler.GetProject).Methods("GET")
	r.Route.HandleFunc("/project", projectHandler.CreateProject).Methods("POST")
	r.Route.HandleFunc("/project/{id}", projectHandler.EditProject).Methods("PUT")
	r.Route.HandleFunc("/project/{id}", projectHandler.RemoveProject).Methods("DELETE")
}