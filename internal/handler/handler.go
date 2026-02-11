package handler

import (
	"Go-IssueTracker-API/internal/service"
	"net/http"
	"encoding/json"
	"Go-IssueTracker-API/internal/model"
	"strconv"
)

type Handler struct {
	issueService *service.IssueService
}

func NewHandler(issueService *service.IssueService) *Handler {
	return &Handler{issueService: issueService}
}

func (h* Handler) CreateIssue(w http.ResponseWriter, r *http.Request) {
	var issue model.Issue

	err := json.NewDecoder(r.Body).Decode(&issue) // парсим тело запроса в структуру Issue
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	id, err := h.issueService.CreateIssue(r.Context(), &issue) // вызываем сервис для создания новой issue
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]int{"id": id}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) GetIssueByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id") // парсим ID из URL
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid issue ID", http.StatusBadRequest)
		return
	}

	issue, err := h.issueService.GetIssueByID(r.Context(), id) // вызываем сервис для получения issue по ID
	if err != nil {
		http.Error(w, "issue not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(issue)
}

func (h *Handler) UpdateIssue(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id") // парсим ID из URL
	id, err := strconv.Atoi(idStr) 
	if err != nil {
		http.Error(w, "Invalid issue ID", http.StatusBadRequest)
		return
	}

	var issue model.Issue
	err = json.NewDecoder(r.Body).Decode(&issue) // парсим тело запроса в структуру Issue
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	issue.ID = id // устанавливаем ID из URL в структуру Issue
	err = h.issueService.UpdateIssue(r.Context(), &issue) // вызываем сервис для обновления issue
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) DeleteIssue(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id") // парсим ID из URL
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid issue ID", http.StatusBadRequest)
		return
	}

	err = h.issueService.DeleteIssue(r.Context(), id) // вызываем сервис для удаления issue по ID
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) ListIssues(w http.ResponseWriter, r *http.Request) {
	issues, err := h.issueService.ListIssues(r.Context()) // вызываем сервис для получения списка всех issues
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(issues)
}