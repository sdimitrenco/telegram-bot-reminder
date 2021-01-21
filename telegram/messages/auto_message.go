package messages

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"telegram-bot-reminder/database/repositories"
)

func SendMessage(ctx context.Context) {
	db := ctx.Value("db").(*gorm.DB)

	userRepo := repositories.NewUser(db)
	arr := userRepo.GetAllRecords()

	for _, value := range arr {
		chatId := value.ChatId
		fmt.Println(chatId)
	}

	//for {
	//time.Sleep(duration)
	//	duration := time.Second * 4
	//
	//
	//
	//}
}
