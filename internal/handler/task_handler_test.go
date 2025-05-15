package handler

import (
	"bytes"
	"encoding/json"
	"errors"
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

		var response dto.Response
		json.NewDecoder(resp.Body).Decode(&response)
		assert.Equal(t, "success", response.Status)

		// Проверяем вызов мок-методов
		mockService.AssertExpectations(t)
	})
	t.Run("invalid json request", func(t *testing.T) {
		invalidBody := []byte(`{invalid json}`)

		req, err := http.NewRequest("POST", "/tasks", bytes.NewReader(invalidBody))
		assert.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)

		var response map[string]string
		json.NewDecoder(resp.Body).Decode(&response)
		assert.Equal(t, "invalid request", response["error"])
	})
	t.Run("service returns error", func(t *testing.T) {
		taskReq := &dto.CreateTaskRequest{
			Title:       "Test Title",
			Description: "Test Description",
		}
		body, _ := json.Marshal(taskReq)

		mockService.On("CreateTask", taskReq).
			Return(0, errors.New("db error")).Once()

		req, err := http.NewRequest("POST", "/tasks", bytes.NewReader(body))
		assert.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)

		var response map[string]string
		json.NewDecoder(resp.Body).Decode(&response)
		assert.Equal(t, "failed to create task", response["error"])
	})
}
