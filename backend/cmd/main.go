package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	app "github.com/Ribas160/flash-cards"
	"github.com/Ribas160/flash-cards/pkg/handler"
	"github.com/Ribas160/flash-cards/pkg/repository"
	"github.com/Ribas160/flash-cards/pkg/service"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing config: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	db, err := repository.NewMysqlDB(repository.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		DBName:   os.Getenv("DB_NAME"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewServices(repos)
	handler := handler.NewHandler(services)

	srv := new(app.Server)
	go func() {
		if err := srv.Run(viper.GetString("port"), handler.InitRoutes()); err != nil {
			logrus.Fatalf("error occurred while running http server: %s", err.Error())
		}
	}()

	logrus.Println("Flash cards app stated")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Println("Flash cards app shutting down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("Error occurred on server shutting down", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("Error occurred on db connection close: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
