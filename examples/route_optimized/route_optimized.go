package main

import (
	"log"

	valhalla "github.com/Kamion/valhalla-go-bindings"
)

func main() {
	c := valhalla.New("http://localhost:8002")

	request := valhalla.RouteRequest{}
	request.Locations = append(request.Locations, valhalla.Location{
		Lat: 42.015576,
		Lon: -87.841412,
	})

	request.Locations = append(request.Locations, valhalla.Location{
		Lat: 41.818609,
		Lon: -87.623402,
	})

	request.Costing = "auto"
	request.DirectionsOptions.Units = "miles"
	request.DirectionsOptions.Narrative = false
	request.CostingOptions.Auto.CountryCrossingPenalty = 2000.0
	request.ID = "Chicago sightseeing"

	route, err := c.RouteOptimized(request)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%+v\n", route)
}
