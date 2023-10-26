package main

import (
	"context"
	fund_management_information_system "fund-management-information-system"
	_ "fund-management-information-system/docs"
	"fund-management-information-system/pkg/handler"
	"fund-management-information-system/pkg/repository"
	"fund-management-information-system/pkg/repository/postgres"
	"fund-management-information-system/pkg/service"
	"github.com/golang-migrate/migrate/v4"
	migrate_postgres "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/gookit/slog"
	"github.com/jmoiron/sqlx"
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
func applyMigrations(db *sqlx.DB) error {
	driver, err := migrate_postgres.WithInstance(db.DB, &migrate_postgres.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://shema/",
		"postgres", driver)
	if err != nil {
		return err
	}

	if err := m.Up(); err != nil {
		if err == migrate.ErrNoChange {
			// Изменений не было
			return nil
		}
		return err
	}

	return nil
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

	if err := initConfig(); err != nil {
		slog.Fatalf("Ошибка инициализации конфига")
	}
	db, err := postgres.NewPostgresDB(postgres.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err := applyMigrations(db); err != nil {
		slog.Fatalf("Ошибка в миграции", err.Error())
	}

	if err != nil {
		slog.Fatalf("Ошибка подключения к БД")
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(fund_management_information_system.Server)

	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
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
