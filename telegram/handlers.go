package telegram

import (
	"context"
	"fmt"
	"github.com/yanzay/tbot"
	"gorm.io/gorm"
	"telegram-bot-reminder/database/repositories"
	"telegram-bot-reminder/telegram/messages"
	"time"
)

func Handle(ctx context.Context, client *tbot.Client, server *tbot.Server) {
	Start(ctx, client, server)
	go messages.SendMessage(ctx, client)

	//send today's daily text
	server.HandleMessage("📗 Стих на cегодня", func(m *tbot.Message) {
		db := ctx.Value("db").(*gorm.DB)
		dtRepo := repositories.NewDailyText(db)
		t := time.Now()
		date := t.Format("2006/01/02")
		dt := dtRepo.FindByDate(date)

		text := fmt.Sprintf("🗓️ <b>%s</b>\n\n<i>%s</i> \n\n%s ", dt.Title, dt.Script, dt.Text)
		_, _ = client.SendMessage(m.Chat.ID, text, tbot.OptParseModeHTML)
	})

	//send tomorrow's daily text
	server.HandleMessage("📘 Стих на завтра", func(m *tbot.Message) {
		db := ctx.Value("db").(*gorm.DB)
		dtRepo := repositories.NewDailyText(db)
		t := time.Now().Add(24 * time.Hour)
		date := t.Format("2006/01/02")
		dt := dtRepo.FindByDate(date)

		text := fmt.Sprintf("🗓️ <b>%s</b>\n\n<i>%s</i> \n\n%s ", dt.Title, dt.Script, dt.Text)
		_, _ = client.SendMessage(m.Chat.ID, text, tbot.OptParseModeHTML)
	})

}
