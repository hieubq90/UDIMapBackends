package models

import "github.com/achiku/xml"

type GetTramDoTrieuResponse struct {
	XMLName xml.Name `xml:"getDSTramdotrieuResponse"`
	Result  string   `xml:"getDSTramdotrieuResult"`
}
