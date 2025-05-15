package service

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"log/slog"
	"testing"
	"todo_service/internal/dto"
	"todo_service/internal/model"
	"todo_service/internal/repo/mocks"
)

func TestCreateTask(t *testing.T) {
	logger := slog.Default()
	mockRepo := new(mocks.Repository)

	taskService := NewTaskService(mockRepo, logger)

	t.Run("success", func(t *testing.T) {
		req := &dto.CreateTaskRequest{
			Title:       "Test Title",
			Description: "Test Description",
		}

		// Проверяем, что репозиторий вызывается с task, у которого нужные поля
		mockRepo.On("CreateTask", mock.Anything, mock.MatchedBy(func(task *model.Task) bool {
			return task.Title == req.Title &&
				task.Description == req.Description &&
				task.Status == "in_progress"
		})).Return(1, nil).Once()

		id, err := taskService.CreateTask(req)
		assert.NoError(t, err)
		assert.Equal(t, 1, id)

		mockRepo.AssertExpectations(t)
	})

	t.Run("repository error", func(t *testing.T) {
		req := &dto.CreateTaskRequest{
			Title:       "Fail Title",
			Description: "Fail Description",
		}

		mockRepo.On("CreateTask", mock.Anything, mock.Anything).
			Return(0, errors.New("db error")).Once()

		id, err := taskService.CreateTask(req)
		assert.Error(t, err)
		assert.Equal(t, 0, id)

		mockRepo.AssertExpectations(t)
	})
}
