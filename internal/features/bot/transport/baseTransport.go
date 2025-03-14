package transport

import (
	bot2 "bot-test/internal/features/bot"
	"bot-test/internal/features/bot/features"
	"github.com/go-telegram/bot"
	"github.com/lalkalol1907/tg-bot-stepper/stepper"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
)

type BaseTransport struct {
	stepper *stepper.Stepper
	logger  *otelzap.Logger
	service bot2.IService

	commands bot2.ICommands

	bot *bot.Bot
}

func (t *BaseTransport) registerFeatures() {
	t.stepper.
		AddFeature("AddGood", "/add-good", features.NewAddGoodFeature(t.service)).
		AddSingleStepCommand("/delete-good", t.commands.DeleteGood).
		AddSingleStepCommand("/get-goods", t.commands.GetGoods).
		AddCallbackHandler(t.commands.CallbackHandler)
}

func NewTransport(
	stepper *stepper.Stepper,
	logger *otelzap.Logger,
	service bot2.IService,
	bot *bot.Bot,
	commands bot2.ICommands,
) *BaseTransport {
	return &BaseTransport{
		stepper:  stepper,
		logger:   logger,
		service:  service,
		bot:      bot,
		commands: commands,
	}
}
