package models

type Favourite struct {
	Id       int64
	UserId   int64
	Name     string
	Lon      float64
	Lat      float64
	Selected bool
}
