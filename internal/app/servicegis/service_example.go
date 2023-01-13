package servicegis

import (
	"context"
	"log"

	"github.com/syalsr/GIS/internal/usecase"
	gis "github.com/syalsr/GIS/internal/usecase"
	api "github.com/syalsr/GIS/pkg/GIS-api/GIS/v1"

	"google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type GrpcGIS struct {
	api.UnimplementedGISServer
	gis gis.Interface
}

func NewGrcpGIS() *GrpcGIS{
	return &GrpcGIS{
		gis: usecase.NewGIS(),
	}
}

// CreateStop - handler for create stop
func (g *GrpcGIS) CreateStop(ctx context.Context, stop *api.RequestStop) (*emptypb.Empty, error) {
	err := g.gis.CreateStop()
	if err != nil {
		return nil, err
	}
	log.Printf("%d %d %s", stop.Latitude, stop.Longitude, stop.Name)
	return nil, nil
}

// CreateBus - handler for create bus
func (g *GrpcGIS) CreateBus(ctx context.Context, bus *api.RequestBus) (*emptypb.Empty, error) {
	err := g.gis.CreateBus()
	if err != nil {
		return nil, err
	}
	status.Code(err)
	return nil, nil
}

// BuildRoute - handler for build route from one stop to another stop
func (g *GrpcGIS) BuildRoute(ctx context.Context, route *api.RequestRoute) (*api.ResponseRoute, error) {
	req, err := g.gis.BuildRoute()
	return &api.ResponseRoute{
		Error: err.Error(),
		TotalTime: req.TotalTime,
		Stops: req.Stops,
	 }, nil
}
