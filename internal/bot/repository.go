package bot

import "context"

type Repository interface {
	SaveMessage(ctx context.Context, chatId int64, message string) error
	GetMessage(ctx context.Context, chatId int64) (string, error)
}
