// Package entities contains entities that are necessary for business logic and not dependent on transport
package entities

import (
	"strconv"
	"strings"
)

// LatLng represents pair of latitude/longitude
type LatLng string

// Valid validates LatLng to make sure that it is a valid latitude/longitude pair
func (l LatLng) Valid() bool {
	s := strings.Split(string(l), ",")
	if len(s) != 2 {
		return false
	}

	for _, v := range s {
		if _, err := strconv.ParseFloat(v, 64); err != nil {
			return false
		}
	}

	return true
}

// Route represents a single route from source to destination
type Route struct {
	Destination LatLng
	Duration    float32
	Distance    float32
}
