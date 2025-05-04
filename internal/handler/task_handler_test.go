package handler

import (
	"bytes"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"log/slog"
	"net/http"
	"testing"
	"todo_service/internal/dto"
	"todo_service/internal/service/mocks"
)

func TestCreateTask(t *testing.T) {

	logger := slog.Default()

	mockService := new(mocks.Service)
	handler := NewTaskHandler(mockService, logger)

	app := fiber.New()
	app.Post("/tasks", handler.CreateTask)

	t.Run("valid request", func(t *testing.T) {
		taskReq := &dto.CreateTaskRequest{
			Title:       "Test Title",
			Description: "Test Description",
		}

		body, _ := json.Marshal(taskReq)

		mockService.On("CreateTask", taskReq).
			Return(1, nil).Once()

		req, err := http.NewRequest("POST", "/tasks", bytes.NewReader(body))
		assert.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")

		// Выполняем запрос
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusCreated, resp.StatusCode)

		// Проверяем ответ
		//var response dto.Response
		//json.NewDecoder(resp.Body).Decode(&response)
		//assert.Equal(t, "success", response.Status)

		// Проверяем вызов мок-методов
		mockService.AssertExpectations(t)
	})
}
