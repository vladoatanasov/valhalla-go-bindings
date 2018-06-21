package valhalla

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func (c *Client) RouteOptimized(request RouteRequest) (RouteResponse, error) {
	r, err := json.Marshal(request)
	if err != nil {
		return RouteResponse{}, err
	}

	response, status, err := c.request("GET", "optimized_route", bytes.NewReader(r))
	if err != nil {
		return RouteResponse{}, err
	}

	if status >= 400 {
		result := ValhallaError{}
		err = json.Unmarshal(response, &result)
		if err != nil {
			return RouteResponse{}, err
		}

		return RouteResponse{}, fmt.Errorf("%+v", result)
	}

	result := RouteResponse{}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return RouteResponse{}, err
	}

	return result, nil
}
