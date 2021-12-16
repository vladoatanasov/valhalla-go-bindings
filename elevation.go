package valhalla

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type ElevationRequest struct {
	HeightPrecision  *int             `json:"height_precision,omitempty"`
	ID               string           `json:"id"`
	Range            bool             `json:"range"`
	ShapeFormat      string           `json:"shape_format,omitempty"`
	EncodedPolyline  string           `json:"encoded_polyline,omitempty"`
	Shape            []ElevationPoint `json:"shape,omitempty"`
	ResampleDistance *int             `json:"resample_distance,omitempty"`
}

type ElevationPoint struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

type ElevationResponse struct {
	EncodedPolyline string      `json:"encoded_polyline"`
	RangeHeight     [][]float64 `json:"range_height"`
	Height          []float64   `json:"height"`
}

func (c *Client) Height(request ElevationRequest) (ElevationResponse, error) {
	//shape_format
	if len(request.EncodedPolyline) > 0 && len(request.ShapeFormat) == 0 {
		request.ShapeFormat = "polyline6"
	}
	r, err := json.Marshal(request)
	fmt.Printf("%q", r)
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
