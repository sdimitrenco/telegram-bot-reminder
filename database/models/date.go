package models

type Date struct {
	Id         int32 `gorm:"primaryKey"`
	StringDate string
	UserID     int32
}

func (d *Date) GetStringDate() string {
	return d.StringDate
}

func (d *Date) GetUserId() int32 {
	return d.UserID
}
