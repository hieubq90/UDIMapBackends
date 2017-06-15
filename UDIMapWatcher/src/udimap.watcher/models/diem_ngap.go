package models

//"tendiemngap":null,
//"dongap":0.1,
//"ngay":10,
//"thang":6,
//"nam":2017,
//"idtuyenduong":807,
//"idquan":18,
//"tenduong":"Cay Tram",
//"tenquan":"Gò Vấp",
//"tudoan":"",
//"dendoan":"",
//"iddiemngap":2224,
//"ngaycn":"\/Date(1497083765717)\/",
//"lat":"",
//"lng":"",
//"hinh":"",
//"tinhtrang":1,
//"dukien":"",
//"thoidiem":"10/06/2017 15g36",
//"canhbaongap":"Ngập nhẹ, các phương tiện hạn chế lưu thông"

type DiemNgap struct {
	ID           int64   `json:"iddiemngap"`
	Lat          string  `json:"lat"`
	Lng          string  `json:"lng"`
	RoadName     string  `json:"tenquan"`
	DistrictName string  `json:"tenduong"`
	From         string  `json:"tudoan"`
	To           string  `json:"dendoan"`
	FloodDeep    float64 `json:"dongap"`
	Status       int64   `json:"tinhtrang"`
	Expected     string  `json:"dukien"`
	Warning      string  `json:"canhbaongap"`
	LastUpdate   string  `json:"thoidiem"`
}
