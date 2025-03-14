package transport

import (
	bot3 "bot-test/internal/features/bot"
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

	// TODO: Register Webhook
	t.e.GET("/bot", func(c echo.Context) error {
		t.bot.WebhookHandler().ServeHTTP(c.Response(), c.Request())
		return nil
	})

	t.bot.StartWebhook(ctx)
}

func NewWebhookTransport(
	stepper *stepper.Stepper,
	logger *otelzap.Logger,
	service bot3.IService,
	bot *bot.Bot,
	e *echo.Echo,
	commands bot3.ICommands,
) bot3.IBaseTransport {
	return &WebhookTransport{
		BaseTransport: *NewTransport(stepper, logger, service, bot, commands),
		e:             e,
	}
}
