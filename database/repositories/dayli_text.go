package repositories

import (
	"gorm.io/gorm"
	"telegram-bot-reminder/database/models"
)

type DailyText struct {
	db *gorm.DB
}

func NewDailyText(db *gorm.DB) *DailyText {
	return &DailyText{db: db}
}

func (repo DailyText) Create(dt *models.DailyText) *models.DailyText {
	repo.db.Create(&dt)
	return dt
}

func (repo DailyText) FindByDate(date string) *models.DailyText {
	var u models.DailyText
	repo.db.First(&u, "date = ?", date)
	return &u
}
