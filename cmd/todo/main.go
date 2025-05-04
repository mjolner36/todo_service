package main

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"todo_service/config"
	"todo_service/internal/api"
	"todo_service/internal/handler"
	"todo_service/internal/repo"
	"todo_service/internal/service"
)

func main() {

	err := godotenv.Load(".env")

	// Загружаем конфигурацию из переменных окружения
	var cfg config.AppConfig
	if err := envconfig.Process("", &cfg); err != nil {
		//slog.Fatal(errors.Wrap(err, "failed to load configuration"))
	}

	//TODO: инициализация логгера
	log := slog.Default()

	//подключение к PostgresSQL
	taskRepository, err := repo.NewRepository(context.Background(), cfg.PostgresSQL)
	if err != nil {
		//log.Fatal(errors.Wrap(err, "failed to initialize repository"))
	}
	//Создание сервиса с бизнес-логикой
	taskService := service.NewTaskService(taskRepository, log)

	//Handler
	taskHandler := handler.NewTaskHandler(taskService, log)

	//Инициализация API
	app := api.NewRouters(&api.Routers{TaskHandler: taskHandler})

	go func() {
		err := app.Listen(":" + cfg.Rest.ListenAddress)
		if err != nil {
			return
		}
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	<-signalChan
}
