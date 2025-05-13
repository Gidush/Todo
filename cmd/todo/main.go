package main

import (
	"context"
	"log"
	"todo/internal/app"
	"todo/internal/config"
	"todo/internal/service/task"
	"todo/internal/storage/postgres"

	_ "todo/cmd/todo/docs"
	"todo/cmd/todo/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
	"github.com/jackc/pgx/v5/pgxpool"
)

// @title Todo API
// @version 1.0
// @description API для управления задачами
// @host localhost:8080
// @BasePath /
func main() {
	conf, err := config.Load()
	if err != nil {
		log.Fatal("can`t load config: ", err)
	}
	pool, err := pgxpool.New(context.Background(), conf.Database.ConnString)
	if err != nil {
		log.Fatal("can`t connect to database: ", err)
	}
	defer pool.Close()

	strg := postgres.NewStorage(pool)
	taskService := task.NewTaskService(strg)
	service := app.New(taskService)

	app := fiber.New(fiber.Config{})

	app.Get("/swagger/*", swagger.HandlerDefault)

	app.Use(middleware.ErrorResponseHandler)
	app.Use(recover.New(recover.Config{EnableStackTrace: true}))

	app.Post("/tasks", service.CreateTask)
	app.Get("/tasks", service.GetAllTasks)
	app.Put("/tasks/:id", service.UpdateTask)
	app.Delete("/tasks/:id", service.DeleteTask)

	log.Fatal(app.Listen(":" + conf.Server.Port))
}
