package models

//{
//"idtram":"1037",
//"tentram":"Huỳnh Tấn Phát",
//"vitri":"873 Huỳnh Tấn Phát, phường Phú Thuận",
//"lat":"10.73129800",
//"lng":"106.73192500",
//"dosaungap":0.16,
//"idtrangthai":1,
//"status":1,
//"tentrangthai":"Tăng",
//"thoidiem":"10/06/2017 15g40"
//}

type TramQuanTracNgap struct {
	ID         int64   `gorm:"primary_key" json:"id"`
	Name       string  `json:"tentram"`
	Lat        float64 `json:"lat"`
	Lng        float64 `json:"lng"`
	Address    string  `json:"vitri"`
	FloodDeep  float64 `json:"dosaungap"`
	StatusID   int64   `json:"idtrangthai"`
	Status     int64   `json:"status"`
	StatusText string  `json:"tentrangthai"`
	LastUpdate string  `json:"thoidiem"`
}

// set TramDoMua's table name to be `profiles`
func (TramQuanTracNgap) TableName() string {
	return "TramQuanTracNgap"
}
