package task

import (
	"context"
	"time"
	"todo/internal/model"
	"todo/internal/serrors"
	"todo/internal/storage"
)

type TaskService struct {
	taskStorage storage.TaskStorage
}

func NewTaskService(taskStorage storage.TaskStorage) *TaskService {
	return &TaskService{
		taskStorage: taskStorage,
	}
}

func (s *TaskService) CreateTask(ctx context.Context, task model.Task) (model.Task, error) {
	if err := task.Validate(); err != nil {
		return model.Task{}, serrors.ErrInvalidArgument.WithDetails("task is invalid").WithWrapped(err)
	}

	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()
	newTask, err := s.taskStorage.Create(ctx, task)
	if err != nil {
		return model.Task{}, serrors.ErrInternalServer.WithWrapped(err)
	}

	return newTask, nil
}

func (s *TaskService) UpdateTask(ctx context.Context, task model.Task) error {
	if err := task.Validate(); err != nil {
		return serrors.ErrInvalidArgument.WithDetails("task is invalid").WithWrapped(err)
	}

	task.UpdatedAt = time.Now()
	err := s.taskStorage.Update(ctx, task)
	if err != nil {
		return serrors.ErrInternalServer.WithWrapped(err)
	}

	return nil
}

func (s *TaskService) DeleteTask(ctx context.Context, ID int) error {
	if ID < 0 {
		return serrors.ErrInvalidArgument.WithDetails("ID is below zero")
	}

	if err := s.taskStorage.Delete(ctx, ID); err != nil {
		return serrors.ErrInternalServer.WithWrapped(err)
	}

	return nil
}

func (s *TaskService) GetAll(ctx context.Context) ([]model.Task, error) {
	tasks, err := s.taskStorage.GetAll(ctx)
	if err != nil {
		return nil, serrors.ErrInternalServer.WithWrapped(err)
	}

	return tasks, nil
}
