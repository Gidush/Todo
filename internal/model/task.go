package model

import (
	"time"
	"todo/pkg/utils/validation"
	"unicode/utf8"
)

type Task struct {
	ID          int
	Title       string
	Description string
	Status      TaskStatus
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (t Task) Validate() error {
	var builder validation.ResultBuilder

	builder.Add(utf8.RuneCountInString(t.Title) < 1000,
		"Title %q can not be more than 1000 symbols", t.Title)

	builder.Add(len(t.Title) != 0,
		"Title can not be empty")

	builder.Add(t.Status.IsValid(),
		"Status %q should be one of 'new', 'done', 'in_progress'.", t.Status)

	return builder.Validate()
}

type TaskStatus string

const (
	TaskStatusNew        TaskStatus = "new"
	TaskStatusDone       TaskStatus = "done"
	TaskStatusInProgress TaskStatus = "in_progress"
)

func (status TaskStatus) String() string {
	return string(status)
}

func (s TaskStatus) IsValid() bool {
	return s == TaskStatusNew || s == TaskStatusDone || s == TaskStatusInProgress
}
