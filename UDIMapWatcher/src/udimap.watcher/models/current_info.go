package models

type CurrentInfo struct {
	Err             int32          `json:"err"`
	ListCameras     []*Camera      `json:"list_cameras"`
	ListRainTracker []*TramDoMua   `json:"list_rain_tracker"`
	ListTideTracker []*TramDoTrieu `json:"list_tide_tracker"`
}
