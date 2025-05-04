package dto

type CreateTaskRequest struct {
	Title       string `json:"name" validate:"required"`
	Description string `json:"description"`
}
