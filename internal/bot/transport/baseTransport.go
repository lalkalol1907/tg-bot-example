package transport

import (
	bot2 "bot-test/internal/bot"
	"bot-test/internal/bot/features/example"
	"github.com/go-telegram/bot"
	"github.com/lalkalol1907/tg-bot-stepper/stepper"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
)

type BaseTransport struct {
	stepper *stepper.Stepper
	logger  *otelzap.Logger
	service bot2.Service

	bot *bot.Bot
}

func (t *BaseTransport) registerFeatures() {
	t.stepper.AddFeature("example", "/start", example.NewExampleFeature(t.service))
}

func NewTransport(
	stepper *stepper.Stepper,
	logger *otelzap.Logger,
	service bot2.Service,
	bot *bot.Bot,
) *BaseTransport {
	return &BaseTransport{
		stepper: stepper,
		logger:  logger,
		service: service,
		bot:     bot,
	}
}
