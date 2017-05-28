package models

import "github.com/achiku/xml"

type GetTramDoMuaBodyContent struct {
	XMLName xml.Name `xml:"http://udchcmc.local/ getDSTramdomua"`
	Key     string   `xml:"key"`
}
