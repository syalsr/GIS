package repository

import (
	"context"

	"github.com/syalsr/GIS/internal/model"
)

type Repository interface {
	IsStrExist(ctx context.Context, table, column, value string) bool
	CreateStop(ctx context.Context, stop model.Stop)
	UpdateStop(ctx context.Context, stop model.Stop)	
	CreateBus(ctx context.Context, bus model.Bus)
	UpdateBus(ctx context.Context, bus model.Bus)	
}
