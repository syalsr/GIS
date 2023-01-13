package usecase

import (
	"context"
	"fmt"

	"github.com/syalsr/GIS/internal/config"
	"github.com/syalsr/GIS/internal/model"
	"github.com/syalsr/GIS/internal/repository"

)

// Interface - interface for GIS
type Interface interface {
	CreateStop(ctx context.Context, name string, roadDistance []model.RoadDistance, latitude, longitude int32)
	CreateBus(ctx context.Context, name string, stop []string, IsRoundtrip bool)
	BuildRoute(ctx context.Context, from, to string) (*model.ResponseRoute, error)
}

// GIS - type implement Interface
type GIS struct {
	R repository.Repository
}

// NewGIS - create new obj which implement Interface
func NewGIS(ctx context.Context, cfg *config.App) Interface {
	return &GIS{
		R: repository.NewClient(cfg),
	}
}

// CreateStop - create stop
func (g *GIS) CreateStop(ctx context.Context, name string, roadDistance []model.RoadDistance, longitude, latitude int32) {
	if g.R.IsStrExist(ctx, "stop", "name", name) {
		g.R.UpdateStop(ctx, model.Stop{Name: name, Longitude: longitude, Latitude: latitude})
		return
	}
	g.R.CreateStop(ctx, model.Stop{Name: name, Longitude: longitude, Latitude: latitude})
}

// CreateBus - create bus
func (g *GIS) CreateBus(ctx context.Context, name string, stop []string, isRoundtrip bool) {
	if g.R.IsStrExist(ctx, "bus", "name", name) {
		g.R.UpdateBus(ctx, model.Bus{Name: name, IsRoundtrip: isRoundtrip})
		return
	}
	g.R.CreateBus(ctx, model.Bus{Name: name, IsRoundtrip: isRoundtrip})
}

// BuildRoute - build route from one stop to another
func (g *GIS) BuildRoute(ctx context.Context, from, to string) (*model.ResponseRoute, error) {

	return nil, nil
}
