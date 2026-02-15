package handler_test

import (
	"Go-IssueTracker-API/internal/handler"
	"Go-IssueTracker-API/internal/model"
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/go-chi/chi/v5"
)

type MockService struct {
	CreateFunc  func(ctx context.Context, issue *model.Issue) (int, error)
	GetByIDFunc func(ctx context.Context, id int) (*model.Issue, error)
	UpdateFunc  func(ctx context.Context, issue *model.Issue) error
	DeleteFunc  func(ctx context.Context, id int) error
	ListFunc    func(ctx context.Context) ([]*model.Issue, error)
}

func (m *MockService) CreateIssue(ctx context.Context, issue *model.Issue) (int, error) {
	return m.CreateFunc(ctx, issue)
}

func (m *MockService) GetIssueByID(ctx context.Context, id int) (*model.Issue, error) {
	return m.GetByIDFunc(ctx, id)
}

func (m *MockService) UpdateIssue(ctx context.Context, issue *model.Issue) error {
	return m.UpdateFunc(ctx, issue)
}

func (m *MockService) DeleteIssue(ctx context.Context, id int) error {
	return m.DeleteFunc(ctx, id)
}

func (m *MockService) ListIssues(ctx context.Context) ([]*model.Issue, error) {
	return m.ListFunc(ctx)
}

func TestCreateIssue(t *testing.T) {
	called := false
	mockService := &MockService{
		CreateFunc: func(ctx context.Context, issue *model.Issue) (int, error) {
			called = true
			return 1, nil
		},
	}

	h := handler.NewHandler(mockService)

	body := &model.Issue{
		Title:       "Test",
		Description: "Desc",
		Status:      "open",
	}

	jsonBody, _ := json.Marshal(body)

	req := httptest.NewRequest(http.MethodPost, "/issues", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	res := httptest.NewRecorder()

	h.CreateIssue(res, req)

	// Проверяем статус
	if res.Code != http.StatusCreated {
		t.Fatalf("expected status 201, got %d", res.Code)
	}

	// Проверяем, что сервис вызвался
	if !called {
		t.Fatal("expected CreateIssue to be called")
	}

	// Проверяем тело ответа
	var response map[string]int
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		t.Fatalf("cannot decode response: %v", err)
	}

	if response["id"] != 1 {
		t.Fatalf("expected id 1, got %d", response["id"])
	}
}

func TestGetByID(t *testing.T) {
	called := false

	mockService := &MockService{
		GetByIDFunc: func(ctx context.Context, id int) (*model.Issue, error) {
			called = true
			return &model.Issue{
				ID: 1,
				Title: "title test",
				Description: "desc test",
				Status: "open",
			}, nil
		},
	}

	h := handler.NewHandler(mockService)
	r := chi.NewRouter()
	r.Get("/issue/{id}", h.GetIssueByID)

	req := httptest.NewRequest(http.MethodGet, "/issue/1", nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	t.Log("status:", res.Code)

	if !called {
		t.Fatal("expected CreateIssue to be called")
	}

	if res.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", res.Code)
	}

	var response model.Issue

	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		t.Fatalf("cannot decode response: %v", err)
	}

	if response.ID != 1 {
		t.Fatalf("expected valid issue, got %v", response)
	}
}

func TestUpdateIssue(t *testing.T) {
	called := false

	mockService := &MockService{
		UpdateFunc: func(ctx context.Context, issue *model.Issue) error {
			called = true
			return nil
		},
	}

	h := handler.NewHandler(mockService)
	r := chi.NewRouter()
	r.Put("/issue/{id}", h.UpdateIssue)

	body := &model.Issue{
		ID: 1,
		Title: "title test",
		Description: "desc test",
		Status: "open",
	}

	jsonBody, _ := json.Marshal(body)

	req := httptest.NewRequest(http.MethodPut, "/issue/1", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	t.Log("status:", res.Code)

	if !called {
		t.Fatal("expected UpdateIssue to be called")
	}

	if res.Code != http.StatusNoContent {
		t.Fatalf("expected status 200, got %d", res.Code)
	}
}

func TestDeleteIssue(t *testing.T) {
	called := false

	mockService := &MockService{
		DeleteFunc: func(ctx context.Context, id int) error {
			called = true
			return nil
		},
	}

	h := handler.NewHandler(mockService)
	r := chi.NewRouter()
	r.Delete("/issue/{id}", h.DeleteIssue)

	req := httptest.NewRequest(http.MethodDelete, "/issue/1", nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	t.Log("status:", res.Code)

	if !called {
		t.Fatal("expected DeleteIssue to be called")
	}

	if res.Code != http.StatusNoContent {
		t.Fatalf("expected status 204, got %d", res.Code)
	}
}

func TestListIssues(t *testing.T) {
	called := false

	mockService := &MockService{
		ListFunc: func(ctx context.Context) ([]*model.Issue, error) {
			called = true
			return []*model.Issue{
				{
					ID: 1,
					Title: "title test",
					Description: "desc test",
					Status: "open",
				},
			}, nil
		},
	}

	h := handler.NewHandler(mockService)
	r := chi.NewRouter()
	r.Get("/issues", h.ListIssues)

	req := httptest.NewRequest(http.MethodGet, "/issues", nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	t.Log("status:", res.Code)

	if !called {
		t.Fatal("expected ListIssues to be called")
	}

	if res.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", res.Code)
	}

	var response []*model.Issue

	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		t.Fatalf("cannot decode response: %v", err)
	}

	if len(response) != 1 || response[0].ID != 1 {
		t.Fatalf("expected valid issues list, got %v", response)
	}
}