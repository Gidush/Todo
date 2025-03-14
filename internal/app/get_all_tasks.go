package app

import (
	"todo/internal/app/api_model"
	"todo/internal/app/mapper"

	"github.com/gofiber/fiber/v2"
)

// GetAllTasks возвращает список всех задач.
// @Summary Получить все задачи
// @Description Возвращает список всех задач.
// @Tags tasks
// @Accept json
// @Produce json
// @Success 200 {object} GetAllTasksResponse "Список задач"
// @Failure 500 "Внутренняя ошибка сервера"
// @Router /tasks [get]
func (i *Implementation) GetAllTasks(ctx *fiber.Ctx) error {
	tasks, err := i.taskService.GetAll(ctx.Context())
	if err != nil {
		return err
	}

	tasksResponce := make([]api_model.TaskResponse, 0, len(tasks))

	for _, t := range tasks {
		tasksResponce = append(tasksResponce, mapper.ResponseFromModel(t))
	}

	return ctx.JSON(GetAllTasksResponse{Tasks: tasksResponce})

}

type GetAllTasksResponse struct {
	Tasks []api_model.TaskResponse `json:"tasks"`
}
