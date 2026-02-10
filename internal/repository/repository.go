package repository

import "Go-IssueTracker-API/internal/model"

type issueRepository interface {
	CreateIssue(title string, description string) (int, error)
	GetIssueByID(id int) (*model.Issue, error)
	UpdateIssue(id int, title string, description string, status string) error
	DeleteIssue(id int) error
	ListIssues() ([]*model.Issue, error)
}