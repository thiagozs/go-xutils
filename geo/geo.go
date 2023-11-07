package geo

import (
	"fmt"
	"strconv"
)

type Geo struct{}

func New() *Geo {
	return &Geo{}
}

func (g *Geo) Latitude(lat string) bool {
	return NewGeoValidator[string]().Latitude(lat)
}

func (g *Geo) Longitude(lon string) bool {
	return NewGeoValidator[string]().Longitude(lon)
}

func (g *Geo) Coordinates(lat, lon string) bool {
	return NewGeoValidator[string]().Coordinates(lat, lon)
}

func (g *Geo) LatitudeFloat64(lat float64) bool {
	return NewGeoValidator[float64]().Latitude(lat)
}

func (g *Geo) LongitudeFloat64(lon float64) bool {
	return NewGeoValidator[float64]().Longitude(lon)
}

func (g *Geo) CoordinatesFloat64(lat, lon float64) bool {
	return NewGeoValidator[float64]().Coordinates(lat, lon)
}

type GeoValidator[T GeoRKind] struct{}

type GeoRKind interface {
	string | float64
}

func NewGeoValidator[T GeoRKind]() *GeoValidator[T] {
	return &GeoValidator[T]{}
}

func (g *GeoValidator[T]) Latitude(lat T) bool {
	v, err := strconv.ParseFloat(fmt.Sprint(lat), 64)
	if err != nil {
		return false
	}

	return v >= -90 && v <= 90
}

func (g *GeoValidator[T]) Longitude(lon T) bool {
	v, err := strconv.ParseFloat(fmt.Sprint(lon), 64)
	if err != nil {
		return false
	}

	return v >= -180 && v <= 180
}

func (g *GeoValidator[T]) Coordinates(lat, lon T) bool {
	return g.Latitude(lat) && g.Longitude(lon)
}
