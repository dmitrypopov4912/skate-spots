package memory

import (
	"context"
	"skate-spots/internal/domain"
	"sync"
)

type SpotRepository struct {
	mu    sync.RWMutex
	spots map[string]*domain.Spot
}

func NewSpotRepository() *SpotRepository {
	return &SpotRepository{
		spots: make(map[string]*domain.Spot),
	}
}

func (r *SpotRepository) Create(ctx context.Context, spot *domain.Spot) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if ctx.Err() != nil {
		return ctx.Err()
	}
	r.spots[spot.ID] = spot
	return nil
}

func (r *SpotRepository) GetByID(ctx context.Context, id string) (*domain.Spot, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	if ctx.Err() != nil {
		return nil, ctx.Err()
	}
	if val, ok := r.spots[id]; ok {
		return val, nil
	}
	return nil, domain.ErrSpotNotFound
}

func (r *SpotRepository) List(ctx context.Context) ([]*domain.Spot, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	if ctx.Err() != nil {
		return nil, ctx.Err()
	}
	spotsList := make([]*domain.Spot, 0, len(r.spots))
	for _, spot := range r.spots {
		spotsList = append(spotsList, spot)
	}
	return spotsList, nil
}
