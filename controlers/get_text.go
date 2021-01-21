package controlers

import (
	"context"
	"os"
	"telegram-bot-reminder/parsers"
	"time"
)

func GetDailyText(ctx context.Context) {
	for {
		t := time.Now().Add(24 * time.Hour)
		frt := t.Format("2006/01/02")
		url := os.Getenv("JW_URL") + frt
		parsers.ParserDailyVerse(ctx, url, frt)

		duration := time.Hour * 24
		time.Sleep(duration)

	}
}
