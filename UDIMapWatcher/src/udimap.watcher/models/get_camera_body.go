package models

import "github.com/achiku/xml"

type GetCameraBodyContent struct {
	XMLName xml.Name `xml:"http://udchcmc.local/ getCamera"`
	Key     string   `xml:"key"`
}
