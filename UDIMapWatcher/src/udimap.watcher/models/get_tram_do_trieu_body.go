package models

import "github.com/achiku/xml"

type GetTramDoTrieuBodyContent struct {
	XMLName xml.Name `xml:"http://udchcmc.local/ getDSTramdotrieu"`
	Key     string   `xml:"key"`
}
