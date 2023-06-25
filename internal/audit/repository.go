package audit

import (
	"context"
	"github.com/22Fariz22/analytics/internal/audit/entity"
	"github.com/22Fariz22/analytics/pkg/logger"
)

// Repo interface for storages
type Repo interface {
	Save(ctx context.Context, l logger.Interface, data *entity.Analytics) error
}
