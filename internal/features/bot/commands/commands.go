package commands

import (
	bot2 "bot-test/internal/features/bot"
	"context"
	"fmt"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

var (
	deleteGoodPrefix = "good:delete"
	deleteChatPrefix = "chat:delete"
)

type Commands struct {
	service bot2.IService
}

func (c *Commands) GetGoods(ctx context.Context, b *bot.Bot, update *models.Update) error {
	chatID := update.Message.Chat.ID

	goods, err := c.service.GetGoodsByOwnerId(ctx, chatID)
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

func (c *Commands) GetChats(ctx context.Context, b *bot.Bot, update *models.Update) error {
	chatID := update.Message.Chat.ID

	chats, err := c.service.GetChatsByOwnerId(ctx, chatID) // TODO: подумать тут и в миграции насчет имени чата
	if err != nil {
		return err
	}

	result := "Твои чаты:\n"

	for i, ch := range chats {
		result += fmt.Sprintf("%d", ch.Id)
		if i != len(chats)-1 {
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

	goods, err := c.service.GetGoodsByOwnerId(ctx, chatID)
	if err != nil {
		return err
	}

	keyboard := make([][]models.InlineKeyboardButton, 0)

	row := make([]models.InlineKeyboardButton, 0)

	for i, g := range goods {
		row = append(row, models.InlineKeyboardButton{
			Text:         g.Name,
			CallbackData: fmt.Sprintf("%s:%s", deleteGoodPrefix, g.Id), // TODO: Вынести ключ в константы
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

func (c *Commands) DeleteChat(ctx context.Context, b *bot.Bot, update *models.Update) error {
	chatID := update.Message.Chat.ID

	chats, err := c.service.GetChatsByOwnerId(ctx, chatID)
	if err != nil {
		return err
	}

	keyboard := make([][]models.InlineKeyboardButton, 0)

	row := make([]models.InlineKeyboardButton, 0)

	for i, ch := range chats {
		row = append(row, models.InlineKeyboardButton{
			Text:         fmt.Sprintf("%d", ch.Id),
			CallbackData: fmt.Sprintf("%s:%d", deleteChatPrefix, ch.Id), // TODO: Вынести ключ в константы
		})
		if i%2 == 1 || i == len(chats)-1 {
			keyboard = append(keyboard, row)
			row = make([]models.InlineKeyboardButton, 0)
		}
	}

	_, err = b.SendMessage(ctx, &bot.SendMessageParams{
		Text:        "Выбери чат для удаления",
		ChatID:      chatID,
		ReplyMarkup: models.InlineKeyboardMarkup{InlineKeyboard: keyboard},
	})
	return err
}

func NewCommands(service bot2.IService) bot2.ICommands {
	return &Commands{service: service}
}
