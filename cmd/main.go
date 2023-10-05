package main

import (
	fund_management_information_system "fund-management-information-system"
	"fund-management-information-system/pkg/handler"
	"fund-management-information-system/pkg/repository"
	"fund-management-information-system/pkg/repository/postgres"
	"fund-management-information-system/pkg/service"
	"github.com/gookit/slog"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("developer")

	return viper.ReadInConfig()
}
func main() {
	slog.Configure(func(logger *slog.SugaredLogger) {
		f := logger.Formatter.(*slog.TextFormatter)
		f.EnableColor = true
	})

	if err := initConfig(); err != nil {
		slog.Fatalf("Ошибка инициализации конфига")
	}
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
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		slog.Fatalf("Ошибка запуска сервера: %s", err.Error())
	}
	//slog.Println("Сервер запущен")
}
