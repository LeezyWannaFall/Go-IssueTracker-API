package main

import (
	"Go-IssueTracker-API/internal/handler"
	"Go-IssueTracker-API/internal/repository"
	"Go-IssueTracker-API/internal/service"

	"database/sql"
	"log"
	"net/http"
	"fmt"

	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"
	"Go-IssueTracker-API/internal/config"
)

func main() {
	// TODO: init config:
	cfg, err := config.LoadConfig("config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	// init database: postgres
	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.Storage.Postgres.User,
		cfg.Storage.Postgres.Password,
		cfg.Storage.Postgres.Host,
		cfg.Storage.Postgres.Port, 
		cfg.Storage.Postgres.Database,
	)


	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("Cannot connect to database:", err)
	}

	// create repository, service and handler
	repo := repository.NewPostgresIssueRepository(db)
	svc := service.NewIssueService(repo)
	h := handler.NewHandler(svc)
	
	// init router: chi
	r := chi.NewRouter()
	r.Post("/issues", h.CreateIssue)
	r.Get("/issues/{id}", h.GetIssueByID)
	r.Put("/issues/{id}", h.UpdateIssue)
	r.Delete("/issues/{id}", h.DeleteIssue)
	r.Get("/issues", h.ListIssues)

	// run server
	addr := fmt.Sprintf(":%d", cfg.Server.Port)
	
	log.Printf("Server running on %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, r))
}