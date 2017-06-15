package models

import "github.com/achiku/xml"

type GetDSQuanTracNgapResponse struct {
	XMLName xml.Name `xml:"getDSQuantracngapResponse"`
	Result  string   `xml:"getDSQuantracngapResult"`
}
