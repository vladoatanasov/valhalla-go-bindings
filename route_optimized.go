package valhalla

import (
	"bytes"
	"encoding/json"
)

func (c *Client) RouteOptimized(request RouteRequest) (RouteResponse, error) {
	r, err := json.Marshal(request)
	if err != nil {
		return RouteResponse{}, err
	}

	response, err := c.request("GET", "optimized_route", bytes.NewReader(r))
	if err != nil {
		return RouteResponse{}, err
	}

	result := RouteResponse{}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return RouteResponse{}, err
	}

	return result, nil
}
