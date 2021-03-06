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
	errSrcIsNotValid     = newHTTPError("\"src\" is not a valid pair of latitude and longitude", 400)
	errDstIsMandatory    = newHTTPError("query parameter \"dst\" is mandatory", 400)
	errDstIsNotValid     = newHTTPError("\"dst\" is not a valid pair of latitude and longitude", 400)
	errInternal          = newHTTPError("internal server error", 500)
)

type routesRequest struct {
	src entities.LatLng
	dst []entities.LatLng
}

type routesResponse struct {
	Source entities.LatLng        `json:"source"`
	Routes []*routesResponseRoute `json:"routes"`
}

type routesResponseRoute struct {
	Destination entities.LatLng `json:"destination"`
	Duration    float32         `json:"duration"`
	Distance    float32         `json:"distance"`
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

		rr, errs := svc.ListRoutes(request.src, request.dst)
		if len(errs) > 0 {
			return nil, errInternal
		}

		var routes []*routesResponseRoute

		for _, r := range rr {
			routes = append(routes, &routesResponseRoute{
				Destination: r.Destination,
				Duration:    r.Duration,
				Distance:    r.Distance,
			})
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
		src := entities.LatLng(src[0])
		if !src.Valid() {
			return nil, errSrcIsNotValid
		}

		request.src = src
	}

	if _, ok := values["dst"]; !ok {
		return nil, errDstIsMandatory
	}

	for _, v := range values["dst"] {
		dst := entities.LatLng(v)
		if !dst.Valid() {
			return nil, errDstIsNotValid
		}

		request.dst = append(request.dst, dst)
	}

	return &request, nil
}

func encodeRoutesResponse(_ context.Context, w nethttp.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
