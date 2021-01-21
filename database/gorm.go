package database

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
	"telegram-bot-reminder/database/models"
)

//connect to database
func Database() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	autoMigrate(db)

	return db, err

}

//create or drop tables
func autoMigrate(db *gorm.DB) {
	if os.Getenv("DROP_TABLES") == "yes" {
		_ = db.Migrator().DropTable(&models.User{})
		_ = db.Migrator().DropTable(&models.Date{})
		_ = db.Migrator().DropTable(&models.DailyText{})
	}

	if os.Getenv("CREATE_TABLE") == "yes" {
		if err := db.AutoMigrate(&models.User{}); err != nil {
			fmt.Println(err)
		}
		if err := db.AutoMigrate(&models.Date{}); err != nil {
			fmt.Println(err)
		}
		if err := db.AutoMigrate(&models.DailyText{}); err != nil {
			fmt.Println(err)
		}
	}

}
