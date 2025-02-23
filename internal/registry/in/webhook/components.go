package webhook

import (
	"bot-test/internal/bot"
	bot2 "github.com/go-telegram/bot"
)

type Components struct {
	Bot *bot2.Bot

	BotTransport bot.Transport
}
