package features

import (
	bot2 "bot-test/internal/features/bot"
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/lalkalol1907/tg-bot-stepper/stepper"
	"github.com/lalkalol1907/tg-bot-stepper/types"
)

type AddGoodFeature struct {
	service bot2.IService
}

func (f *AddGoodFeature) Start(ctx context.Context, b *bot.Bot, update *models.Update) (types.StepExecutionResult, error) {
	if _, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Введи название товара", // TODO: возвращать клавиатуру с товарами
	}); err != nil {
		return types.StepExecutionResult{}, err
	}

	nextStep := "add"
	return types.StepExecutionResult{NextStep: &nextStep, IsFinal: false}, nil
}

func (f *AddGoodFeature) AddGood(ctx context.Context, b *bot.Bot, update *models.Update) (types.StepExecutionResult, error) {
	name := update.Message.Text
	chatId := update.Message.Chat.ID

	err := f.service.SaveGood(ctx, chatId, name)
	if err != nil {
		return types.StepExecutionResult{}, err
	}

	if _, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Товар добавлен",
	}); err != nil {
		return types.StepExecutionResult{}, err
	}

	return types.StepExecutionResult{NextStep: nil, IsFinal: true}, nil
}

func NewAddGoodFeature(service bot2.IService) *stepper.Feature {
	feature := stepper.NewFeature()

	featureClass := &AddGoodFeature{service: service}

	feature.AddStep("start", featureClass.Start)
	feature.AddStep("add", featureClass.AddGood)

	return feature
}
