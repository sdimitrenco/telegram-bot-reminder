package repositories

import (
	"gorm.io/gorm"
	"telegram-bot-reminder/database/models"
)

type Date struct {
	db *gorm.DB
}

func NewDate(db *gorm.DB) *Date {
	return &Date{db: db}
}

func (d Date) Create(date *models.Date) *models.Date {
	d.db.Create(&date)
	return date
}

func (d Date) Save(date *models.Date) *models.Date {
	d.db.Save(&date)
	return date
}

func (d Date) FindByUserId(date string) *models.Date {
	var dt models.Date
	d.db.First(&dt, date)
	return &dt
}
