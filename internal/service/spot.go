package service

import (
	"context"
	"skate-spots/internal/domain"
	"skate-spots/internal/repository"
	"time"

	"github.com/google/uuid"
)

type CreateSpotInput struct {
	Name string
	Lat  float64
	Lon  float64
}

type SpotService struct {
	repo repository.SpotRepository
}

func NewSpotService(repo repository.SpotRepository) *SpotService {
	return &SpotService{repo: repo}
}

func (s *SpotService) CreateSpot(ctx context.Context, input CreateSpotInput) (*domain.Spot, error) {
	if ctx.Err() != nil {
		return nil, ctx.Err()
	}
	newspot := &domain.Spot{
		ID:        uuid.New().String(),
		Name:      input.Name,
		Latitude:  input.Lat,
		Longitude: input.Lon,
		AddedAt:   time.Now().UTC(),
	}
	if err := newspot.Validate(); err != nil {
		return nil, err
	}
	if err := s.repo.Create(ctx, newspot); err != nil {
		return nil, err
	}
	return newspot, nil
}

func (s *SpotService) GetSpotByID(ctx context.Context, id string) (*domain.Spot, error) {
	if ctx.Err() != nil {
		return nil, ctx.Err()
	}
	if id == "" {
		return nil, domain.ErrInvalidID
	}
	spot, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return spot, nil
}

func (s *SpotService) ListSpots(ctx context.Context) ([]*domain.Spot, error) {
	if ctx.Err() != nil {
		return nil, ctx.Err()
	}
	return s.repo.List(ctx)
}
