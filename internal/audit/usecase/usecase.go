package usecase

import (
	"context"

	"github.com/22Fariz22/analytics/internal/audit"
	"github.com/22Fariz22/analytics/internal/audit/entity"
	"github.com/22Fariz22/analytics/pkg/logger"
)

type useCase struct {
	repo audit.Repo
}

// NewUseCase create usecase
func NewUseCase(repo audit.Repo) *useCase {
	return &useCase{repo: repo}
}

// Save delivery data to usecase method
func (u *useCase) Save(ctx context.Context, l logger.Interface, data *entity.Analytics) error {
	return u.repo.Save(ctx, l, data)
}
