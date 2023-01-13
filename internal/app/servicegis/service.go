package servicegis

import (
	"context"

	"github.com/syalsr/GIS/internal/config"
	"github.com/syalsr/GIS/internal/model"
	"github.com/syalsr/GIS/internal/usecase"
	gis "github.com/syalsr/GIS/internal/usecase"
	api "github.com/syalsr/GIS/pkg/GIS/v1"

	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type GrpcGIS struct {
	api.UnimplementedGISServer
	gis gis.Interface
}

func NewGrcpGIS(ctx context.Context, cfg *config.App) *GrpcGIS {
	return &GrpcGIS{
		gis: usecase.NewGIS(ctx, cfg),
	}
}

// CreateStop - handler for create stop
func (g *GrpcGIS) CreateStop(ctx context.Context, stop *api.RequestStop) (*emptypb.Empty, error) {
	roadDistance := make([]model.RoadDistance, 0, len(stop.RoadDistance))
	for _, item := range stop.RoadDistance {
		roadDistance = append(roadDistance, model.RoadDistance{
			Name:      item.Name,
			Curvature: item.Curvature,
			Velocity:  item.Velocity,
		})
	}
	g.gis.CreateStop(ctx, stop.Name, roadDistance, stop.Longitude, stop.Latitude)

	return nil, nil
}

// CreateBus - handler for create bus
func (g *GrpcGIS) CreateBus(ctx context.Context, bus *api.RequestBus) (*emptypb.Empty, error) {
	g.gis.CreateBus(ctx, bus.Name, bus.Stop, bus.IsRoundtrip)
	return nil, nil
}

// BuildRoute - handler for build route from one stop to another stop
func (g *GrpcGIS) BuildRoute(ctx context.Context, route *api.RequestRoute) (*api.ResponseRoute, error) {
	req, err := g.gis.BuildRoute(ctx, route.From, route.To)
	if err != nil {
		return nil, err
	}
	return &api.ResponseRoute{
		Error:     err.Error(),
		TotalTime: req.TotalTime,
		Stops:     req.Stops,
	}, nil
}
