package models

import "github.com/achiku/xml"

type GetTramDoMuaResponse struct {
	XMLName xml.Name `xml:"getDSTramdomuaResponse"`
	Result  string   `xml:"getDSTramdomuaResult"`
}
