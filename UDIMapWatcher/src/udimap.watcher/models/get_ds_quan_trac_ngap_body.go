package models

import "github.com/achiku/xml"

type GetDSQuanTracNgapBodyContent struct {
	XMLName xml.Name `xml:"http://udchcmc.local/ getDSQuantracngap"`
	Key     string   `xml:"key"`
}
