package service

import (
	"context"
	"errors"
	"hse-2026-golang-project/internal/repository"
)

type IssueService struct {
	repo *repository.ProjectRepository
}

func NewIssueService(repo *repository.ProjectRepository) *IssueService {
	return &IssueService{repo: repo}
}

func (s *IssueService) GetByProjectKey(ctx context.Context, key string) (interface{}, error) {

	projects, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	for _, p := range projects {
		if p.Key == key {
			return s.repo.GetIssuesByProject(ctx, p.JiraID)
		}
	}

	return nil, errors.New("project not found")
}