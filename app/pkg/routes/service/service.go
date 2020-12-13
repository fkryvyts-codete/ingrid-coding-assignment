// Package service contains business logic for the application
package service

import (
	"errors"
	"sort"
	"sync"

	"github.com/go-kit/kit/log"

	"github.com/fkryvyts-codete/ingrid-coding-assignment/pkg/osrm"
	"github.com/fkryvyts-codete/ingrid-coding-assignment/pkg/routes/entities"
)

// Service represents application service
type Service interface {
	ListRoutes(src entities.LatLng, dst []entities.LatLng) ([]*entities.Route, []error)
}

// NewService creates new service instance
func NewService(logger log.Logger) Service {
	return &loggingMiddleware{
		logger: logger,
		svc: &service{
			client: osrm.NewClient(),
		},
	}
}

type service struct {
	client osrm.Client
}

type fetchRouteResult struct {
	err   error
	route *entities.Route
}

func (s *service) ListRoutes(src entities.LatLng, dst []entities.LatLng) ([]*entities.Route, []error) {
	var (
		result []*entities.Route
		errs   []error
	)

	var wg sync.WaitGroup

	c := make(chan *fetchRouteResult)

	for _, v := range dst {
		wg.Add(1)

		go s.fetchRoute(src, v, c, &wg)
	}

	go func() {
		wg.Wait()
		close(c)
	}()

	for r := range c {
		if r.err == nil {
			result = append(result, r.route)
		} else {
			errs = append(errs, r.err)
		}
	}

	sort.Slice(result, func(i, j int) bool {
		if result[i].Duration < result[j].Duration {
			return true
		}
		if result[i].Duration > result[j].Duration {
			return false
		}
		return result[i].Distance < result[j].Distance
	})

	return result, errs
}

func (s *service) fetchRoute(src, dst entities.LatLng, c chan *fetchRouteResult, wg *sync.WaitGroup) {
	defer wg.Done()

	resp, err := s.client.Driving(src, dst)

	switch {
	case err != nil:
		c <- &fetchRouteResult{err: err}
	case resp.Code != osrm.StatusOk:
		c <- &fetchRouteResult{err: errors.New("response status is not Ok")}
	case len(resp.Routes) == 0:
		c <- &fetchRouteResult{err: errors.New("empty routes list")}
	default:
		c <- &fetchRouteResult{
			route: &entities.Route{
				Destination: dst,
				Duration:    resp.Routes[0].Duration,
				Distance:    resp.Routes[0].Distance,
			},
		}
	}
}
