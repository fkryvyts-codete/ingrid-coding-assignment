// Package http contains logic for handling http requests and translating them to service method calls
package http

import (
	"context"
	"encoding/json"
	nethttp "net/http"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"

	"github.com/fkryvyts-codete/ingrid-coding-assignment/pkg/routes/entities"
	"github.com/fkryvyts-codete/ingrid-coding-assignment/pkg/routes/service"
)

var (
	errSrcIsMandatory    = newHTTPError("query parameter \"src\" is mandatory", 400)
	errMultipleSrcValues = newHTTPError("multiple values for query parameter \"src\" are not supported", 400)
	errDstIsMandatory    = newHTTPError("query parameter \"dst\" is mandatory", 400)
	errInternal          = newHTTPError("internal server error", 500)
)

type routesRequest struct {
	src entities.LatLng
	dst []entities.LatLng
}

type routesResponse struct {
	Source entities.LatLng   `json:"source"`
	Routes []*entities.Route `json:"routes"`
}

// RegisterHandlers registers handlers for handling incoming requests
func RegisterHandlers(mux *nethttp.ServeMux, logger log.Logger) {
	svc := service.NewService(logger)

	routesHandler := httptransport.NewServer(
		makeRoutesEndpoint(svc),
		decodeRoutesRequest,
		encodeRoutesResponse,
	)

	mux.Handle("/routes", routesHandler)
}

func makeRoutesEndpoint(svc service.Service) endpoint.Endpoint {
	return func(_ context.Context, req interface{}) (interface{}, error) {
		request, ok := req.(*routesRequest)
		if !ok {
			return nil, errInternal
		}

		routes, err := svc.ListRoutes(request.src, request.dst)
		if err != nil {
			return nil, errInternal
		}

		return &routesResponse{
			Source: request.src,
			Routes: routes,
		}, nil
	}
}

func decodeRoutesRequest(_ context.Context, r *nethttp.Request) (interface{}, error) {
	var request routesRequest

	values := r.URL.Query()

	switch src, ok := values["src"]; {
	case !ok:
		return nil, errSrcIsMandatory
	case len(src) > 1:
		return nil, errMultipleSrcValues
	default:
		request.src = entities.LatLng(src[0])
	}

	if _, ok := values["dst"]; !ok {
		return nil, errDstIsMandatory
	}

	for _, v := range values["dst"] {
		request.dst = append(request.dst, entities.LatLng(v))
	}

	return &request, nil
}

func encodeRoutesResponse(_ context.Context, w nethttp.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
