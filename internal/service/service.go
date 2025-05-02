package service

import (
	"github.com/gofiber/fiber/v2"
	"log/slog"
)

type Service interface {
	CreateTask(c *fiber.Ctx) error
}

type service struct {
	log *slog.Logger
}

func NewService(log *slog.Logger) Service {
	return &service{log: log}
}

func (service *service) CreateTask(c *fiber.Ctx) error {
	return nil
}
