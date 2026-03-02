package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/JingolBong/jira-connector/jira-backend/internal/service"
)

type ProjectHandler struct {
	service *service.ProjectService
}

func NewProjectHandler(s *service.ProjectService) *ProjectHandler {
	return &ProjectHandler{service: s}
}

func (h *ProjectHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	data, _ := h.service.GetAll(r.Context())
	json.NewEncoder(w).Encode(data)
}

func (h *ProjectHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	h.service.Delete(r.Context(), id)
	w.WriteHeader(http.StatusNoContent)
}