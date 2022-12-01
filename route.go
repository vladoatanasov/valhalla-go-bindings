package valhalla

import (
	"bytes"

	easyjson "github.com/mailru/easyjson"
)

type RouteRequest struct {
	Locations         []Location        `json:"locations,omitempty"`
	Costing           string            `json:"costing,omitempty"`
	DirectionsOptions DirectionsOptions `json:"directions_options,omitempty"`
	ID                string            `json:"id,omitempty"`
	CostingOptions    CostingOptions    `json:"costing_options,omitempty"`
	AvoidLocations    []Location        `json:"avoid_locations,omitempty"`
	Verbose           bool              `json:"verbose,omitempty"`

	// isochrone request
	Contours   []Contour `json:"contours,omitempty"`
	Polygons   bool      `json:"polygons,omitempty"`
	Denoise    float32   `json:"denoise,omitempty"`
	Generalize float32   `json:"generalize,omitempty"`
}

type Contour struct {
	Time  int    `json:"time,omitempty"`
	Color string `json:"color,omitempty"`
}

type RouteResponse struct {
	ID   string `json:"id"`
	Trip struct {
		Language      string `json:"language"`
		Status        int    `json:"status"`
		Units         string `json:"units"`
		StatusMessage string `json:"status_message"`
		Legs          []struct {
			Shape   string `json:"shape"`
			Summary struct {
				MaxLon float64 `json:"max_lon"`
				MaxLat float64 `json:"max_lat"`
				Time   float64 `json:"time"`
				Length float64 `json:"length"`
				MinLat float64 `json:"min_lat"`
				MinLon float64 `json:"min_lon"`
			} `json:"summary"`
			Maneuvers []Maneuver `json:"maneuvers"`
		} `json:"legs"`
		Summary struct {
			MaxLon float64 `json:"max_lon"`
			MaxLat float64 `json:"max_lat"`
			Time   float64 `json:"time"`
			Length float64 `json:"length"`
			MinLat float64 `json:"min_lat"`
			MinLon float64 `json:"min_lon"`
		} `json:"summary"`
		Locations []struct {
			OriginalIndex int     `json:"original_index"`
			Type          string  `json:"type"`
			Lon           float64 `json:"lon"`
			Lat           float64 `json:"lat"`
			SideOfStreet  string  `json:"side_of_street"`
		} `json:"locations"`
	} `json:"trip"`
	Alternates []RouteResponse `json:"alternates"`
}

type Maneuver struct {
	TravelMode                       string   `json:"travel_mode"`
	BeginShapeIndex                  int      `json:"begin_shape_index"`
	Length                           float64  `json:"length"`
	Time                             float64  `json:"time"`
	Type                             int      `json:"type"`
	EndShapeIndex                    int      `json:"end_shape_index"`
	Instruction                      string   `json:"instruction"`
	VerbalPreTransitionInstruction   string   `json:"verbal_pre_transition_instruction"`
	TravelType                       string   `json:"travel_type"`
	StreetNames                      []string `json:"street_names,omitempty"`
	VerbalTransitionAlertInstruction string   `json:"verbal_transition_alert_instruction,omitempty"`
	VerbalPostTransitionInstruction  string   `json:"verbal_post_transition_instruction,omitempty"`
	Sign                             struct {
		ExitBranchElements []struct {
			Text string `json:"text"`
		} `json:"exit_branch_elements"`
	} `json:"sign,omitempty"`
}

type ValhallaError struct {
	ErrorCode  int    `json:"error_code"`
	Error      string `json:"error"`
	StatusCode int    `json:"status_code"`
	Status     string `json:"status"`
}

func (c *Client) Route(request RouteRequest) (RouteResponse, error) {
	r, err := easyjson.Marshal(request)
	if err != nil {
		return RouteResponse{}, err
	}

	response, err := c.request("GET", "route", bytes.NewReader(r))
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
