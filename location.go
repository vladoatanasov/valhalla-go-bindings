package valhalla

type Location struct {
	Lat    float32 `json:"lat,omitempty"`
	Lon    float32 `json:"lon,omitempty"`
	Street string  `json:"street,omitempty"`
}
