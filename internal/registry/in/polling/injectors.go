package polling

import (
	"bot-test/internal/features/bot/commands"
	"bot-test/internal/features/bot/respository"
	"bot-test/internal/features/bot/service"
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

func WithRepository(c *ex.Components[Components]) (ex.CloseFunc, error) {
	c.In.Repository = respository.NewRepository(c.DB)
	return nil, nil
}

func WithService(c *ex.Components[Components]) (ex.CloseFunc, error) {
	c.In.Service = service.NewService(c.In.Repository)
	return nil, nil
}

func WithCommands(c *ex.Components[Components]) (ex.CloseFunc, error) {
	c.In.Commands = commands.NewCommands(c.In.Service)
	return nil, nil
}

func WithPollTransport(c *ex.Components[Components]) (ex.CloseFunc, error) {
	c.In.BotTransport = transport.NewPollTransport(c.BotStepper, c.Logger, c.In.Service, c.In.Bot, c.In.Commands)
	return c.In.BotTransport.Stop, nil
}
