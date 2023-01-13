package usecase

import "github.com/syalsr/GIS/internal/model"

type Interface interface {
	CreateStop() error
	CreateBus() error
	BuildRoute() (*model.RequestRoute, error)
}

type GIS struct {

}

func NewGIS() Interface {
	return &GIS{}
}

func (g *GIS) CreateStop() error{
	return nil
}
func (g *GIS) CreateBus() error{
	return nil
}
func (g *GIS) BuildRoute() (*model.RequestRoute, error){
	return nil, nil
}