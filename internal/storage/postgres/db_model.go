package postgres

import (
	"time"
	"todo/internal/model"
)

type TaskDB struct {
	ID          int       `db:"id"`
	Title       string    `db:"title"`
	Description *string   `db:"description"`
	Status      string    `db:"status"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

func (task TaskDB) toModel() model.Task {
	var description string
	if task.Description != nil {
		description = *task.Description
	}

	return model.Task{
		ID:          task.ID,
		Title:       task.Title,
		Description: description,
		Status:      model.TaskStatus(task.Status),
		CreatedAt:   task.CreatedAt,
		UpdatedAt:   task.UpdatedAt,
	}
}

func taskDBFromModel(t model.Task) TaskDB {
	return TaskDB{
		ID:          t.ID,
		Title:       t.Title,
		Description: &t.Description,
		Status:      t.Status.String(),
		CreatedAt:   t.CreatedAt,
		UpdatedAt:   t.UpdatedAt,
	}
}
