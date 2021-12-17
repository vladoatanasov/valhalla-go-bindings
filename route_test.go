package valhalla

import (
	"math"
	"testing"
)

var client = New("http://localhost:8002")

func almostEqual(a, b, tolerance float64) bool {
	threshold := a * tolerance
	return math.Abs(a-b) <= threshold
}
func TestRoute(t *testing.T) {
	type test struct {
		waypoints   []Location
		costing     string
		id          string
		want        float64
		tolerance   float64
		expectError bool
		skip        bool
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
			expectError: false,
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
					Lat:  -41.4985,
					Lon:  173.2444,
					City: "Nelson",
				},
			},
			id:          "Wellington -> Nelson",
			costing:     "auto",
			expectError: false,
			want:        250,
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
					Lat:  -33.8688,
					Lon:  151.2093,
					City: "Sydney",
				},
			},
			id:          "Wellington -> Sydney",
			costing:     "auto",
			expectError: true,
			want:        238,
			tolerance:   0.05,
		},
		{
			waypoints: []Location{
				{
					Lat:  -31.9523,
					Lon:  115.8613,
					City: "Perth",
				},
				{
					Lat:  -33.8688,
					Lon:  151.2093,
					City: "Sydney",
				},
			},
			id:          "Perth -> Sydney",
			costing:     "auto",
			expectError: false,
			want:        3933,
			tolerance:   0.05,
		},
		{
			waypoints: []Location{
				{
					Lat:  -31.9523,
					Lon:  115.8613,
					City: "Perth",
				},

				{
					Lat:  -33.8688,
					Lon:  151.2093,
					City: "Sydney",
				},
				{
					Lat:  -16.9203,
					Lon:  145.7710,
					City: "Cairns",
				},
			},
			id:          "Perth -> Sydney -> Cairns",
			costing:     "auto",
			expectError: true,
			want:        6346,
			tolerance:   0.05,
		},
		{
			waypoints: []Location{
				{
					Lat:  -21.2129,
					Lon:  -159.7823,
					City: "Avarua",
				},
				{
					Lat:  -21.2338709,
					Lon:  -159.8231969,
					City: "Wigmore's Superstore",
				},
			},
			id:          "Avarua -> Wigmore's Superstore",
			costing:     "auto",
			expectError: false,
			want:        13,
			tolerance:   0.05,
			skip:        true,
		},
	}

	for _, tc := range tests {
		if tc.skip {
			t.Logf("Skipped %s", tc.id)
			continue
		}
		routeRequest := RouteRequest{}
		routeRequest.Costing = tc.costing
		routeRequest.ID = tc.id
		routeRequest.Locations = append(routeRequest.Locations, tc.waypoints...)

		route, err := client.Route(routeRequest)

		if tc.expectError {
			if err == nil {
				t.Errorf("Expected an error and did not get it trip %s", route.ID)
			}
		} else {
			if err != nil {
				t.Fatalf("Fatal Error: %+v", err)
			}
			if !almostEqual(route.Trip.Summary.Length, tc.want, tc.tolerance) {
				t.Errorf("Route %s failed: Expected %fkm got %fkm", tc.id, tc.want, route.Trip.Summary.Length)
			}
		}

	}

}

func BenchmarkRoute(b *testing.B) {
	routeRequest := RouteRequest{}
	routeRequest.Costing = "auto"
	routeRequest.Locations = []Location{
		{
			Lat:  -31.9523,
			Lon:  115.8613,
			City: "Perth",
		},
		{
			Lat:  -33.8688,
			Lon:  151.2093,
			City: "Sydney",
		},
	}
	for i := 0; i < b.N; i++ {
		client.Route(routeRequest)
	}
}
