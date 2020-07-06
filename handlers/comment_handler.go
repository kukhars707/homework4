package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/kukhars707/homework4/services"
	"net/http"
)

type CommentHandler struct {
	service services.CommentService
}

func NewCommentHandler(service services.CommentService) CommentHandler {
	return CommentHandler{service: service}
}

func (h CommentHandler) GetComments(w http.ResponseWriter, r *http.Request) {
	taskID := mux.Vars(r)["taskId"]

	comments, err := h.service.GetComments(taskID)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(comments)
	if err != nil {
		panic(err)
		return
	}
}

func (h CommentHandler) CreateComments(w http.ResponseWriter, r *http.Request) {
	taskID := mux.Vars(r)["taskId"]
	text := r.FormValue("text")

	comment, err := h.service.CreateComment(taskID, text)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(comment)
}

func (h CommentHandler) EditComments(w http.ResponseWriter, r *http.Request) {
	commentID := mux.Vars(r)["commentId"]
	text := r.FormValue("text")

	comment, err := h.service.CreateComment(commentID, text)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(comment)
}

func (h CommentHandler) RemoveComments(w http.ResponseWriter, r *http.Request) {
	commentID := mux.Vars(r)["commentId"]

	err := h.service.RemoveComments(commentID)

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