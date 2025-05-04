package api

import (
	"github.com/gofiber/fiber/v2"
	"todo_service/internal/handler"
)

type Routers struct {
	TaskHandler handler.TaskHandler
}

func NewRouters(routers *Routers) *fiber.App {
	app := fiber.New()
	group := app.Group("/v1")

	group.Post("tasks", routers.TaskHandler.CreateTask)

	return app
}
