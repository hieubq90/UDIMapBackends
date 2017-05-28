package models

import "github.com/achiku/xml"

type GetCameraResponse struct {
	XMLName xml.Name `xml:"getCameraResponse"`
	Result  string   `xml:"getCameraResult"`
}
