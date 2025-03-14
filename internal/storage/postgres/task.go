package postgres

import (
	"context"
	"fmt"
	"time"
	"todo/internal/model"

	"github.com/georgysavva/scany/v2/pgxscan"
)

const (
	QueryReturning = " RETURNING id, title, description, status, created_at, updated_at"
	QueryInsert    = `
	INSERT INTO tasks (
		title,
		description,
		status,
		created_at,
		updated_at
	) VALUES ($1, $2, $3, $4, $5)
	`
	QueryUpdate = `
	UPDATE tasks
	SET title = $2,
    	description = $3,
    	status = $4,
    	updated_at = $5
	WHERE ID = $1
	`
	QuerySelect = "SELECT * FROM tasks"
	QueryDelete = `
	DELETE FROM tasks
	WHERE ID = $1
	`
)

func (s *Storage) Create(ctx context.Context, task model.Task) (model.Task, error) {
	query := QueryInsert + QueryReturning

	queryCtx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	taskDB := taskDBFromModel(task)
	var itemDB TaskDB
	if err := pgxscan.Get(queryCtx, s.db, &itemDB, query,
		taskDB.Title,
		taskDB.Description,
		taskDB.Status,
		taskDB.CreatedAt,
		taskDB.UpdatedAt,
	); err != nil {
		return model.Task{}, fmt.Errorf("create task failed: %w", err)
	}
	return itemDB.toModel(), nil
}

func (s *Storage) Update(ctx context.Context, task model.Task) error {
	query := QueryUpdate
	queryCtx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	taskDB := taskDBFromModel(task)
	if _, err := s.db.Exec(queryCtx, query,
		taskDB.ID,
		taskDB.Title,
		taskDB.Description,
		taskDB.Status,
		taskDB.UpdatedAt,
	); err != nil {
		return fmt.Errorf("update task failed: %w", err)
	}
	return nil
}

func (s *Storage) GetAll(ctx context.Context) ([]model.Task, error) {
	query := QuerySelect
	queryCtx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	var itemDBs []TaskDB
	if err := pgxscan.Select(queryCtx, s.db, &itemDBs, query); err != nil {
		return nil, fmt.Errorf("get tasks failed: %w", err)
	}

	tasks := make([]model.Task, 0, len(itemDBs))
	for _, itemDB := range itemDBs {
		tasks = append(tasks, itemDB.toModel())
	}
	return tasks, nil
}

func (s *Storage) Delete(ctx context.Context, ID int) error {
	query := QueryDelete
	queryCtx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	if _, err := s.db.Exec(queryCtx, query, ID); err != nil {
		return err
	}
	return nil
}
