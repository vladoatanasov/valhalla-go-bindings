package valhalla

type CostingOptions struct {
	Auto  CostingOption `json:"auto,omitempty"`
	Truck CostingOption `json:"truck,omitempty"`
}

type CostingOption struct {
	CountryCrossingPenalty int `json:"country_crossing_penalty,omitempty"`
	TollBoothPenalty       int `json:"toll_booth_penalty,omitempty"`
	CountryCrossingCost    int `json:"country_crossing_cost,omitempty"`

	// truck specific costing options
	Hazmat   bool    `json:"hazmat,omitempty"`
	Weight   float32 `json:"weight,omitempty"`
	AxleLoad float32 `json:"axle_load,omitempty"`
	Height   float32 `json:"height,omitempty"`
	Width    float32 `json:"width,omitempty"`
	Length   float32 `json:"length,omitempty"`
}
