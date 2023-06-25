package worker

import (
	"context"
	"errors"
	"github.com/22Fariz22/analytics/internal/audit"
	"github.com/22Fariz22/analytics/internal/audit/entity"
	"github.com/22Fariz22/analytics/pkg/logger"
	"sync"
)

// Pool структура для воркера
type Pool struct {
	l        logger.Interface
	wg       sync.WaitGroup
	once     sync.Once
	shutDown chan struct{}
	mainCh   chan workerData
	UC       audit.UseCase
}

// NewWorkerPool создание воркера
func NewWorkerPool(l logger.Interface, UC audit.UseCase) *Pool {
	return &Pool{
		l:        l,
		wg:       sync.WaitGroup{},
		once:     sync.Once{},
		shutDown: make(chan struct{}),
		mainCh:   make(chan workerData, 10),
		UC:       UC,
	}
}

// workerData структура содержания воркера
type workerData struct {
	data *entity.Analytics
}

// AddJob запуск в handler
func (w *Pool) AddJob(ctx context.Context, l logger.Interface, dataAnalytics *entity.Analytics) error {
	select {
	case <-w.shutDown:
		return errors.New("all channels are closed")
	case w.mainCh <- workerData{
		data: dataAnalytics,
	}:
		return nil
	}
}

// RunWorkers запуск в App
func (w *Pool) RunWorkers(count int) {
	for i := 0; i < count; i++ {
		w.wg.Add(1)
		go func() {
			defer w.wg.Done()
			for {
				select {
				case <-w.shutDown:
					return
				case data, ok := <-w.mainCh:
					if !ok {
						return
					}
					err := w.UC.Save(context.Background(), w.l, data.data)
					if err != nil {
						w.l.Info("", err)
					}
				}
			}
		}()
	}
}

// Stop остановка воркера
func (w *Pool) Stop() {
	w.once.Do(func() {
		close(w.shutDown)
		close(w.mainCh)
	})
	w.wg.Wait()
}
