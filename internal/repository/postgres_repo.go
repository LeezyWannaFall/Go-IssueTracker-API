package repository

import (
	"context"
	"database/sql"
	"Go-IssueTracker-API/internal/model"
	"errors"
)

type PostgresIssueRepository struct {
	db *sql.DB
}

func NewPostgresIssueRepository(db *sql.DB) *PostgresIssueRepository {
	return &PostgresIssueRepository{db: db}
}

func (r *PostgresIssueRepository) CreateIssue(ctx context.Context, issue *model.Issue) (int, error) {
	var id int
	query := `
		INSERT INTO issues (title, description, status)
		VALUES ($1, $2, $3)
		RETURNING id
	`
	err := r.db.QueryRowContext(ctx, query, issue.Title, issue.Description, issue.Status).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *PostgresIssueRepository) GetIssueByID(ctx context.Context, id int) (*model.Issue, error) {
	var issue model.Issue

	query := "SELECT id, title, description, status FROM issues WHERE id = $1"
	err := r.db.QueryRowContext(ctx, query, id).Scan(&issue.ID, &issue.Title, &issue.Description, &issue.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("not found")
		}
		return nil, err
	}

	return &issue, nil
}

func (r *PostgresIssueRepository) UpdateIssue(ctx context.Context, issue *model.Issue) error {
	query := `UPDATE issues
		SET title = $1,
			description = $2,
			status = $3
		WHERE id = $4;`

	result, err := r.db.ExecContext(ctx, query, issue.Title, issue.Description, issue.Status, issue.ID)
	if err != nil {
		return errors.New("invalid request")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("not found")
	}

	return nil
}

func (r *PostgresIssueRepository) DeleteIssue(ctx context.Context, id int) error {
	query := "DELETE FROM issues WHERE id = $1"
	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("not found")
	}

	return nil
}

func (r *PostgresIssueRepository) ListIssues(ctx context.Context) ([]*model.Issue, error) {
	query := "SELECT id, title, description, status FROM issues"
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var issues []*model.Issue
	for rows.Next() {
		var issue model.Issue
		err := rows.Scan(&issue.ID, &issue.Title, &issue.Description, &issue.Status)
		if err != nil {
			return nil, err
		}
		issues = append(issues, &issue)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return issues, nil
}