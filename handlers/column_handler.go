package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/kukhars707/homework4/services"
	"net/http"
)

type ColumnHandler struct {
	service services.ColumnService
}

func NewColumnHandler(service services.ColumnService) ColumnHandler {
	return ColumnHandler{service: service}
}

func (h ColumnHandler) GetColumns(w http.ResponseWriter, r *http.Request) {
	projectID := mux.Vars(r)["projectId"]
	columns, err := h.service.GetColumns(projectID)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(columns)
	if err != nil {
		panic(err)
		return
	}
}

func (h ColumnHandler) GetColumn(w http.ResponseWriter, r *http.Request) {
	projectID := mux.Vars(r)["projectId"]
	columnID := mux.Vars(r)["columnId"]

	column, err := h.service.GetColumn(projectID, columnID)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(column)
	if err != nil {
		panic(err)
		return
	}
}

func (h ColumnHandler) CreateColumn(w http.ResponseWriter, r *http.Request) {
	projectID := mux.Vars(r)["projectId"]
	name := r.FormValue("name")

	column, err := h.service.CreateColumn(projectID, name)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(column)
}

func (h ColumnHandler) EditColumn(w http.ResponseWriter, r *http.Request) {
	projectID := mux.Vars(r)["projectId"]
	columnID := mux.Vars(r)["columnId"]
	name := r.FormValue("name")

	column, err := h.service.EditColumn(projectID, columnID, name)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(column)
}

func (h ColumnHandler) RemoveColumn(w http.ResponseWriter, r *http.Request) {
	projectID := mux.Vars(r)["projectId"]
	columnID := mux.Vars(r)["columnId"]

	column, err := h.service.RemoveColumn(projectID, columnID)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(column)
	if err != nil {
		panic(err)
		return
	}
}