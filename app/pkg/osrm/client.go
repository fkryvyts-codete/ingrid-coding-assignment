// Package osrm contains methods for calling OSRM API
package osrm

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/viper"

	"github.com/fkryvyts-codete/ingrid-coding-assignment/pkg/routes/entities"
)

// StatusOk shows whether or not response is successful
const StatusOk = "Ok"

// Client represents API client
type Client interface {
	Driving(src, dst entities.LatLng) (*Response, error)
}

// NewClient returns new API client instance
func NewClient() Client {
	return &client{}
}

// Response represents API response
type Response struct {
	Code   string  `json:"code"`
	Routes []Route `json:"routes"`
}

// Route represents route in API response
type Route struct {
	Duration float32 `json:"duration"`
	Distance float32 `json:"distance"`
}

type client struct {
}

//nolint:gosec
func (c *client) Driving(src, dst entities.LatLng) (*Response, error) {
	url := fmt.Sprintf("%s/route/v1/driving/%s;%s?overview=false", viper.GetString("osrm.url"), src, dst)

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var resp Response

	err = json.Unmarshal(data, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
