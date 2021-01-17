package repositories

import (
	"gorm.io/gorm"
	"telegram-bot-reminder/database/models"
)

type UserRepo struct {
	db *gorm.DB
}

func NewDateUser(db *gorm.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (user UserRepo) Create(u *models.User) *models.User {
	user.db.Create(&u)
	return u
}

func (user UserRepo) Save(u *models.User) *models.User {
	user.db.Save(&u)
	return u
}

func (user UserRepo) FindByUserId(id int32) *models.User {
	var u models.User
	user.db.First(&u, id)
	return &u
}
