package app

import (
	"fmt"
	"github.com/22Fariz22/analytics/internal/audit"
	handler "github.com/22Fariz22/analytics/internal/audit/delivery/http"
	"github.com/22Fariz22/analytics/internal/audit/storage"
	"github.com/22Fariz22/analytics/internal/audit/usecase"
	"github.com/22Fariz22/analytics/internal/audit/worker"
	"github.com/22Fariz22/analytics/internal/config"
	"github.com/22Fariz22/analytics/pkg/logger"
	"github.com/22Fariz22/analytics/pkg/postgres"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
)

type app struct {
	cfg        *config.Config
	httpServer *http.Server
	UC         audit.UseCase
}

// NewApp create
func NewApp(cfg *config.Config) *app {

	// Repository
	db, err := postgres.New(cfg.DatabaseURI, postgres.MaxPoolSize(2))
	if err != nil {
		log.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}

	repo := storage.NewPGRepository(db)

	return &app{
		cfg:        cfg,
		httpServer: nil,
		UC:         usecase.NewUseCase(repo),
	}
}

func (a *app) Run() {
	l := logger.New("debug")
	l.Info("app start")

	workers := worker.NewWorkerPool(l, a.UC)
	workers.RunWorkers(10)
	defer workers.Stop()

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	hd := handler.NewHandler(a.UC, a.cfg, workers, l)

	r.Post("/analitycs", hd.Analitycs)

	http.ListenAndServe(":8080", r)
}
