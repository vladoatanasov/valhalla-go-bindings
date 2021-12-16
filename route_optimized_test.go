package valhalla

import (
	"testing"
)

func TestRouteOptimized(t *testing.T) {
	type test struct {
		waypoints   []Location
		costing     string
		id          string
		want        float64
		tolerance   float64
		expectError bool
		path        []int
	}

	tests := []test{
		{
			waypoints: []Location{
				{
					Lat:  -41.2924,
					Lon:  174.7787,
					City: "Wellington",
				},
				{
					Lat:  -36.8509,
					Lon:  174.7645,
					City: "Auckland",
				},
			},
			id:          "Wellington -> Auckland",
			costing:     "auto",
			expectError: true,
			want:        643,
			tolerance:   0.05,
		},
		{
			waypoints: []Location{
				{
					Lat:  -41.2924,
					Lon:  174.7787,
					City: "Wellington",
				},
				{
					Lat:  -41.1381,
					Lon:  174.8472,
					City: "Porirua",
				},
				{
					Lat:  -41.2127,
					Lon:  174.8997,
					City: "Lower Hutt",
				},
				{
					Lat:  -41.1249,
					Lon:  175.0656,
					City: "Upper Hutt",
				},
				{
					Lat:  -41.2274,
					Lon:  174.8850,
					City: "Petone",
				},
				{
					Lat:  -41.2924,
					Lon:  174.7787,
					City: "Wellington",
				},
			},
			id:          "Wellington -> [Porirua, Lower Hutt, Upper Hutt, Petone] -> Wellington",
			costing:     "auto",
			expectError: false,
			want:        86,
			path:        []int{0, 4, 2, 3, 1, 5},
			tolerance:   0.05,
		},
	}

	for _, tc := range tests {
		routeRequest := RouteRequest{}
		routeRequest.Costing = tc.costing
		routeRequest.ID = tc.id
		routeRequest.Locations = append(routeRequest.Locations, tc.waypoints...)

		route, err := client.RouteOptimized(routeRequest)

		if tc.expectError {
			if err == nil {
				t.Errorf("Expected an error and did not get it: %s", route.ID)
			}
		} else {
			if err != nil {
				t.Fatalf("Fatal Error: %s %+v", err, tc.id)
			}
			for i, l := range route.Trip.Locations {
				if l.OriginalIndex != tc.path[i] {
					t.Errorf("Expected %d got %d", tc.path[i], l.OriginalIndex)
				}
			}

			if !almostEqual(route.Trip.Summary.Length, tc.want, tc.tolerance) {
				t.Errorf("Route %s failed: Expected %f got %f", tc.id, tc.want, route.Trip.Summary.Length)
			}
		}

	}

}
