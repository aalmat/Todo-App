package main

import (
	"github.com/aalmat/todo"
	"github.com/aalmat/todo/pkg/handler"
	"github.com/aalmat/todo/pkg/repository"
	"github.com/aalmat/todo/pkg/service"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper"
	"os"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	//runMigrationUp()

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing config: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error initializing config: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		viper.GetString("db.host"),
		viper.GetString("db.port"),
		viper.GetString("db.username"),
		os.Getenv("DB_PASSWORD"),
		viper.GetString("db.dbname"),
		viper.GetString("db.sslmode"),
	})

	if err != nil {
		logrus.Fatalf("Database connection err: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	server := new(todo.Server)
	if err := server.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("server error: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()

}

func runMigrationUp() {
	m, err := migrate.New(
		"file://schema",
		"postgres://postgres:postgres@localhost:5432/todo-db?sslmode=disable")
	if err != nil {
		logrus.Fatal(err)
	}
	if err := m.Up(); err != nil {
		logrus.Fatal(err)
	}

}

func runMigrationDown() {
	m, err := migrate.New(
		"file://schema",
		"postgres://postgres:postgres@localhost:5432/todo-db?sslmode=disable")
	if err != nil {
		logrus.Fatal(err)
	}
	if err := m.Down(); err != nil {
		logrus.Fatal(err)
	}

}
