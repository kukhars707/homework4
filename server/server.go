package server

import (
	"database/sql"
	"fmt"
	"github.com/kukhars707/homework4/handlers"
	"github.com/kukhars707/homework4/repositories"
	"github.com/kukhars707/homework4/server/router"
	"github.com/kukhars707/homework4/services"
	"net/http"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "sergey"
	dbname   = "taskmanager"
)

type Server struct {
	routes router.Router
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Start() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+ "dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	projectRepository := repositories.NewProjectRepository(db)
	columnRepository := repositories.NewColumnRepository(db)
	taskRepository := repositories.NewTaskRepository(db)
	commentRepository := repositories.NewCommentRepository(db)

	projectService := services.NewProjectService(projectRepository)
	columnService := services.NewColumnService(columnRepository)
	taskService := services.NewTaskService(taskRepository)
	commentService := services.NewCommentService(commentRepository)

	projectHandler := handlers.NewProjectHandler(projectService)
	columnHandler := handlers.NewColumnHandler(columnService)
	taskHandler := handlers.NewTaskHandler(taskService)
	commentHandler := handlers.NewCommentHandler(commentService)

	s.routes = router.NewRouter()
	s.routes.CreateRoutes(projectHandler, columnHandler, taskHandler, commentHandler)

	serve := &http.Server{
		Addr: ":8000",
		Handler: s.routes.Route,
	}

	serve.ListenAndServe()
}