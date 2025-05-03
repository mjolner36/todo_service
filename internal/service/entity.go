package service

type Task struct {
	Title       string `json:"name" validate:"required"`
	Description string `json:"description"`
}
