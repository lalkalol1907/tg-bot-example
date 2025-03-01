package example

import (
	bot2 "bot-test/internal/features/bot"
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/lalkalol1907/tg-bot-stepper/stepper"
	"github.com/lalkalol1907/tg-bot-stepper/types"
)

type Feature struct {
	service bot2.IService
}

func (f *Feature) Start(ctx context.Context, b *bot.Bot, update *models.Update) (types.StepExecutionResult, error) {
	if _, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Введи сообщение",
	}); err != nil {
		return types.StepExecutionResult{}, err
	}

	nextStep := "registerMessage"
	return types.StepExecutionResult{NextStep: &nextStep, IsFinal: false}, nil
}

func (f *Feature) RegisterMessage(ctx context.Context, b *bot.Bot, update *models.Update) (types.StepExecutionResult, error) {
	message := update.Message.Text
	chatId := update.Message.Chat.ID

	err := f.service.SaveMessage(ctx, chatId, message)
	if err != nil {
		return types.StepExecutionResult{}, err
	}

	if _, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Я сохранил брат, напиши что-нибудт, чтобы узнать",
	}); err != nil {
		return types.StepExecutionResult{}, err
	}

	nextStep := "getMessage"
	return types.StepExecutionResult{NextStep: &nextStep, IsFinal: false}, nil
}

func (f *Feature) GetMessage(ctx context.Context, b *bot.Bot, update *models.Update) (types.StepExecutionResult, error) {
	chatId := update.Message.Chat.ID

	text, err := f.service.GetMessage(ctx, chatId)
	if err != nil {
		return types.StepExecutionResult{}, err
	}

	if _, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   text,
	}); err != nil {
		return types.StepExecutionResult{}, err
	}

	return types.StepExecutionResult{NextStep: nil, IsFinal: true}, nil
}

func NewExampleFeature(service bot2.IService) *stepper.Feature {
	feature := stepper.NewFeature()

	featureClass := &Feature{service: service}

	feature.AddStep("start", featureClass.Start)
	feature.AddStep("registerMessage", featureClass.RegisterMessage)
	feature.AddStep("getMessage", featureClass.GetMessage)

	return feature
}
