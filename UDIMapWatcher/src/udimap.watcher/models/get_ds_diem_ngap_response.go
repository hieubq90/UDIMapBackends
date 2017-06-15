package models

import "github.com/achiku/xml"

type GetDSDiemNgapResponse struct {
	XMLName xml.Name `xml:"getDSDiemNgapResponse"`
	Result  string   `xml:"getDSDiemNgapResult"`
}
