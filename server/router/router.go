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

func (r *Router) CreateRoutes(
	projectHandler handlers.ProjectHandler,
	columnHandler handlers.ColumnHandler,
	taskHandler handlers.TaskHandler,
	commentHandler handlers.CommentHandler) {
	r.Route.HandleFunc("/project", projectHandler.GetProjects).Methods("GET")
	r.Route.HandleFunc("/project/{id}", projectHandler.GetProject).Methods("GET")
	r.Route.HandleFunc("/project", projectHandler.CreateProject).Methods("POST")
	r.Route.HandleFunc("/project/{id}", projectHandler.EditProject).Methods("PUT")
	r.Route.HandleFunc("/project/{id}", projectHandler.RemoveProject).Methods("DELETE")

	r.Route.HandleFunc("/{projectId}/column", columnHandler.GetColumns).Methods("GET")
	r.Route.HandleFunc("/{projectId}/column", columnHandler.CreateColumn).Methods("POST")
	r.Route.HandleFunc("/{projectId}/{columnId}", columnHandler.EditColumn).Methods("PUT")
	r.Route.HandleFunc("/{columnId}", columnHandler.RemoveColumn).Methods("DELETE")

	r.Route.HandleFunc("/{projectId}/tasks", taskHandler.GetTasks).Methods("GET")
	r.Route.HandleFunc("/task/{taskId}", taskHandler.GetTask).Methods("GET")
	r.Route.HandleFunc("/{projectId}/{columnId}/task", taskHandler.CreateTasks).Methods("POST")
	r.Route.HandleFunc("/{projectId}/{columnId}/{taskId}", taskHandler.EditTasks).Methods("PUT")
	r.Route.HandleFunc("/task/{taskId}", taskHandler.RemoveTasks).Methods("DELETE")

	r.Route.HandleFunc("/{taskId}/comment", commentHandler.GetComments).Methods("GET")
	r.Route.HandleFunc("/{taskId}/comment", commentHandler.CreateComments).Methods("POST")
	r.Route.HandleFunc("/comment/{commentId}", commentHandler.EditComments).Methods("PUT")
	r.Route.HandleFunc("/{commentId}", commentHandler.RemoveComments).Methods("DELETE")
}