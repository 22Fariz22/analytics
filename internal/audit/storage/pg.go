package storage

import (
	"context"
	"fmt"
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
func (p *pgRepository) Save(ctx context.Context, l logger.Interface, data *entity.Analytics) error {
	fmt.Println("repo Save.")

	_, err := p.Pool.Exec(ctx, `insert into audit (user_id, time) values($1, $2)`, data.UserID, data.UploadedAt)
	if err != nil {
		fmt.Println("err exec:", err)
	}

	return nil
}
