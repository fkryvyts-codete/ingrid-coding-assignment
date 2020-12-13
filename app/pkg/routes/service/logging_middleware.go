// Package service contains business logic for the application
package service

import (
	"fmt"

	"github.com/go-kit/kit/log"

	"github.com/fkryvyts-codete/ingrid-coding-assignment/pkg/routes/entities"
)

type loggingMiddleware struct {
	logger log.Logger
	svc    Service
}

func (l *loggingMiddleware) ListRoutes(src entities.LatLng, dst []entities.LatLng) ([]*entities.Route, []error) {
	l.logger.Log("ListRoutes", "begin", "arguments", fmt.Sprintf("(%v, %v)", src, dst))
	defer l.logger.Log("ListRoutes", "end")

	r, errs := l.svc.ListRoutes(src, dst)

	for _, err := range errs {
		l.logger.Log("Error", err.Error())
	}

	return r, errs
}
