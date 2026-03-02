package service

import (
	"context"
	"github.com/JingolBong/jira-connector/jira-backend/internal/repository"
)

type GraphService struct {
	repo *repository.ProjectRepository
	analyzed map[string]bool
}

func NewGraphService(repo *repository.ProjectRepository) *GraphService {
	return &GraphService{
		repo: repo,
		analyzed: make(map[string]bool),
	}
}

func (s *GraphService) Make(ctx context.Context, projectKey string, task int) error {

	projects, _ := s.repo.GetAll(ctx)

	for _, p := range projects {
		if p.Key == projectKey {
			s.analyzed[projectKey] = true
			return nil
		}
	}

	return nil
}

func (s *GraphService) Get(ctx context.Context, projectKey string, task int) (map[string]int, error) {

	projects, _ := s.repo.GetAll(ctx)

	for _, p := range projects {
		if p.Key == projectKey {

			issues, _ := s.repo.GetIssuesByProject(ctx, p.JiraID)

			result := make(map[string]int)

			for _, i := range issues {
				result[i.Priority]++
			}

			return result, nil
		}
	}

	return nil, nil
}

func (s *GraphService) IsAnalyzed(projectKey string) bool {
	return s.analyzed[projectKey]
}