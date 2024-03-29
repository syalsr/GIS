package repository

import (
	"context"

	"github.com/syalsr/GIS/internal/model"
)

type Repository interface {
	IsStrExist(ctx context.Context, table, column, value string) bool
	GetIDByName(ctx context.Context, table, column, name string) (int, error)
	Stop
	Bus
}

type Stop interface {
	CreateStop(ctx context.Context, stop model.Stop)
	UpdateStop(ctx context.Context, stop model.Stop)
	CreateCurvature(ctx context.Context, stopFrom model.Stop, roadDistance model.RoadDistance)
}

type Bus interface {
	CreateBus(ctx context.Context, bus model.Bus)
	UpdateBus(ctx context.Context, bus model.Bus)
	CreateTrip(ctx context.Context, stopName []string, busName string)
}