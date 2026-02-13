package service_test

import (
	"Go-IssueTracker-API/internal/model"
	"Go-IssueTracker-API/internal/service"
	"context"
	"testing"
	"errors"
)

type MockRepo struct {
	CreateFunc     func(ctx context.Context, issue *model.Issue) (int, error)
	GetByIDFunc    func(ctx context.Context, id int) (*model.Issue, error)
	UpdateFunc     func(ctx context.Context, issue *model.Issue) error
	DeleteFunc     func(ctx context.Context, id int) error
	ListFunc       func(ctx context.Context) ([]*model.Issue, error)
}

func (m *MockRepo) CreateIssue(ctx context.Context, issue *model.Issue) (int, error) {
	return m.CreateFunc(ctx, issue)
}

func (m *MockRepo) GetIssueByID(ctx context.Context, id int) (*model.Issue, error) {
	return m.GetByIDFunc(ctx, id)
}

func (m *MockRepo) UpdateIssue(ctx context.Context, issue *model.Issue) error {
	return m.UpdateFunc(ctx, issue)
}

func (m *MockRepo) DeleteIssue(ctx context.Context, id int) error {
	return m.DeleteFunc(ctx, id)
}

func (m *MockRepo) ListIssues(ctx context.Context) ([]*model.Issue, error) {
	return m.ListFunc(ctx)
}

func TestCreateIssue(t *testing.T) {
	called := false
    mockRepo := &MockRepo{
        CreateFunc: func(ctx context.Context, issue *model.Issue) (int, error) {
			called = true
            return 1, nil
        },
    }

    service := service.NewIssueService(mockRepo)

	issue := &model.Issue{
		Title:       "Test Issue",
		Description: "This is a test issue",
		Status:      "open",
	}

	_, err := service.CreateIssue(context.Background(), issue)

    if err != nil {
        t.Fatalf("expected no error, got %v", err)
    }

	if !called {
		t.Fatal("expected CreateIssue to be called")
	}
}

func TestGetIssueByID(t *testing.T) {
	called := false
	mockRepo := &MockRepo{
		GetByIDFunc: func(ctx context.Context, id int) (*model.Issue, error) {
			called = true
			return &model.Issue{
				ID: id,
				Title: "test",	
			}, nil
		},
	}

	service := service.NewIssueService(mockRepo)

	issue, err := service.GetIssueByID(context.Background(), 1)

	if issue == nil {
		t.Fatalf("expected issue, got nil")
	}

	if issue.ID != 1 {
		t.Fatalf("expected ID 1, got %d", err)
	}

	if !called {
		t.Fatal("expected GetIssueByID to be called")
	}
}

func TestUpdateIssue(t *testing.T) {
	called := false
	mockRepo := &MockRepo{
		UpdateFunc: func(ctx context.Context, issue *model.Issue) (error) {
			called = true
			return nil
		},
	}

	service := service.NewIssueService(mockRepo)

	issue := &model.Issue{
		Title:       "Test Issue",
		Description: "This is a test issue",
		Status:      "open",
	}

	err := service.UpdateIssue(context.Background(), issue)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if !called {
		t.Fatal("expected UpdateIssue to be called")
	}
}

func TestUpdateIssue_InvalidStatus(t *testing.T) {
	called := false
	mockRepo := &MockRepo{
		UpdateFunc: func(ctx context.Context, issue *model.Issue) (error) {
			called = true
			return nil
		},
	}

	service := service.NewIssueService(mockRepo)

	issue := &model.Issue{
		Title:       "Test Issue",
		Description: "This is a test issue",
		Status:      "123",
	}

	err := service.UpdateIssue(context.Background(), issue)

	if err == nil {
		t.Fatalf("expected error, got nil")
	}

	if called {
		t.Fatal("expected UpdateIssue not to be called")
	}
}

func TestUpdateIssue_ReturnError(t *testing.T) {
	expErr := errors.New("invalid status")
	called := false
	mockRepo := &MockRepo{
		UpdateFunc: func(ctx context.Context, issue *model.Issue) (error) {
			called = true
			return expErr
		},
	}

	service := service.NewIssueService(mockRepo)
	issue := &model.Issue{
		Status: "open",
	}
	err := service.UpdateIssue(context.Background(), issue)

	if err == nil {
		t.Fatalf("expected error, got nil")
	}

	if err.Error() != expErr.Error() {
		t.Fatalf("expected %v, got %v", expErr, err)
	}

	if !called {
		t.Fatal("expected UpdateIssue to be called")
	}
}

func TestDeleteIssue(t *testing.T) {
	mockRepo := &MockRepo{
		DeleteFunc: func(ctx context.Context, id int) error {
			return nil
		},
	}

	service := service.NewIssueService(mockRepo)
	err := service.DeleteIssue(context.Background(), 1)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestListIssues(t *testing.T) {
	called := false
	mockRepo := &MockRepo{
		ListFunc: func(ctx context.Context) ([]*model.Issue, error) {
			called = true
			return []*model.Issue{
				{ID: 1, Title: "Test Issue 1"},
				{ID: 2, Title: "Test Issue 2"},
			}, nil
		},
	}

	service := service.NewIssueService(mockRepo)
	issues, err := service.ListIssues(context.Background())
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(issues) != 2 {
		t.Fatalf("expected 2 issues, got %d", len(issues))
	}

	if !called {
		t.Fatal("expected ListIssues to be called")
	}
}