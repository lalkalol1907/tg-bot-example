package bot

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type ICommands interface {
	GetGoods(ctx context.Context, b *bot.Bot, update *models.Update) error
	DeleteGood(ctx context.Context, b *bot.Bot, update *models.Update) error

	GetChats(ctx context.Context, b *bot.Bot, update *models.Update) error
	DeleteChat(ctx context.Context, b *bot.Bot, update *models.Update) error

	CallbackHandler(ctx context.Context, b *bot.Bot, update *models.CallbackQuery) error
}
