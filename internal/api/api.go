package api

import (
	"github.com/gofiber/fiber/v2"
	"todo_service/internal/service"
)

type Routers struct {
	service service.Service
}

func NewRouters(routers *Routers) *fiber.App {
	app := fiber.New()
	group := app.Group("/v1")

	group.Post("create_task", routers.service.CreateTask)

	return app
}
