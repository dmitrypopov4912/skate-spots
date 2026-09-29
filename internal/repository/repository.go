package repository

import (
	"context"
	"skate-spots/internal/domain"
)

type SpotRepository interface {
	Create(ctx context.Context, spot *domain.Spot) error
	GetByID(ctx context.Context, id string) (*domain.Spot, error)
	List(ctx context.Context) ([]*domain.Spot, error)
}
