package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/JingolBong/jira-connector/jira-backend/internal/service"
)

type GraphHandler struct {
	service *service.GraphService
}

func NewGraphHandler(s *service.GraphService) *GraphHandler {
	return &GraphHandler{service: s}
}

func (h *GraphHandler) Make(w http.ResponseWriter, r *http.Request) {

	task, _ := strconv.Atoi(mux.Vars(r)["task"])
	project := r.URL.Query().Get("project")

	h.service.Make(r.Context(), project, task)
	w.WriteHeader(http.StatusOK)
}

func (h *GraphHandler) Get(w http.ResponseWriter, r *http.Request) {

	task, _ := strconv.Atoi(mux.Vars(r)["task"])
	project := r.URL.Query().Get("project")

	data, _ := h.service.Get(r.Context(), project, task)
	json.NewEncoder(w).Encode(data)
}

func (h *GraphHandler) IsAnalyzed(w http.ResponseWriter, r *http.Request) {

	project := r.URL.Query().Get("project")
	json.NewEncoder(w).Encode(h.service.IsAnalyzed(project))
}