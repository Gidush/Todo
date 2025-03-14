package app

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// DeleteTask удаляет задачу по её ID.
// @Summary Удалить задачу
// @Description Удаляет задачу с указанным ID.
// @Tags tasks
// @Accept json
// @Produce json
// @Param id path int true "ID задачи"
// @Success 200 "Задача успешно удалена"
// @Failure 400 "Неверный формат ID задачи"
// @Failure 500 "Внутренняя ошибка сервера"
// @Router /tasks/{id} [delete]
func (i *Implementation) DeleteTask(ctx *fiber.Ctx) error {
	ID, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return err
	}

	err = i.taskService.DeleteTask(ctx.Context(), ID)
	if err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusOK)
}
