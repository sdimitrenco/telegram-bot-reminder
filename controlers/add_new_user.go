package controlers

import (
	"context"
	"gorm.io/gorm"
	"telegram-bot-reminder/database/models"
	"telegram-bot-reminder/database/repositories"
)

func AddNewUser(
	ctx context.Context,
	username string,
	firstName string,
	lastName string,
	chatId string) {

	db := ctx.Value("db").(*gorm.DB)
	userRepo := repositories.NewUser(db)

	user := userRepo.FindByChatId(chatId)

	if user.ChatId == "" {
		userRepo.Create(&models.User{
			Username:  username,
			FirstName: firstName,
			LastName:  lastName,
			ChatId:    chatId,
		})
	}

}
