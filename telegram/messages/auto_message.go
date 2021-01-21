package messages

import (
	"context"
	"fmt"
	"github.com/yanzay/tbot"
	"gorm.io/gorm"
	"os"
	"telegram-bot-reminder/database/repositories"
	"time"
)

func SendMessage(ctx context.Context, client *tbot.Client) {
	db := ctx.Value("db").(*gorm.DB)

	for {
		duration := time.Hour * 10
		time.Sleep(duration)

		userRepo := repositories.NewUser(db)
		arr := userRepo.GetAllRecords()

		for _, value := range arr {
			chatId := value.ChatId
			_, _ = client.SendMessage(chatId, os.Getenv("AUTO_TEXT"), tbot.OptParseModeMarkdown)
			fmt.Println(chatId)
		}

	}
}
