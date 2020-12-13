// Package service contains business logic for the application
package service

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/fkryvyts-codete/ingrid-coding-assignment/pkg/osrm"
	"github.com/fkryvyts-codete/ingrid-coding-assignment/pkg/routes/entities"
)

type mockClient struct {
	data map[string]*osrm.Response
}

func (m *mockClient) Driving(src, dst entities.LatLng) (*osrm.Response, error) {
	k := fmt.Sprintf("%s;%s", src, dst)
	if v, ok := m.data[k]; ok {
		return v, nil
	}

	return &osrm.Response{Code: "InvalidQuery"}, nil
}

func TestService(t *testing.T) {
	client := &mockClient{
		data: map[string]*osrm.Response{
			"13.388860,52.517037;13.397634,52.529407": {
				Code: osrm.StatusOk,
				Routes: []*osrm.Route{{
					Duration: 251.5,
					Distance: 1884.8,
				}},
			},
			"13.388860,52.517037;13.428555,52.523219": {
				Code: osrm.StatusOk,
				Routes: []*osrm.Route{{
					Duration: 394.2,
					Distance: 3841.7,
				}},
			},
			"13.388860,52.517037;13.428555,52.523218": {
				Code: osrm.StatusOk,
				Routes: []*osrm.Route{{
					Duration: 394.2,
					Distance: 3841.6,
				}},
			},
		},
	}

	s := &service{client: client}

	// test successful method call
	routes, errs := s.ListRoutes(
		"13.388860,52.517037",
		[]entities.LatLng{
			"13.428555,52.523219",
			"13.397634,52.529407",
			"13.428555,52.523218",
		},
	)

	assert.Len(t, errs, 0, "errors slice should be empty")
	assert.Len(t, routes, 3, "there should be 3 routes")

	expectedRoutes := []*entities.Route{
		{
			Destination: "13.397634,52.529407",
			Duration:    251.5,
			Distance:    1884.8,
		},
		{
			Destination: "13.428555,52.523218",
			Duration:    394.2,
			Distance:    3841.6,
		},
		{
			Destination: "13.428555,52.523219",
			Duration:    394.2,
			Distance:    3841.7,
		},
	}

	assert.Equal(t, expectedRoutes, routes)

	// test calling method with invalid coordinates
	routes, errs = s.ListRoutes(
		"13.388860,52.517037",
		[]entities.LatLng{
			"55.555555,55.555555",
			"13.397634,52.529407",
			"13.428555,52.523218",
		},
	)

	assert.Len(t, errs, 1, "there should be one error")
	assert.Len(t, routes, 0, "there should be no routes if there is an error")
	assert.Equal(t, "response status is not Ok", errs[0].Error())
}
