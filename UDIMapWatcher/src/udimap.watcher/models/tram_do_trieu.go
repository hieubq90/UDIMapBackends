package models

import (
	"strconv"

	"github.com/jinzhu/gorm"
)

type TramDoTrieu struct {
	ID         int64   `gorm:"primary_key" json:"id,omitempty"`
	IDText     string  `json:"idtramtrieu"`
	Name       string  `json:"tentramtrieu"`
	Lat        float64 `json:"lat"`
	Lng        float64 `json:"lng"`
	Address    string  `json:"vitri"`
	TideLevel  float64 `json:"muctrieu"`
	StatusID   int64   `json:"idtrangthai"`
	Status     int64   `json:"status"`
	StatusText string  `json:"tentrangthai"`
	ShortName  string  `json:"viettat"`
	LastUpdate string  `json:"thoidiem"`
}

// set TramDoTrieu's table name to be `profiles`
func (TramDoTrieu) TableName() string {
	return "TramDoTrieu"
}

func (c *TramDoTrieu) BeforeCreate(scope *gorm.Scope) error {
	id, err := strconv.ParseInt(c.IDText, 10, 64)
	if err == nil {
		scope.SetColumn("ID", id)
		return nil
	}
	return err
}
