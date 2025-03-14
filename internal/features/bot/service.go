package bot

import "context"

type IService interface {
	SaveGood(ctx context.Context, ownerId int64, name string) error
	DeleteGood(ctx context.Context, goodId string) error
}
