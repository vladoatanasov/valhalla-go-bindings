package valhalla

type Location struct {
	Lat                 float64 `json:"lat,omitempty"`
	Lon                 float64 `json:"lon,omitempty"`
	Street              string  `json:"street,omitempty"`
	Type                string
	Heading             int
	HeadingTolerance    int
	MinimumReachability int
	Radius              int
	RankCandidates      bool
	Name                string
	City                string
	State               string
	PostalCode          string
	Country             string
	Phone               string
	URL                 string
	SideOfStreet        string //left, right
	DateTime            string //2015-12-29T08:00
}
