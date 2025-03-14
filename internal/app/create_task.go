package app

import (
	"encoding/json"
	"todo/internal/app/api_model"
	"todo/internal/app/mapper"
	"todo/internal/model"

	"github.com/gofiber/fiber/v2"
)

// CreateTask создает новую задачу.
// @Summary Создать новую задачу
// @Description Создает новую задачу с указанным заголовком и описанием.
// @Tags tasks
// @Accept json
// @Produce json
// @Param request body CreateRequest true "Данные для создания задачи"
// @Success 200 {object} CreateResponse "Задача успешно создана"
// @Failure 400 string "Неверный формат запроса"
// @Failure 500 "Внутренняя ошибка сервера"
// @Router /tasks [post]
func (i *Implementation) CreateTask(ctx *fiber.Ctx) error {
	rawData := ctx.Body()

	var request CreateRequest
	err := json.Unmarshal(rawData, &request)
	if err != nil {
		return err
	}

	task := model.Task{
		Title:       request.Title,
		Description: request.Description,
		Status:      model.TaskStatusNew,
	}
	newTask, err := i.taskService.CreateTask(ctx.Context(), task)
	if err != nil {
		return err
	}

	taskResponse := mapper.ResponseFromModel(newTask)
	return ctx.JSON(CreateResponse{Task: taskResponse})
}

type CreateRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type CreateResponse struct {
	Task api_model.TaskResponse `json:"task"`
}
