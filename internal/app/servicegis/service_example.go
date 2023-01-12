package servicegis

import (
	"context"

	api "github.com/syalsr/GIS/pkg/GIS-api/GIS/v1"
	"github.com/rs/zerolog/log"
)

type GIS struct {
	api.UnimplementedGISServer
}

func (g *GIS) CreateStop(ctx context.Context, stop *api.RequestStop) (*api.ResponseWithOnlyError, error) {
	
	return nil, nil
}
func (g *GIS) CreateBus(ctx context.Context, bus *api.RequestBus) (*api.ResponseWithOnlyError, error) {
	return nil, nil
}
func (g *GIS) BuildRoute(ctx context.Context, route *api.RequestRoute) (*api.ResponseRoute, error) {
	return nil, nil
}
