package polling

import (
	"bot-test/internal/features/bot/transport"
	"bot-test/internal/registry/ex"
	"github.com/go-telegram/bot"
)

func WithBot(c *ex.Components[Components]) (ex.CloseFunc, error) {
	b, err := bot.New(
		c.Config.Bot.Token,
		bot.WithDefaultHandler(c.BotStepper.Handle),
	)
	if err != nil {
		return nil, err
	}
	c.In.Bot = b
	return nil, nil
}

func WithPollTransport(c *ex.Components[Components]) (ex.CloseFunc, error) {
	c.In.BotTransport = transport.NewPollTransport(c.BotStepper, c.Logger, c.Service, c.In.Bot)
	return c.In.BotTransport.Stop, nil
}
