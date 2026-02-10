package repository

import (
	"Go-IssueTracker-API/internal/model"
	"context"
)

type IssueRepository interface {
	CreateIssue(ctx context.Context, issue *model.Issue) (int, error)
	GetIssueByID(ctx context.Context, id int) (*model.Issue, error)
	UpdateIssue(ctx context.Context, issue *model.Issue) error
	DeleteIssue(ctx context.Context, id int) error
	ListIssues(ctx context.Context) ([]*model.Issue, error)
}