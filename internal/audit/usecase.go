package audit

import (
	"context"
	"github.com/22Fariz22/analytics/internal/audit/entity"
	"github.com/22Fariz22/analytics/pkg/logger"
)

// UseCase interface for usecase
type UseCase interface {
	Save(ctx context.Context, l logger.Interface, data *entity.Analytics) error
}
