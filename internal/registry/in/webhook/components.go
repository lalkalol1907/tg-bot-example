package webhook

import (
	"bot-test/internal/features/bot"
	bot2 "github.com/go-telegram/bot"
)

type Components struct {
	Bot *bot2.Bot

	BotTransport bot.IBaseTransport
}
