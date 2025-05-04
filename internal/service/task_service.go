package service

import (
	"context"
	"log/slog"
	"time"
	"todo_service/internal/dto"
	"todo_service/internal/model"
	"todo_service/internal/repo"
)

type service struct {
	log  *slog.Logger
	repo repo.Repository
}

func NewTaskService(repo repo.Repository, log *slog.Logger) Service {
	return &service{repo: repo, log: log}
}

type Service interface {
	CreateTask(task *dto.CreateTaskRequest) (int, error)
}

func (service *service) CreateTask(oldTask *dto.CreateTaskRequest) (int, error) {
	newTask := &model.Task{
		Title:       oldTask.Title,
		Description: oldTask.Description,
		Status:      "in_progress",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	return service.repo.CreateTask(context.Background(), newTask)
}
