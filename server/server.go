package server

import (
	"database/sql"
	"fmt"
	"github.com/kukhars707/homework4/repositories"
	"github.com/kukhars707/homework4/server/router"
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

	projectRepository := repositories.NewProjectRepository(db)
}