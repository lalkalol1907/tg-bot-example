package bot

import (
	"bot-test/pkg/models"
	"context"
)

type IRepository interface {
	AddGood(ctx context.Context, ownerId int64, name string) error
	DeleteGood(ctx context.Context, goodId string) error
	GetGoodsByOwnerId(ctx context.Context, ownerId int64) ([]*models.Good, error)

	AddChat(ctx context.Context, ownerId int64, chatId int64) error
	DeleteChat(ctx context.Context, chatId int64) error
	GetChatsByOwnerId(ctx context.Context, ownerId int64) ([]*models.Chat, error)
}
