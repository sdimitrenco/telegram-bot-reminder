package telegram

import (
	"github.com/yanzay/tbot"
)

var today = tbot.KeyboardButton{
	Text:            "📗 Стих на cегодня",
	RequestContact:  false,
	RequestLocation: false,
	RequestPoll:     nil,
}

var tomorrow = tbot.KeyboardButton{
	Text:            "📘 Стих на завтра",
	RequestContact:  false,
	RequestLocation: false,
	RequestPoll:     nil,
}

var profile = tbot.KeyboardButton{
	Text:            "🙋 ‍Завтра ведет",
	RequestContact:  false,
	RequestLocation: false,
	RequestPoll:     nil,
}

func MenuButtons() *tbot.ReplyKeyboardMarkup {
	return &tbot.ReplyKeyboardMarkup{
		Keyboard: [][]tbot.KeyboardButton{
			{today, tomorrow, profile},
		},
		ResizeKeyboard:  true,
		OneTimeKeyboard: false,
	}
}
