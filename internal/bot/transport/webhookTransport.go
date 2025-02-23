package transport

import (
	bot2 "bot-test/internal/bot"
	"context"
	"github.com/go-telegram/bot"
	"github.com/labstack/echo/v4"
	"github.com/lalkalol1907/tg-bot-stepper/stepper"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
)

type WebhookTransport struct {
	BaseTransport

	e *echo.Echo
}

func (t *WebhookTransport) Run(ctx context.Context) {
	t.registerFeatures()
	t.bot.Start(ctx) // TODO:
}

func NewWebhookTransport(
	stepper *stepper.Stepper,
	logger *otelzap.Logger,
	service bot2.Service,
	bot *bot.Bot,
	e *echo.Echo,
) bot2.WebhookTransport {
	return &WebhookTransport{
		BaseTransport: *NewTransport(stepper, logger, service, bot),
		e:             e,
	}
}
