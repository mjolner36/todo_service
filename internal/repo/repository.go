package repo

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
	CreateTask(ctx context.Context, task Task) (int, error)
}

type repository struct {
	pool *pgxpool.Pool
}

func NewRepository() (Repository, error) {
	return nil, nil
}

func (repo *repository) CreateTask(ctx context.Context, task Task) (int, error) {
	return 0, nil
}
