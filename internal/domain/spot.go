package domain

import (
	"errors"
	"strings"
	"time"
	"unicode/utf8"
)

type SpotFeature string

const (
	HasHighLedge SpotFeature = "high_ledge"
	HasLowLedge  SpotFeature = "low_ledge"
	HasHighRail  SpotFeature = "high_rail"
	HasLowRail   SpotFeature = "low_rail"
	HasGap       SpotFeature = "gap"
	HasStairs    SpotFeature = "stairs"
)

var (
	ErrCoordinates   = errors.New("invalid coordinates")
	ErrEmptyName     = errors.New("empty/short name")
	ErrSpotNotFound  = errors.New("spot not found")
	ErrInvalidFigure = errors.New("invalid figure")
)

type Spot struct {
	ID        string        `json:"id"`
	Name      string        `json:"name"`
	Latitude  float64       `json:"latitude"`
	Longitude float64       `json:"longitude"`
	AddedAt   time.Time     `json:"added_at"`
	Light     bool          `json:"light"`
	Figures   []SpotFeature `json:"types"`
}

func (s Spot) Validate() error {
	trimmed := strings.TrimSpace(s.Name)
	if trimmed == "" || utf8.RuneCountInString(trimmed) < 3 {
		return ErrEmptyName
	}
	if s.Latitude > 90 || s.Latitude < -90 {
		return ErrCoordinates
	}
	if s.Longitude < -180 || s.Longitude > 180 {
		return ErrCoordinates
	}
	for _, elem := range s.Figures {
		switch elem {
		case HasHighLedge, HasLowLedge, HasHighRail, HasLowRail, HasGap, HasStairs:
		default:
			return ErrInvalidFigure
		}
	}
	return nil
}
