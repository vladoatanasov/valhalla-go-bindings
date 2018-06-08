package valhalla

type DirectionsOptions struct {
	Units     string `json:"units,omitempty"`
	Narrative bool   `json:"narrative,omitempty"`
	Language  string `json:"language,omitempty"`
}
