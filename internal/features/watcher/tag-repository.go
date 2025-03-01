package watcher

import (
	"bot-test/pkg/models"
	"context"
)

type ITagRepository interface {
	GetTagsByOwnerId(ctx context.Context, ownerId int64) ([]*models.Tag, error)
}
