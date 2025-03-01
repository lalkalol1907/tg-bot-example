package webhook

import (
	"bot-test/internal/features/bot/transport"
	"bot-test/internal/registry/ex"
	"github.com/go-telegram/bot"
)

func WithWebhookTransport(c *ex.Components[Components]) (ex.CloseFunc, error) {
	c.In.BotTransport = transport.NewWebhookTransport(c.BotStepper, c.Logger, c.Service, c.In.Bot, c.Http)
	return nil, nil
}

func WithBot(c *ex.Components[Components]) (ex.CloseFunc, error) {
	b, err := bot.New(c.Config.Bot.Token, []bot.Option{
		bot.WithDefaultHandler(c.BotStepper.Handle),
	}...)
	if err != nil {
		return nil, err
	}
	c.In.Bot = b
	return nil, nil
}
