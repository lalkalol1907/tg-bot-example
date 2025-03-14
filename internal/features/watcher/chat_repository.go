package watcher

import (
	"bot-test/pkg/models"
	"context"
)

type IChatRepository interface {
	GetChatById(ctx context.Context, chatId int64, workerId int64) (*models.Chat, error)
}
