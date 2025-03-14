package storage

import (
	"context"
	"todo/internal/model"
)

type TaskStorage interface {
	Create(ctx context.Context, task model.Task) (model.Task, error)
	Update(ctx context.Context, task model.Task) error
	Delete(ctx context.Context, ID int) error
	GetAll(ctx context.Context) ([]model.Task, error)
}
