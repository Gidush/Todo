package mapper

import (
	"todo/internal/app/api_model"
	"todo/internal/model"
)

func ResponseFromModel(t model.Task) api_model.TaskResponse {
	return api_model.TaskResponse{
		ID:          t.ID,
		Title:       t.Title,
		Description: t.Description,
		Status:      t.Status.String(),
		CreatedAt:   t.CreatedAt,
		UpdatedAt:   t.UpdatedAt,
	}
}
