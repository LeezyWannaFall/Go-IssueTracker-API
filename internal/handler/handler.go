package handler

import (
	"Go-IssueTracker-API/internal/service"
)

type Handler struct {
	issueService service.IssueService
}

func NewHandler(issueService service.IssueService) *Handler {
	return &Handler{issueService: issueService}
}
