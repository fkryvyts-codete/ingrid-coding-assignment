// Package http contains logic for handling http requests and translating them to service method calls
package http

import (
	"context"
	"encoding/json"
	"fmt"
	nethttp "net/http"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"

	"github.com/fkryvyts-codete/ingrid-coding-assignment/pkg/routes/entities"
)

var (
	errSrcIsMandatory    = newHTTPError("query parameter \"src\" is mandatory", 400)
	errMultipleSrcValues = newHTTPError("multiple values for query parameter \"src\" are not supported", 400)
	errDstIsMandatory    = newHTTPError("query parameter \"dst\" is mandatory", 400)
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
	return func(_ context.Context, req interface{}) (interface{}, error) {
		fmt.Println(req)

		return routesResponse{
			Routes: []entities.Route{},
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

	return request, nil
}

func encodeRoutesResponse(_ context.Context, w nethttp.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
