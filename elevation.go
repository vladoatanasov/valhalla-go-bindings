package valhalla

import (
	"bytes"
	"encoding/json"
)

type ElevationRequest struct {
	HeightPrecision  int    `json:"height_precision"`
	ID               string `json:"id"`
	Range            bool   `json:"range"`
	ShapeFormat      string `json:"shape_format"`
	EncodedPolyline  string `json:"encoded_polyline"`
	ResampleDistance int    `json:"resample_distance"`
}

type ElevationResponse struct {
	EncodedPolyline string      `json:"encoded_polyline"`
	RangeHeight     [][]float64 `json:"range_height"`
	Height          []float64   `json:"height"`
}

func (c *Client) Height(request ElevationRequest) (ElevationResponse, error) {
	r, err := json.Marshal(request)
	if err != nil {
		return ElevationResponse{}, err
	}

	response, err := c.request("GET", "height", bytes.NewReader(r))
	if err != nil {
		return ElevationResponse{}, err
	}

	result := ElevationResponse{}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return ElevationResponse{}, err
	}

	return result, nil
}
