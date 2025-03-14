package commands

import (
	"context"
	"fmt"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"strconv"
	"strings"
)

func (c *Commands) CallbackHandler(ctx context.Context, b *bot.Bot, update *models.CallbackQuery) error {
	command := update.Data

	fmt.Println("aboba")

	if strings.HasPrefix(command, deleteGoodPrefix) {
		spl := strings.Split(command, ":")
		goodID := spl[len(spl)-1]
		err := c.service.DeleteGood(ctx, goodID)

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

	if strings.HasPrefix(command, deleteChatPrefix) {
		spl := strings.Split(command, ":")
		chatIdRaw := spl[len(spl)-1]

		chatId, err := strconv.ParseInt(chatIdRaw, 10, 64)
		if err == nil {
			err = c.service.DeleteChat(ctx, chatId)
		}

		text := "Чат удален"

		if err != nil {
			text = "Ошибка удаления чата"
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
