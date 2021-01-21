package controlers

import (
	"context"
	"gorm.io/gorm"
	"telegram-bot-reminder/database/models"
	"telegram-bot-reminder/database/repositories"
)

func AddNewUser(ctx context.Context, name string, chatId string) {
	db := ctx.Value("db").(*gorm.DB)
	userRepo := repositories.NewUser(db)

	userRepo.Create(&models.User{
		FullName: name,
		ChatId:   chatId,
	})

}
