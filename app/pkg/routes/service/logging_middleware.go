// Package service contains business logic for the application
package service

import (
	"github.com/go-kit/kit/log"

	"github.com/fkryvyts-codete/ingrid-coding-assignment/pkg/routes/entities"
)

type loggingMiddleware struct {
	logger log.Logger
	svc    Service
}

func (l *loggingMiddleware) ListRoutes(src entities.LatLng) ([]entities.Route, error) {
	l.logger.Log("ListRoutes", "begin", "src", src)
	defer l.logger.Log("ListRoutes", "end")

	r, err := l.svc.ListRoutes(src)
	if err != nil {
		l.logger.Log("Error occurred", err.Error())
	}

	return r, err
}
