package app

import (
	"encoding/json"
	"strconv"
	"todo/internal/model"

	"github.com/gofiber/fiber/v2"
)

// UpdateTask обновляет задачу по её ID.
// @Summary Обновить задачу
// @Description Обновляет задачу с указанным ID.
// @Tags tasks
// @Accept json
// @Produce json
// @Param id path int true "ID задачи"
// @Param request body UpdateRequest true "Данные для обновления задачи"
// @Success 200 "Задача успешно обновлена"
// @Failure 400 string "Неверный формат запроса"
// @Failure 500 "Внутренняя ошибка сервера"
// @Router /tasks/{id} [put]
func (i *Implementation) UpdateTask(ctx *fiber.Ctx) error {
	ID, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return err
	}
	rawData := ctx.Body()

	var request UpdateRequest
	err = json.Unmarshal(rawData, &request)
	if err != nil {
		return err
	}

	task := model.Task{
		ID:          ID,
		Title:       request.Title,
		Description: request.Description,
		Status:      model.TaskStatus(request.Status),
	}
	err = i.taskService.UpdateTask(ctx.Context(), task)
	if err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusOK)
}

type UpdateRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}
