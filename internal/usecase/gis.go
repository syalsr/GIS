package usecase

import "github.com/syalsr/GIS/internal/model"

type GIS interface {
	CreateStop() error
	CreateBus() error
	BuildRoute() (model.RequestRoute, error)
}
