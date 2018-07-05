package valhalla

import (
	"bytes"
	"encoding/json"

	"github.com/paulmach/orb/geojson"
)

func (c *Client) Isochrone(request RouteRequest) (*geojson.FeatureCollection, error) {
	r, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	response, err := c.request("GET", "isochrone", bytes.NewReader(r))
	if err != nil {
		return nil, err
	}

	result, err := geojson.UnmarshalFeatureCollection(response)
	if err != nil {
		return nil, err
	}

	return result, nil
}
