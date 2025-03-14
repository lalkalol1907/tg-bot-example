package commands

import (
	bot2 "bot-test/internal/features/bot"
	"context"
	"fmt"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"strings"
)

type Commands struct {
	repository bot2.IRepository
}

func (c *Commands) GetGoods(ctx context.Context, b *bot.Bot, update *models.Update) error {
	chatID := update.Message.Chat.ID

	goods, err := c.repository.GetGoods(ctx, chatID)
	if err != nil {
		return err
	}

	result := "Твои товары:\n"

	for i, g := range goods {
		result += g.Name
		if i != len(goods)-1 {
			result += "\n"
		}
	}

	_, err = b.SendMessage(ctx, &bot.SendMessageParams{
		Text:   result,
		ChatID: chatID,
	})
	return err
}

func (c *Commands) DeleteGood(ctx context.Context, b *bot.Bot, update *models.Update) error {
	chatID := update.Message.Chat.ID

	goods, err := c.repository.GetGoods(ctx, chatID)
	if err != nil {
		return err
	}

	keyboard := make([][]models.InlineKeyboardButton, 0)

	row := make([]models.InlineKeyboardButton, 0)

	for i, g := range goods {
		row = append(row, models.InlineKeyboardButton{
			Text:         g.Name,
			CallbackData: fmt.Sprintf("good:delete:%s", g.Id), // TODO: Вынести ключ в константы
		})
		if i%2 == 1 || i == len(goods)-1 {
			keyboard = append(keyboard, row)
			row = make([]models.InlineKeyboardButton, 0)
		}
	}

	_, err = b.SendMessage(ctx, &bot.SendMessageParams{
		Text:        "Выбери товар для удаления",
		ChatID:      chatID,
		ReplyMarkup: models.InlineKeyboardMarkup{InlineKeyboard: keyboard},
	})
	return err
}

func (c *Commands) CallbackHandler(ctx context.Context, b *bot.Bot, update *models.CallbackQuery) error {
	command := update.Data

	fmt.Println("aboba")

	if strings.HasPrefix(command, "good:delete") {
		spl := strings.Split(command, ":")
		goodID := spl[len(spl)-1]
		err := c.repository.DeleteGood(ctx, goodID)

		text := "Товар удален"

		if err != nil {
			text = "Ошибка удаления товара"
		}

		_, err = b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
			CallbackQueryID: update.ID,
			Text:            text,
			ShowAlert:       true,
		})
		return err
	}

	return nil
}

func NewCommands(repository bot2.IRepository) bot2.ICommands {
	return &Commands{repository: repository}
}
