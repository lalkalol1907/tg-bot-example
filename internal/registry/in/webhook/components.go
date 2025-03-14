package webhook

import (
	"bot-test/internal/features/bot"
	bot2 "github.com/go-telegram/bot"
)

type Components struct {
	Bot      *bot2.Bot
	Commands bot.ICommands

	Repository bot.IRepository
	Service    bot.IService

	BotTransport bot.IBaseTransport
}
