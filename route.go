package valhalla

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type RouteRequest struct {
	Locations         []Location        `json:"locations,omitempty"`
	Costing           string            `json:"costing,omitempty"`
	DirectionsOptions DirectionsOptions `json:"directions_options,omitempty"`
	ID                string            `json:"id,omitempty"`
}

func (c *Client) Route(request RouteRequest) error {
	r, err := json.Marshal(request)
	if err != nil {
		return err
	}

	response, err := c.request("GET", "route", bytes.NewReader(r))
	if err != nil {
		return err
	}

	fmt.Println(string(response))

	return nil
}
