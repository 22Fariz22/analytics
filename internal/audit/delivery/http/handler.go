package http

import (
	"context"
	"fmt"
	"github.com/22Fariz22/analytics/internal/audit"
	"github.com/22Fariz22/analytics/internal/audit/worker"
	"github.com/22Fariz22/analytics/internal/config"
	"github.com/22Fariz22/analytics/pkg/logger"
	"io"
	"log"
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

func (h *Handler) Analitycs(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	log.Println("handler GetAnalytics.")

	//var dataAnalytics *entity.Analytics

	payload, err := io.ReadAll(r.Body)
	if err != nil {
		h.l.Error("can't read body request", err)
		http.Error(w, "", 500)
	}

	fmt.Println("payload", string(payload))

	//if err := json.Unmarshal(payload, &dataAnalytics); err != nil {
	//	h.l.Info("error unmarshall", err)
	//	return
	//}

	h.Workers.AddJob(ctx, h.l) //add data from unmarshalled data

	//status 202
	w.WriteHeader(http.StatusAccepted)
}
