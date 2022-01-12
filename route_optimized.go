package valhalla

import (
	"bytes"

	easyjson "github.com/mailru/easyjson"
)

func (c *Client) RouteOptimized(request RouteRequest) (RouteResponse, error) {
	r, err := easyjson.Marshal(request)
	if err != nil {
		return RouteResponse{}, err
	}

	response, err := c.request("GET", "optimized_route", bytes.NewReader(r))
	if err != nil {
		return RouteResponse{}, err
	}

	result := RouteResponse{}
	err = easyjson.Unmarshal(response, &result)
	if err != nil {
		return RouteResponse{}, err
	}

	return result, nil
}
