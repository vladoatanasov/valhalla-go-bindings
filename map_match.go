package valhalla

import (
	"bytes"
	"encoding/json"

	easyjson "github.com/mailru/easyjson"
)

type TraceRouteRequest struct {
	BeginTime         *float64          `json:"begin_time,omitempty"`
	Durations         []int64           `json:"durations,omitempty"`
	UseTimestamps     bool              `json:"use_timestamps"`
	Costing           string            `json:"costing,omitempty"`
	DirectionsOptions DirectionsOptions `json:"directions_options,omitempty"`
	TraceOptions      TraceOptions      `json:"trace_options,omitempty"`
	LinearReferences  bool              `json:"linear_references"`
	EncodedPolyline   string            `json:"encoded_polyline"`
	ShapeMatch        string            `json:"shape_match,omitempty"`
}
type TraceOptions struct {
	SearchRadius          *float64 `json:"search_radius,omitempty"`
	GPSAccuracy           *float64 `json:"gps_accuracy,omitempty"`
	BreakageDistance      *float64 `json:"breakage_distance,omitempty"`
	InterpolationDistance *float64 `json:"interpolation_distance,omitempty"`
}
type TraceAttributesRequest struct {
	EncodedPolyline string       `json:"encoded_polyline"`
	Costing         string       `json:"costing,omitempty"`
	TraceOptions    TraceOptions `json:"trace_options,omitempty"`
	ShapeMatch      string       `json:"shape_match,omitempty"`
}
type TraceAttributesResponse struct {
	Edges           []Edge  `json:"edges"`
	OSMChangeSet    int     `json:"osm_changeset"`
	ConfidenceScore float64 `json:"confidence_score"`
	Shape           string  `json:"shape"`
	Units           string  `json:"units"`
	Admins          []struct {
		StateCode   string `json:"state_code"`
		StateText   string `json:"state_text"`
		CountryText string `json:"country_text"`
		CountryCode string `json:"country_code"`
	} `json:"admins"`
}

type Edge struct {
	Names                []string `json:"names"`
	Length               float64  `json:"length"`
	Speed                float64  `json:"speed"`
	RoadClass            string   `json:"road_class"`
	BeginHeading         float64  `json:"begin_heading"`
	EndHeading           float64  `json:"end_heading"`
	BeginShapeIndex      int      `json:"begin_shape_index"`
	EndShapeIndex        int      `json:"end_shape_index"`
	Traversability       string   `json:"traversability"`
	Use                  string   `json:"use"`
	Toll                 bool     `json:"toll"`
	Unpaved              bool     `json:"unpaved"`
	Tunnel               bool     `json:"tunnel"`
	Bridge               bool     `json:"bridge"`
	Roundabout           bool     `json:"roundabout"`
	InternalIntersection bool     `json:"internal_intersection"`
	DriveOnRight         bool     `json:"drive_on_right"`
	Surface              string   `json:"surface"`
	Sign                 struct {
		ExitNumber []string `json:"exit_number"`
		ExitBranch []string `json:"exit_branch"`
		ExitToward []string `json:"exit_toward"`
		ExitName   []string `json:"exit_name"`
	} `json:"sign"`
	TravelMode       string  `json:"travel_mode"`
	VehicleType      string  `json:"vehicle_type"`
	PedestrianType   string  `json:"pedestrian_type"`
	BicycleType      string  `json:"bicycle_type"`
	TransitType      string  `json:"transit_type"`
	ID               int     `json:"id"`
	WayID            int     `json:"way_id"`
	WeightedGrade    float64 `json:"weighted_grade"`
	MaxUpwardGrade   float64 `json:"max_upward_grade"`
	MaxDownwardGrade float64 `json:"max_downward_grade"`
	MeanElevation    float64 `json:"mean_elevation"`
	LaneCount        int     `json:"lane_count"`
	CycleLane        string  `json:"cycle_lane"`
	BicycleNetwork   int     `json:"bicycle_network"`
	SACScale         int     `json:"sac_scale"`
	Shoulder         bool    `json:"shoulder"`
	Sidewalk         string  `json:"sidewalk"`
	Density          float64 `json:"density"`
	SpeedLimit       float64 `json:"speed_limit"`
	TruckSpeed       float64 `json:"truck_speed"`
	TruckRoute       bool    `json:"truck_route"`
	EndNode          struct {
		IntersectingEdges []struct {
			BeginHeading            float64 `json:"begin_heading"`
			FromEdgeNameConsistency bool    `json:"from_edge_name_consistency"`
			ToEdgeNameConsistency   bool    `json:"to_edge_name_consistency"`
			Driveability            string  `json:"drivability"`
			Cyclability             string  `json:"cyclability"`
			Walkability             string  `json:"walkability"`
			Use                     string  `json:"use"`
			RoadClass               string  `json:"road_class"`
		} `json:"intersecting_edges"`
		ElapsedTime float64 `json:"elapsed_time"`
		AdminIndex  int     `json:"admin_index"`
		Type        string  `json:"type"`
		Fork        bool    `json:"fork"`
		TimeZone    string  `json:"time_zone"`
	} `json:"end_node"`
}

func (c *Client) TraceRoute(request TraceRouteRequest) (RouteResponse, error) {
	r, err := easyjson.Marshal(request)
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
func (c *Client) TraceAttributes(request TraceAttributesRequest) (TraceAttributesResponse, error) {
	r, err := easyjson.Marshal(request)
	result := TraceAttributesResponse{}

	if err != nil {
		return result, err
	}
	response, err := c.request("POST", "trace_attributes", bytes.NewReader(r))
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}
