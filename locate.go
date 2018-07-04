package valhalla

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type LocateResponse []struct {
	InputLon float64 `json:"input_lon"`
	InputLat float64 `json:"input_lat"`
	Nodes    []struct {
		TrafficSignal bool    `json:"traffic_signal"`
		Type          string  `json:"type"`
		Lat           float64 `json:"lat"`
		NodeID        struct {
			ID     int   `json:"id"`
			Value  int64 `json:"value"`
			TileID int   `json:"tile_id"`
			Level  int   `json:"level"`
		} `json:"node_id"`
		Access struct {
			Wheelchair bool `json:"wheelchair"`
			Taxi       bool `json:"taxi"`
			HOV        bool `json:"HOV"`
			Truck      bool `json:"truck"`
			Emergency  bool `json:"emergency"`
			Pedestrian bool `json:"pedestrian"`
			Car        bool `json:"car"`
			Bus        bool `json:"bus"`
			Bicycle    bool `json:"bicycle"`
		} `json:"access"`
		Lon            float64 `json:"lon"`
		EdgeCount      int     `json:"edge_count"`
		Administrative struct {
			TimeZonePosix               string `json:"time_zone_posix"`
			StandardTimeZoneName        string `json:"standard_time_zone_name"`
			Iso31661                    string `json:"iso_3166-1"`
			DaylightSavingsTimeZoneName string `json:"daylight_savings_time_zone_name"`
			Country                     string `json:"country"`
			Iso31662                    string `json:"iso_3166-2"`
			State                       string `json:"state"`
		} `json:"administrative"`
		IntersectionType string `json:"intersection_type"`
		Density          int    `json:"density"`
		LocalEdgeCount   int    `json:"local_edge_count"`
		ModeChange       bool   `json:"mode_change"`
	} `json:"nodes"`
	Edges []struct {
		EdgeID struct {
			ID     int   `json:"id"`
			Value  int64 `json:"value"`
			TileID int   `json:"tile_id"`
			Level  int   `json:"level"`
		} `json:"edge_id"`
		EdgeInfo struct {
			Names []string `json:"names"`
			Shape string   `json:"shape"`
			WayID int      `json:"way_id"`
		} `json:"edge_info"`
		Edge struct {
			Classification struct {
				Link           bool   `json:"link"`
				Internal       bool   `json:"internal"`
				Surface        string `json:"surface"`
				Classification string `json:"classification"`
			} `json:"classification"`
			GeoAttributes struct {
				Curvature     int     `json:"curvature"`
				WeightedGrade float32 `json:"weighted_grade"`
				Length        int     `json:"length"`
			} `json:"geo_attributes"`
			SpeedType   string `json:"speed_type"`
			BikeNetwork struct {
				Mountain bool `json:"mountain"`
				Local    bool `json:"local"`
				Regional bool `json:"regional"`
				National bool `json:"national"`
			} `json:"bike_network"`
			CycleLane     string `json:"cycle_lane"`
			TrafficSignal bool   `json:"traffic_signal"`
			Unreachable   bool   `json:"unreachable"`
			RoundAbout    bool   `json:"round_about"`
			Access        struct {
				Truck      bool `json:"truck"`
				Pedestrian bool `json:"pedestrian"`
				Wheelchair bool `json:"wheelchair"`
				Taxi       bool `json:"taxi"`
				HOV        bool `json:"HOV"`
				Emergency  bool `json:"emergency"`
				Motorcycle bool `json:"motorcycle"`
				Car        bool `json:"car"`
				Moped      bool `json:"moped"`
				Bus        bool `json:"bus"`
				Bicycle    bool `json:"bicycle"`
			} `json:"access"`
			DestinationOnly          bool   `json:"destination_only"`
			CountryCrossing          bool   `json:"country_crossing"`
			Forward                  bool   `json:"forward"`
			Seasonal                 bool   `json:"seasonal"`
			Use                      string `json:"use"`
			DriveOnRight             bool   `json:"drive_on_right"`
			LaneCount                int    `json:"lane_count"`
			TruckRoute               bool   `json:"truck_route"`
			PartOfComplexRestriction bool   `json:"part_of_complex_restriction"`
			StartRestriction         struct {
				Truck      bool `json:"truck"`
				Pedestrian bool `json:"pedestrian"`
				Wheelchair bool `json:"wheelchair"`
				Taxi       bool `json:"taxi"`
				HOV        bool `json:"HOV"`
				Emergency  bool `json:"emergency"`
				Motorcycle bool `json:"motorcycle"`
				Car        bool `json:"car"`
				Moped      bool `json:"moped"`
				Bus        bool `json:"bus"`
				Bicycle    bool `json:"bicycle"`
			} `json:"start_restriction"`
			Toll              bool `json:"toll"`
			HasExitSign       bool `json:"has_exit_sign"`
			AccessRestriction bool `json:"access_restriction"`
			Bridge            bool `json:"bridge"`
			Tunnel            bool `json:"tunnel"`
			SpeedLimit        int  `json:"speed_limit"`
			NotThru           bool `json:"not_thru"`
			EndRestriction    struct {
				Truck      bool `json:"truck"`
				Pedestrian bool `json:"pedestrian"`
				Wheelchair bool `json:"wheelchair"`
				Taxi       bool `json:"taxi"`
				HOV        bool `json:"HOV"`
				Emergency  bool `json:"emergency"`
				Motorcycle bool `json:"motorcycle"`
				Car        bool `json:"car"`
				Moped      bool `json:"moped"`
				Bus        bool `json:"bus"`
				Bicycle    bool `json:"bicycle"`
			} `json:"end_restriction"`
			Speed   int `json:"speed"`
			EndNode struct {
				ID     int   `json:"id"`
				Value  int64 `json:"value"`
				TileID int   `json:"tile_id"`
				Level  int   `json:"level"`
			} `json:"end_node"`
		} `json:"edge"`
		MinimumReachability int           `json:"minimum_reachability"`
		TrafficSegments     []interface{} `json:"traffic_segments"`
		Distance            float64       `json:"distance"`
		PercentAlong        float64       `json:"percent_along"`
		CorrelatedLon       float64       `json:"correlated_lon"`
		SideOfStreet        string        `json:"side_of_street"`
		CorrelatedLat       float64       `json:"correlated_lat"`
	} `json:"edges"`
}

func (c *Client) Locate(request RouteRequest) (LocateResponse, error) {
	r, err := json.Marshal(request)
	if err != nil {
		return LocateResponse{}, err
	}

	response, status, err := c.request("GET", "locate", bytes.NewReader(r))
	if err != nil {
		return LocateResponse{}, err
	}

	if status >= 400 {
		result := ValhallaError{}
		err = json.Unmarshal(response, &result)
		if err != nil {
			return LocateResponse{}, err
		}

		return LocateResponse{}, fmt.Errorf("%+v", result)
	}

	result := LocateResponse{}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return LocateResponse{}, err
	}

	return result, nil
}
