package repositories

import (
	"gorm.io/gorm"
	"telegram-bot-reminder/database/models"
)

type DateRepo struct {
	db *gorm.DB
}

func NewDateRepo(db *gorm.DB) *DateRepo {
	return &DateRepo{db: db}
}

func (repo DateRepo) Create(d *models.Date) *models.Date {
	repo.db.Create(&d)
	return d
}

func (repo DateRepo) Save(d *models.Date) *models.Date {
	repo.db.Save(&d)
	return d
}

func (repo DateRepo) FindByUserId(date string) *models.Date {
	var d models.Date
	repo.db.First(&d, date)
	return &d
}
