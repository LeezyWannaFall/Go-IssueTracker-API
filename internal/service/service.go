package service

import (
    "context"
	"errors"
    "Go-IssueTracker-API/internal/model"
)

type IssueService struct {
    repo IssueRepository
}

/*
	1.CreateIssue(ctx context.Context, issue *model.Issue) (int, error)
	2.GetIssueByID(ctx context.Context, id int) (*model.Issue, error)
	3.UpdateIssue(ctx context.Context, issue *model.Issue) error
	4.DeleteIssue(ctx context.Context, id int) error
	5.ListIssues(ctx context.Context) ([]*model.Issue, error)
*/

func NewIssueService(repo IssueRepository) *IssueService {
    return &IssueService{repo: repo}
}

func (s *IssueService) CreateIssue(ctx context.Context, issue *model.Issue) (int, error) {
	if issue.Title == "" {
		return 0, errors.New("title is required")
	}

	issue.Status = "open"

	return s.repo.CreateIssue(ctx, issue)
}

func (s *IssueService) GetIssueByID(ctx context.Context, id int) (*model.Issue, error) {
	return s.repo.GetIssueByID(ctx, id)
}

func (s *IssueService) UpdateIssue(ctx context.Context, issue *model.Issue) error {
	if issue.Status != "open" &&
	   issue.Status != "in_progress" &&
	   issue.Status != "done" {
		return errors.New("invalid status")
	}
	
	return s.repo.UpdateIssue(ctx, issue)
}

func (s *IssueService) DeleteIssue(ctx context.Context, id int) error {
	return s.repo.DeleteIssue(ctx, id)
}

func (s *IssueService) ListIssues(ctx context.Context) ([]*model.Issue, error) {
	return s.repo.ListIssues(ctx)
}
