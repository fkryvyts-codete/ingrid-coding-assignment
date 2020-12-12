// Package http contains logic for handling http requests and translating them to service method calls
package http

import (
	"context"
	"encoding/json"
	nethttp "net/http"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"

	"github.com/fkryvyts-codete/ingrid-coding-assignment/pkg/routes/entities"
)

type routesRequest struct {
	src entities.LatLng
	dst []entities.LatLng
}

type routesResponse struct {
	Source entities.LatLng  `json:"source"`
	Routes []entities.Route `json:"routes"`
}

// RegisterHandlers registers handlers for handling incoming requests
func RegisterHandlers(mux *nethttp.ServeMux) {
	routesHandler := httptransport.NewServer(
		makeRoutesEndpoint(),
		decodeRoutesRequest,
		encodeRoutesResponse,
	)

	mux.Handle("/routes", routesHandler)
}

func makeRoutesEndpoint() endpoint.Endpoint {
	return func(_ context.Context, _ interface{}) (interface{}, error) {
		return routesResponse{
			Routes: []entities.Route{},
		}, nil
	}
}

func decodeRoutesRequest(_ context.Context, r *nethttp.Request) (interface{}, error) {
	var request routesRequest

	for k, values := range r.URL.Query() {
		for v := range values {
			if k == "src" {
				request.src = entities.LatLng(v)
			} else {
				request.dst = append(request.dst, entities.LatLng(v))
			}
		}
	}

	return request, nil
}

func encodeRoutesResponse(_ context.Context, w nethttp.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
