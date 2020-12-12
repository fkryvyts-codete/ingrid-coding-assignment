// Package service contains business logic for the application
package service

import (
	"github.com/go-kit/kit/log"

	"github.com/fkryvyts-codete/ingrid-coding-assignment/pkg/routes/entities"
)

// Service represents application service
type Service interface {
	ListRoutes(src entities.LatLng) ([]entities.Route, error)
}

// NewService creates new service instance
func NewService(logger log.Logger) Service {
	return &loggingMiddleware{
		logger: logger,
		svc:    &service{},
	}
}

type service struct {
}

func (s *service) ListRoutes(src entities.LatLng) ([]entities.Route, error) {
	var result []entities.Route

	result = append(result, entities.Route{
		Destination: "0,0",
		Duration:    10.5,
		Distance:    300.0,
	})

	return result, nil
}
