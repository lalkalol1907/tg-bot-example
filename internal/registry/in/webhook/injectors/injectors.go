package injectors

import (
	"bot-test/internal/bot/transport"
	"bot-test/internal/registry/ex"
	"bot-test/internal/registry/in/webhook"
	"github.com/go-telegram/bot"
)

func WithWebhookTransport(c *ex.Components[webhook.Components]) (ex.CloseFunc, error) {
	c.In.BotTransport = transport.NewWebhookTransport(c.BotStepper, c.Logger, c.Service, c.In.Bot, c.Http)
	return nil, nil
}

func WithBot(c *ex.Components[webhook.Components]) (ex.CloseFunc, error) {
	b, err := bot.New(c.Config.Bot.Token, []bot.Option{
		bot.WithDefaultHandler(c.BotStepper.Handle),
	}...)
	if err != nil {
		return nil, err
	}
	c.In.Bot = b
	return nil, nil
}
