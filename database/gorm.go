package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
	"telegram-bot-reminder/database/models"
)

func Database() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	return db, err

}

func autoMigrate(db *gorm.DB) error {
	if os.Getenv("DROP_TABLES") == "yes" {
		_ = db.Migrator().DropTable(&models.User{})
		_ = db.Migrator().DropTable(&models.Date{})
	}

	if os.Getenv("CREATE_TABLE") == "yes" {

	}

	return nil

}
