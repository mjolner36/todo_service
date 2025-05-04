package repo

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"todo_service/config"
	"todo_service/internal/model"
)

type Repository interface {
	CreateTask(ctx context.Context, task *model.Task) (int, error)
}

type repository struct {
	pool *pgxpool.Pool
	//log  *slog.Logger
}

func NewRepository(ctx context.Context, cfg config.PostgresSQL) (Repository, error) {
	// Формируем строку подключения
	connString := fmt.Sprintf(
		`user=%s password=%s host=%s port=%d dbname=%s sslmode=%s 
        pool_max_conns=%d pool_max_conn_lifetime=%s pool_max_conn_idle_time=%s`,
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Name,
		cfg.SSLMode,
		cfg.PoolMaxConns,
		cfg.PoolMaxConnLifetime.String(),
		cfg.PoolMaxConnIdleTime.String(),
	)

	// Парсим конфигурацию подключения
	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, fmt.Errorf("failed to parse PostgreSQL config: %w", err)
	}

	// Оптимизация выполнения запросов (кеширование запросов)
	config.ConnConfig.DefaultQueryExecMode = pgx.QueryExecModeCacheDescribe

	// Создаём пул соединений с базой данных
	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("failed to create PostgreSQL connection pool: %w", err)
	}

	return &repository{pool}, nil
}

func (repo *repository) CreateTask(ctx context.Context, task *model.Task) (int, error) {

	query := `INSERT INTO tasks (title, description, status, created_at)
	          VALUES ($1, $2, $3, $4) RETURNING id`

	var id int
	err := repo.pool.QueryRow(ctx, query, task.Title, task.Description, task.Status, task.CreatedAt).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
