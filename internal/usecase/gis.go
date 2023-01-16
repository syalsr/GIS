package usecase

import (
	"context"

	"github.com/syalsr/GIS/internal/config"
	"github.com/syalsr/GIS/internal/model"
	"github.com/syalsr/GIS/internal/repository"
	"github.com/syalsr/GIS/internal/usecase/router"
)

// Interface - interface for GIS
type Interface interface {
	CreateStop(ctx context.Context, name string, roadDistance []model.RoadDistance, latitude, longitude int32)
	CreateBus(ctx context.Context, name string, stop []string, IsRoundtrip bool)
	BuildRoute(ctx context.Context, from, to string) (*model.ResponseRoute, error)
}

// GIS - type implement Interface
type GIS struct {
	Repo   repository.Repository
	Router router.Interface
}

// NewGIS - create new obj which implement Interface
func NewGIS(ctx context.Context, cfg *config.App) Interface {
	return &GIS{
		Repo:   repository.NewClient(cfg),
		Router: router.NewRouter(),
	}
}

// CreateStop - create stop
func (g *GIS) CreateStop(ctx context.Context, name string, roadDistance []model.RoadDistance, longitude, latitude int32) {
	edges := []router.Edge{}
	for _, item := range roadDistance {
		edges = append(edges, router.Edge{To: router.Vertex{Name: item.Name}, Weight: item.Curvature})
	}

	g.Router.FillData(&router.Vertex{
		Name:  name,
		Edges: edges,
	}, len(roadDistance) == 0)

	if g.Repo.IsStrExist(ctx, "stop", "name", name) {
		g.Repo.UpdateStop(ctx, model.Stop{Name: name, Longitude: longitude, Latitude: latitude})
		return
	}
	g.Repo.CreateStop(ctx, model.Stop{Name: name, Longitude: longitude, Latitude: latitude})
	for _, item := range roadDistance {
		if !g.Repo.IsStrExist(ctx, "stop", "name", item.Name) {
			g.Repo.CreateStop(ctx, model.Stop{Name: item.Name})
		}
		g.Repo.CreateCurvature(ctx, model.Stop{Name: name}, item)
	}

}

// CreateBus - create bus
// Assume that stops already exist
func (g *GIS) CreateBus(ctx context.Context, name string, stop []string, isRoundtrip bool) {
	if g.Repo.IsStrExist(ctx, "bus", "name", name) {
		g.Repo.UpdateBus(ctx, model.Bus{Name: name, IsRoundtrip: isRoundtrip, Stop: stop})
		return
	}
	g.Repo.CreateBus(ctx, model.Bus{Name: name, IsRoundtrip: isRoundtrip, Stop: stop})
	g.Repo.CreateTrip(ctx, stop, name)
}

// BuildRoute - build route from one stop to another
func (g *GIS) BuildRoute(ctx context.Context, from, to string) (*model.ResponseRoute, error) {
	//Check in cache later
	finishVertex := g.Router.BuildRouter(from, to)
	response := &model.ResponseRoute{}
	for finishVertex.From != nil {
		response.Stops = append(response.Stops, finishVertex.Name)
		response.TotalTime += finishVertex.Weight
	}

	return response, nil
}
