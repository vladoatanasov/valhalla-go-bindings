package main

import (
	"log"

	valhalla "github.com/littlemonkeyltd/valhalla-go-bindings"
)

func main() {
	c := valhalla.New("http://localhost:8002")

	request := valhalla.RouteRequest{}
	request.Verbose = true
	request.Locations = append(request.Locations, valhalla.Location{
		Lat: 41.885619,
		Lon: -87.621311,
	})

	request.Costing = "auto"
	request.DirectionsOptions.Units = "miles"
	request.DirectionsOptions.Narrative = false
	request.CostingOptions.Auto.CountryCrossingPenalty = 2000.0
	request.ID = "Chicago sightseeing"

	route, err := c.Locate(request)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%+v\n", route)
}
