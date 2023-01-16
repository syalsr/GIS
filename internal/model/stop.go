package model

type Stop struct {
	Name         string
	Longitude    int32
	Latitude     int32
	RoadDistance []RoadDistance
}

type RoadDistance struct {
	Name      string
	Curvature float64
	Velocity  float64
}
