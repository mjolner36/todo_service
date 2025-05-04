package handler

import (
	"github.com/gofiber/fiber/v2"
	"log/slog"
	"todo_service/internal/dto"
	"todo_service/internal/service"
)

type taskHandler struct {
	taskService service.Service
	log         *slog.Logger
}

func NewTaskHandler(service service.Service, log *slog.Logger) TaskHandler {
	return &taskHandler{taskService: service, log: log}
}

type TaskHandler interface {
	CreateTask(c *fiber.Ctx) error
}

func (taskHandler *taskHandler) CreateTask(c *fiber.Ctx) error {
	task := &dto.CreateTaskRequest{}
	if err := c.BodyParser(task); err != nil {
		taskHandler.log.Error("failed to parse body", "err", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}

	createTask, err := taskHandler.taskService.CreateTask(task)
	if err != nil {
		taskHandler.log.Error("failed to create task", "err", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "failed to create task"})
	}
	return c.Status(fiber.StatusCreated).JSON(createTask)
}
