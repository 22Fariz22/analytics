package http

import (
	"context"
	"encoding/json"
	"github.com/22Fariz22/analytics/internal/audit"
	"github.com/22Fariz22/analytics/internal/audit/entity"
	"github.com/22Fariz22/analytics/internal/audit/worker"
	"github.com/22Fariz22/analytics/internal/config"
	"github.com/22Fariz22/analytics/pkg/logger"
	"io"
	"log"
	"net/http"
	"time"
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

	var dataUser entity.DataUser
	//var headers *entity.HeadersData
	//var body entity.BodyData

	payload, err := io.ReadAll(r.Body)
	if err != nil {
		h.l.Error("can't read body request", err)
		http.Error(w, "", 500)
	}

	if err := json.Unmarshal(payload, &dataUser.Body); err != nil {
		h.l.Info("error unmarshall", err)
		return
	}

	headers := map[string]string{}
	//headers append to struct for repo
	for i := range r.Header {
		headers[i] = r.Header.Get(i)
		//fmt.Println(i, " : ", r.Header.Get(i))
	}
	dataUser.Headers = headers

	userID := r.Header.Get("X-Tantum-Authorization")

	analitycsData := &entity.Analytics{
		UploadedAt: time.Now(),
		UserID:     userID,
		Data:       dataUser,
	}

	h.Workers.AddJob(ctx, h.l, analitycsData) //add data from unmarshalled data

	//status 202
	w.WriteHeader(http.StatusAccepted)
}
