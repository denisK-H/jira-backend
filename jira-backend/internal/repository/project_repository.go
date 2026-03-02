package repository

import (
	"context"
	"database/sql"
	"github.com/JingolBong/jira-connector/internal/db"
	"github.com/JingolBong/jira-connector/internal/models"
)

type ProjectRepository struct {
	storage *db.Storage
	writeDB *sql.DB
}

func NewProjectRepository(storage *db.Storage, writeDB *sql.DB) *ProjectRepository {
	return &ProjectRepository{storage: storage, writeDB: writeDB}
}

func (r *ProjectRepository) GetAll(ctx context.Context) ([]models.Project, error) {
	return r.storage.GetAllProjects(ctx)
}

func (r *ProjectRepository) GetByID(ctx context.Context, id int64) (*models.Project, error) {
	return r.storage.GetProjectByJiraID(ctx, id)
}

func (r *ProjectRepository) Delete(ctx context.Context, id int64) error {
	_, err := r.writeDB.ExecContext(ctx, "DELETE FROM project WHERE jira_id=$1", id)
	return err
}

func (r *ProjectRepository) GetIssuesByProject(ctx context.Context, id int64) ([]models.Issue, error) {
	return r.storage.GetIssuesByProject(ctx, id)
}