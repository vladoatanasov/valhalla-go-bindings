package valhalla

type Location struct {
	Lat    float64 `json:"lat,omitempty"`
	Lon    float64 `json:"lon,omitempty"`
	Street string  `json:"street,omitempty"`
}
