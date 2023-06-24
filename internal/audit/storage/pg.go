package storage

import (
	"context"
	"github.com/22Fariz22/analytics/internal/audit/entity"
	"github.com/22Fariz22/analytics/pkg/logger"
	"github.com/22Fariz22/analytics/pkg/postgres"
)

type pgRepository struct {
	*postgres.Postgres
}

// NewPGRepository create postgres storage
func NewPGRepository(db *postgres.Postgres) *pgRepository {
	return &pgRepository{db}
}

// Save url to db
func (p *pgRepository) Save(ctx context.Context, l logger.Interface, data []*entity.Data) error {

	return nil
}
