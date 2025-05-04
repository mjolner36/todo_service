package dto

type CreateTaskRequest struct {
	Title       string `json:"name" validate:"required"`
	Description string `json:"description"`
}

type Response struct {
	Status string `json:"status"`
	Data   any    `json:"data,omitempty"`
}
