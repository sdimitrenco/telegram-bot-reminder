package models

type User struct {
	Id       int32 `gorm:"primaryKey"`
	FullName string
	ChatId   string
}

func (u *User) GetFullName() string {
	return u.FullName
}

func (u *User) GetChatId() string {
	return u.ChatId
}
