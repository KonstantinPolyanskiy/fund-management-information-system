package main

import (
	"context"
	fund_management_information_system "fund-management-information-system"
	"fund-management-information-system/pkg/handler"
	"fund-management-information-system/pkg/repository"
	"fund-management-information-system/pkg/repository/postgres"
	"fund-management-information-system/pkg/service"
	"github.com/gookit/slog"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("developer")

	return viper.ReadInConfig()
}

// @title Invest fund management system
// @version 1.0

// @host localhost:8080
// @BasePath /

// @securityDefinitions.apiKey ApiKeyAuth
// @in Header
// @name Authorization
func main() {
	var wg sync.WaitGroup
	slog.Configure(func(logger *slog.SugaredLogger) {
		f := logger.Formatter.(*slog.TextFormatter)
		f.EnableColor = true
	})

	/*if err := initConfig(); err != nil {
		slog.Fatalf("Ошибка инициализации конфига")
	}*/
	db, err := postgres.NewPostgresDB(postgres.Config{
		Host:     "localhost",
		Port:     "5436",
		Username: "postgres",
		DBName:   "postgres",
		SSLMode:  "disable",
		Password: "'liderkick123'",
	})

	if err != nil {
		slog.Fatalf("Ошибка подключения к БД")
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(fund_management_information_system.Server)

	go func() {
		if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
			slog.Fatalf("ошибка в запуске сервера - %s", err.Error())
		}
	}()

	wg.Add(2)
	go services.Manager.UpdateWorkInfoProcess(context.Background())
	go services.Client.UpdateInvestmentsInfoProcess(context.Background())
	wg.Done()

	slog.Println("Сервер запущен")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	slog.Println("Сервер завершил работу")

	if err := srv.Shutdown(context.Background()); err != nil {
		slog.Fatalf("ошибка в остановке сервера - %s", err.Error())
	}
	if err := db.Close(); err != nil {
		slog.Fatalf("ошибка в остановке базы данных - %s", err.Error())
	}
}
