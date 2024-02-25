package main

import (
	"log/slog"
	"os"

	"github.com/pavel-ovchinnikov/go-calculator/internal/application"
	"github.com/pavel-ovchinnikov/go-calculator/internal/config"
)

func main() {
	path, exists := os.LookupEnv("CONFIG_PATH")
	if !exists {
		slog.Error("invalid CONFIG_PATH: ")
	}

	cfg, err := config.NewConfig(path)
	if err != nil {
		slog.Error("config: ", err)
		return
	}

	app := application.NewApplication(cfg)
	app.Run()
}
