package valhalla

type CostingOptions struct {
	Auto CostingOption `json:"auto,omitempty"`
}

type CostingOption struct {
	CountryCrossingPenalty int `json:"country_crossing_penalty,omitempty"`
	TollBoothPenalty       int `json:"toll_booth_penalty,omitempty"`
	CountryCrossingCost    int `json:"country_crossing_cost,omitempty"`
}
