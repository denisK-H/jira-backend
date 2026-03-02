package handler

import (
	"encoding/json"
	"net/http"

	"github.com/JingolBong/jira-connector/jira-backend/internal/service"
)

type IssueHandler struct {
	service *service.IssueService
}

func NewIssueHandler(s *service.IssueService) *IssueHandler {
	return &IssueHandler{service: s}
}

func (h *IssueHandler) GetByProject(w http.ResponseWriter, r *http.Request) {

	key := r.URL.Query().Get("project")
	data, _ := h.service.GetByProjectKey(r.Context(), key)

	json.NewEncoder(w).Encode(data)
}