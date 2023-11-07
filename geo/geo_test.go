package geo

import (
	"testing"
)

func TestLatitude(t *testing.T) {
	geo := New()
	tests := []struct {
		lat      string
		expected bool
	}{
		{"45.0", true},
		{"-90.0", true},
		{"90.0", true},
		{"-91.0", false},
		{"91.0", false},
		{"abc", false},
	}

	for _, test := range tests {
		if res := geo.Latitude(test.lat); res != test.expected {
			t.Errorf("Latitude(%s) = %v; want %v", test.lat, res, test.expected)
		}
	}
}

func TestLongitude(t *testing.T) {
	geo := New()
	tests := []struct {
		lon      string
		expected bool
	}{
		{"-180.0", true},
		{"180.0", true},
		{"-181.0", false},
		{"181.0", false},
		{"abc", false},
	}

	for _, test := range tests {
		if res := geo.Longitude(test.lon); res != test.expected {
			t.Errorf("Longitude(%s) = %v; want %v", test.lon, res, test.expected)
		}
	}
}

func TestCoordinates(t *testing.T) {
	geo := New()
	tests := []struct {
		lat      string
		lon      string
		expected bool
	}{
		{"45.0", "90.0", true},
		{"-90.0", "-180.0", true},
		{"90.0", "180.0", true},
		{"-91.0", "180.0", false},
		{"90.0", "181.0", false},
		{"abc", "def", false},
	}

	for _, test := range tests {
		if res := geo.Coordinates(test.lat, test.lon); res != test.expected {
			t.Errorf("Coordinates(%s, %s) = %v; want %v", test.lat, test.lon, res, test.expected)
		}
	}
}

func TestLatitudeFloat64(t *testing.T) {
	geo := New()
	tests := []struct {
		lat      float64
		expected bool
	}{
		{45.0, true},
		{-90.0, true},
		{90.0, true},
		{-91.0, false},
		{91.0, false},
	}

	for _, test := range tests {
		if res := geo.LatitudeFloat64(test.lat); res != test.expected {
			t.Errorf("LatitudeFloat64(%f) = %v; want %v", test.lat, res, test.expected)
		}
	}
}

func TestLongitudeFloat64(t *testing.T) {
	geo := New()
	tests := []struct {
		lon      float64
		expected bool
	}{
		{-180.0, true},
		{180.0, true},
		{-181.0, false},
		{181.0, false},
	}

	for _, test := range tests {
		if res := geo.LongitudeFloat64(test.lon); res != test.expected {
			t.Errorf("LongitudeFloat64(%f) = %v; want %v", test.lon, res, test.expected)
		}
	}
}

func TestCoordinatesFloat64(t *testing.T) {
	geo := New()
	tests := []struct {
		lat      float64
		lon      float64
		expected bool
	}{
		{45.0, 90.0, true},
		{-90.0, -180.0, true},
		{90.0, 180.0, true},
		{-91.0, 180.0, false},
		{90.0, 181.0, false},
	}

	for _, test := range tests {
		if res := geo.CoordinatesFloat64(test.lat, test.lon); res != test.expected {
			t.Errorf("CoordinatesFloat64(%f, %f) = %v; want %v", test.lat, test.lon, res, test.expected)
		}
	}
}
