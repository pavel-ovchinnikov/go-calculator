package application

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/pavel-ovchinnikov/go-calculator/internal/config"
	"github.com/pavel-ovchinnikov/go-calculator/internal/handler"
)

type Application struct {
	cfg *config.Config
}

func (app *Application) Run() {
	sumHandler := handler.NewSumHandler()

	mux := http.NewServeMux()
	mux.HandleFunc("/sum", sumHandler.Handler)

	slog.Info(
		fmt.Sprintf(
			"Start server, address: %v",
			app.cfg.HttpServer.Address),
	)
	err := http.ListenAndServe(app.cfg.HttpServer.Address, mux)
	if err != nil {
		slog.Error("%s", err)
	}
	slog.Info("Finish server")
}

func NewApplication(cfg *config.Config) *Application {
	return &Application{
		cfg: cfg,
	}
}
