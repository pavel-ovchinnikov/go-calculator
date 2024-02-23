package application

import (
	"log/slog"
	"net/http"

	"github.com/pavel-ovchinnikov/go-calculator/internal/handler"
)

type Application struct{}

func (app *Application) Run() {
	sumHandler := handler.NewSumHandler()

	mux := http.NewServeMux()
	mux.HandleFunc("/sum", sumHandler.Handler)

	slog.Info("Start server")
	err := http.ListenAndServe(":3030", mux)
	if err != nil {
		slog.Error("%s", err)
	}
	slog.Info("Finish server")
}

func NewApplication() *Application {
	return &Application{}
}
