package service

import (
    "context"

    "Go-IssueTracker-API/internal/model"
    "Go-IssueTracker-API/internal/repository"
)

type IssueService struct {
    repo repository.IssueRepository
}

/*
	1.CreateIssue(ctx context.Context, issue *model.Issue) (int, error)
	2.GetIssueByID(ctx context.Context, id int) (*model.Issue, error)
	3.UpdateIssue(ctx context.Context, issue *model.Issue) error
	4.DeleteIssue(ctx context.Context, id int) error
	5.ListIssues(ctx context.Context) ([]*model.Issue, error)
*/

func NewIssueService(repo repository.IssueRepository) *IssueService {
    return &IssueService{repo: repo}
}

func (s *IssueService) CreateIssue(ctx context.Context, issue *model.Issue) (int, error) {

}

func (s *IssueService) GetIssueByID(ctx context.Context, id int) (*model.Issue, error) {

}

func (s *IssueService) UpdateIssue(ctx context.Context, issue *model.Issue) error {

}

func (s *IssueService) DeleteIssue(ctx context.Context, id int) error {

}

func (s *IssueService) ListIssues(ctx context.Context) ([]*model.Issue, error) {
	
}
