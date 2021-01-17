package models

type User struct {
	Id       int32 `gorm:"primaryKey"`
	FullName string
}

func (u *User) GetFullName() string {
	return u.FullName
}
