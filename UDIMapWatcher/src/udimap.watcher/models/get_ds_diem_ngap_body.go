package models

import "github.com/achiku/xml"

type GetDSDiemNgapBodyContent struct {
	XMLName xml.Name `xml:"http://udchcmc.local/ getDSDiemNgap"`
	Key     string   `xml:"key"`
}
