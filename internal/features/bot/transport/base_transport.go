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
		AddFeature("AddGood", "/add_good", features.NewAddGoodFeature(t.service)).
		AddFeature("AddChat", "/add_chat", features.NewAddChatFeature(t.service)).
		AddSingleStepCommand("/delete_good", t.commands.DeleteGood).
		AddSingleStepCommand("/get_goods", t.commands.GetGoods).
		AddSingleStepCommand("/delete_chat", t.commands.DeleteChat).
		AddSingleStepCommand("/get_chats", t.commands.GetChats).
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
