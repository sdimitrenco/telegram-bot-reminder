package telegram

import (
	"context"
	"fmt"
	"github.com/yanzay/tbot"
)

func Handle(ctx context.Context, client *tbot.Client, server *tbot.Server) {
	Start(ctx, client, server)
	server.HandleMessage("слово", func(m *tbot.Message) {
		message := fmt.Sprintf("%s\n", "обратный текст")
		_, _ = client.SendMessage(m.Chat.ID, message, tbot.OptParseModeMarkdown)
	})
}
