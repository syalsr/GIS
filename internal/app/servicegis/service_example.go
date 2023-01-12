package servicegis

import (
	"context"

	gis "github.com/syalsr/GIS/internal/usecase"
	api "github.com/syalsr/GIS/pkg/GIS-api/GIS/v1"
)

type GrpcGIS struct {
	api.UnimplementedGISServer
	gis.GIS
}

func (g *GrpcGIS) CreateStop(ctx context.Context, stop *api.RequestStop) (*api.ResponseWithOnlyError, error) {
	err := g.GIS.CreateStop()
	return &api.ResponseWithOnlyError{ Erorr: err.Error() }, nil
}

func (g *GrpcGIS) CreateBus(ctx context.Context, bus *api.RequestBus) (*api.ResponseWithOnlyError, error) {
	err := g.GIS.CreateBus()
	return &api.ResponseWithOnlyError{ Erorr: err.Error() }, nil
}

func (g *GrpcGIS) BuildRoute(ctx context.Context, route *api.RequestRoute) (*api.ResponseRoute, error) {
	req, err := g.GIS.BuildRoute()
	return &api.ResponseRoute{
		Error: err.Error(),
		TotalTime: req.TotalTime,
		Stops: req.Stops,
	 }, nil
}
