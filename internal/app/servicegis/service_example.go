package servicegis

import (
	"context"

	api "github.com/syalsr/GIS/pkg/GIS-api/GIS/v1"
)

type GIS struct {
	api.UnimplementedGISServer
}

func (g *GIS) CreateStop(context.Context, *api.RequestStop) (*api.ResponseStop, error) {
	return nil, nil
}
