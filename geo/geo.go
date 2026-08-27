package geo

import (
	"fmt"
	"strconv"
)

type GeoRKind interface {
	string | float64
}

// IsLatitude reports whether value is a finite latitude from -90 to 90.
func IsLatitude[T GeoRKind](lat T) bool {
	v, err := strconv.ParseFloat(fmt.Sprint(lat), 64)
	if err != nil {
		return false
	}

	return v >= -90 && v <= 90
}

// IsLongitude reports whether value is a finite longitude from -180 to 180.
func IsLongitude[T GeoRKind](lon T) bool {
	v, err := strconv.ParseFloat(fmt.Sprint(lon), 64)
	if err != nil {
		return false
	}

	return v >= -180 && v <= 180
}

// AreCoordinates validates a latitude/longitude pair.
func AreCoordinates[T GeoRKind](lat, lon T) bool {
	return IsLatitude(lat) && IsLongitude(lon)
}
