package models

type Camera struct {
	ID       string `gorm:"primary_key"`
	Address  string
	ImageUrl string
	Lat      float64
	Lng      float64
}

// set Camera's table name to be `profiles`
func (Camera) TableName() string {
	return "CameraNgap"
}

//func (c *Camera) BeforeCreate(scope *gorm.Scope) error {
//	scope.SetColumn("ID", fmt.Sprintf("%f_%f", c.Lat, c.Lng))
//	return nil
//}
