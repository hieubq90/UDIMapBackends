package models

type TramDoMua struct {
	ID         int     `gorm:"primary_key" json:"idtram"`
	Name       string  `json:"tentram"`
	Lat        float64 `json:"lat"`
	Lng        float64 `json:"lng"`
	Address    string  `json:"vitri"`
	WaterLevel float64 `json:"mucnuoc"`
	StatusID   int64   `json:"idtrangthai"`
	Status     int64   `json:"status"`
	StatusText string  `json:"trangthai"`
	ShortName  string  `json:"viettat"`
	LastUpdate string  `json:"thoidiem"`
}

// set TramDoMua's table name to be `profiles`
func (TramDoMua) TableName() string {
	return "TramDoMua"
}
