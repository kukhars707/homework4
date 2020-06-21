package handlers

import (
	"fmt"
	"github.com/kukhars707/homework4/services"
	"net/http"
)

type ColumnHandler struct {
	service services.ColumnService
}

func NewColumnHandler(service services.ColumnService) ColumnHandler {
	return ColumnHandler{service: service}
}

func (h ColumnHandler) GetColumns(w http.ResponseWriter, r *http.Request) ColumnHandler {
	fmt.Println("Not Implement")
}

func (h ColumnHandler) GetColumn(w http.ResponseWriter, r *http.Request) ColumnHandler {
	fmt.Println("Not Implement")
}

func (h ColumnHandler) CreateColumns(w http.ResponseWriter, r *http.Request) ColumnHandler {
	fmt.Println("Not Implement")
}

func (h ColumnHandler) EditColumns(w http.ResponseWriter, r *http.Request) ColumnHandler {
	fmt.Println("Not Implement")
}

func (h ColumnHandler) RemoveColumns(w http.ResponseWriter, r *http.Request) ColumnHandler {
	fmt.Println("Not Implement")
}