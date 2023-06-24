package http

import (
	"context"
	"github.com/22Fariz22/analytics/internal/audit"
	"github.com/22Fariz22/analytics/internal/audit/worker"
	"github.com/22Fariz22/analytics/internal/config"
	"github.com/22Fariz22/analytics/pkg/logger"
	"net/http"
)

// Handler структура хэндлер
type Handler struct {
	UC      audit.UseCase
	Cfg     config.Config
	Workers *worker.Pool
	l       logger.Interface
}

// NewHandler создает хэндлер
func NewHandler(repo audit.UseCase, cfg *config.Config, workers *worker.Pool, l logger.Interface) *Handler {
	return &Handler{
		UC:      repo,
		Cfg:     *cfg,
		Workers: workers,
		l:       l,
	}
}

func (h *Handler) GetAnalytics(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	///unmarshall json

	h.Workers.AddJob(ctx, h.l) //add data from unmarshalled data

	//status 202
	w.WriteHeader(http.StatusAccepted)
}
