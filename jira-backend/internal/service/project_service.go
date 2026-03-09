package service

import (
	"context"

	"hse-2026-golang-project/internal/models"
	"hse-2026-golang-project/jira-backend/internal/repository"
)

type ProjectService struct {
	repo *repository.ProjectRepository
}

func NewProjectService(repo *repository.ProjectRepository) *ProjectService {
	return &ProjectService{repo: repo}
}

func (s *ProjectService) GetAll(ctx context.Context) ([]models.Project, error) {
	return s.repo.GetAll(ctx)
}

func (s *ProjectService) Delete(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}
