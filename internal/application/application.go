package application

import (
	"fmt"
	"log/slog"
	"net/http"

	"gorm.io/gorm"

	"github.com/pavel-ovchinnikov/go-calculator/internal/adapter/sum_repository"
	"github.com/pavel-ovchinnikov/go-calculator/internal/config"
	"github.com/pavel-ovchinnikov/go-calculator/internal/handler"
	service "github.com/pavel-ovchinnikov/go-calculator/internal/services"
)

type Application struct {
	cfg *config.Config
	db  *gorm.DB
}

func (app *Application) Run() {
	// Repository
	rep := sum_repository.NewSumRepository(app.db)

	// Service
	sumService := service.NewSumService(rep)

	// Handler
	sumHandler := handler.NewSumHandler(sumService)

	// Router
	mux := http.NewServeMux()
	mux.HandleFunc("/sum", sumHandler.Handler)

	slog.Info(fmt.Sprintf("Start server, address: %v", app.cfg.HttpServer.Address))
	if err := http.ListenAndServe(app.cfg.HttpServer.Address, mux); err != nil {
		slog.Error("%s", err)
	}
	slog.Info("Finish server")
}

func NewApplication(cfg *config.Config, db *gorm.DB) *Application {
	return &Application{
		cfg: cfg,
		db:  db,
	}
}
