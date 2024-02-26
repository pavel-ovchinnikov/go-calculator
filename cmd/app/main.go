package main

import (
	"log/slog"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/pavel-ovchinnikov/go-calculator/internal/application"
	"github.com/pavel-ovchinnikov/go-calculator/internal/config"
)

func main() {
	// connect database
	dbURL := "postgres://user:password@postgresdb:5432/calculator"
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		slog.Error("connect database:", err)
		return
	}

	// config
	path, exists := os.LookupEnv("CONFIG_PATH")
	if !exists {
		slog.Error("invalid CONFIG_PATH: ")
	}
	cfg, err := config.NewConfig(path)
	if err != nil {
		slog.Error("config: ", err)
		return
	}

	app := application.NewApplication(cfg, db)
	app.Run()
}
