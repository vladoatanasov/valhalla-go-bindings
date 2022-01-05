package valhalla

import (
	"bytes"
	"encoding/json"
)

type TraceRouteRequest struct {
	BeginTime     *float64 `json:"begin_time,omitempty"`
	Durations     []int64  `json:"durations,omitempty"`
	UseTimestamps bool     `json:"use_timestamps"`
	Costing       string   `json:"costing,omitempty"`
	TraceOptions  struct {
		SearchRadius          *float64 `json:"search_radius,omitempty"`
		GPSAccuracy           *float64 `json:"gps_accuracy,omitempty"`
		BreakageDistance      *float64 `json:"breakage_distance,omitempty"`
		InterpolationDistance *float64 `json:"interpolation_distance,omitempty"`
	} `json:"trace_options,omitempty"`
	LinearReferences bool   `json:"linear_references"`
	EncodedPolyline  string `json:"encoded_polyline"`
}

func (c *Client) TraceRoute(request TraceRouteRequest) (RouteResponse, error) {
	r, err := json.Marshal(request)
	result := RouteResponse{}

	if err != nil {
		return result, err
	}
	response, err := c.request("POST", "trace_route", bytes.NewReader(r))
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}
