package bot

import (
	"bot-test/pkg/models"
	"context"
)

type IRepository interface {
	SaveGood(ctx context.Context, ownerId int64, name string) error
	DeleteGood(ctx context.Context, goodId string) error
	GetGoods(ctx context.Context, ownerId int64) ([]*models.Good, error)
}
