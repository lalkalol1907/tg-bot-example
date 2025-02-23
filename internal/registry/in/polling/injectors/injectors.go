package injectors

import (
	"bot-test/internal/bot/transport"
	"bot-test/internal/registry/ex"
	"bot-test/internal/registry/in/polling"
	"github.com/go-telegram/bot"
)

func WithBot(c *ex.Components[polling.Components]) (ex.CloseFunc, error) {
	b, err := bot.New(c.Config.Bot.Token, []bot.Option{
		bot.WithDefaultHandler(c.BotStepper.Handle),
	}...)
	if err != nil {
		return nil, err
	}
	c.In.Bot = b
	return nil, nil
}

func WithPollTransport(c *ex.Components[polling.Components]) (ex.CloseFunc, error) {
	c.In.BotTransport = transport.NewPollTransport(c.BotStepper, c.Logger, c.Service, c.In.Bot)
	return c.In.BotTransport.Stop, nil
}
