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

func (repo Date) Create(d *models.Date) *models.Date {
	repo.db.Create(&d)
	return d
}

func (repo Date) Save(d *models.Date) *models.Date {
	repo.db.Save(&d)
	return d
}

func (repo Date) FindByUserId(date string) *models.Date {
	var d models.Date
	repo.db.First(&d, date)
	return &d
}
