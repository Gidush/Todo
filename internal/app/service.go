package app

import (
	"todo/internal/service/task"
)

type Implementation struct {
	taskService *task.TaskService
}

func New(taskService *task.TaskService) *Implementation {
	return &Implementation{
		taskService: taskService,
	}
}
