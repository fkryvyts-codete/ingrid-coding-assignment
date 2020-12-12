// Package entities contains entities that are necessary for business logic and not dependent on transport
package entities

// LatLng represents pair of latitude/longitude
type LatLng string

// Route represents a single route from source to destination
type Route struct {
	Destination LatLng
	Duration    float32
	Distance    float32
}
