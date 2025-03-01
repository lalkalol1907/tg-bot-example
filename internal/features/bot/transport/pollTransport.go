package transport

import (
	bot3 "bot-test/internal/features/bot"
	"context"
	"github.com/go-telegram/bot"
	"github.com/lalkalol1907/tg-bot-stepper/stepper"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
)

type PollTransport struct {
	BaseTransport
}

func (t *PollTransport) Run(ctx context.Context) {
	t.registerFeatures()
	t.bot.Start(ctx)
}

func (t *PollTransport) Stop() error {
	_, err := t.bot.Close(context.Background())

	return err
}

func NewPollTransport(
	stepper *stepper.Stepper,
	logger *otelzap.Logger,
	service bot3.IService,
	bot *bot.Bot,
) bot3.IPollTransport {
	return &PollTransport{
		*NewTransport(stepper, logger, service, bot),
	}
}
