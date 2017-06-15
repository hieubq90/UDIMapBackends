package models

type FloodPoint struct {
	ID           int64   `gorm:"primary_key" json:"id"`
	Lat          float64 `json:"lat"`
	Lng          float64 `json:"lng"`
	RoadName     string  `json:"tenquan"`
	DistrictName string  `json:"tenduong"`
	From         string  `json:"tudoan"`
	To           string  `json:"dendoan"`
	FloodDeep    float64 `json:"dongap"`
	Status       int64   `json:"tinhtrang"`
	Expected     string  `json:"dukien"`
	Warning      string  `json:"canhbaongap"`
	LastUpdate   string  `json:"thoidiem"`
}
