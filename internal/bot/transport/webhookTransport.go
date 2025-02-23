package transport

import (
	bot2 "bot-test/internal/bot"
	"context"
	"github.com/go-telegram/bot"
	"github.com/lalkalol1907/tg-bot-stepper/stepper"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
)

type WebhookTransport struct {
	Transport
}

func (t *WebhookTransport) Run(ctx context.Context) {
	t.registerFeatures()
	t.bot.Start(ctx) // TODO:
}

func (t *WebhookTransport) Stop() error {
	return nil // TODO:
}

func NewWebhookTransport(
	stepper *stepper.Stepper,
	logger *otelzap.Logger,
	service bot2.Service,
	bot *bot.Bot,
) bot2.Transport {
	return &WebhookTransport{
		*NewTransport(stepper, logger, service, bot),
	}
}
