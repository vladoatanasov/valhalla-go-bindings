package valhalla

type CostingOptions struct {
	Auto CostingOption `json:"auto,omitempty"`
}

type CostingOption struct {
	CountryCrossingPenalty float32 `json:"country_crossing_penalty,omitempty"`
	TollBoothPenalty       float32 `json:"toll_booth_penalty,omitempty"`
}
