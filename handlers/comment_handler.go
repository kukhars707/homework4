package handlers

import (
	"fmt"
	"github.com/kukhars707/homework4/services"
	"net/http"
)

type CommentHandler struct {
	service services.CommentService
}

func NewCommentService(service services.CommentService) CommentHandler {
	return CommentHandler{service: service}
}

func (h CommentHandler) GetComments(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Not implement")
}

func (h CommentHandler) GetComment(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Not implement")
}

func (h CommentHandler) CreateComments(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Not implement")
}

func (h CommentHandler) EditComments(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Not implement")
}

func (h CommentHandler) RemoveComments(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Not implement")
}