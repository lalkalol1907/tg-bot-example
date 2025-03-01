package bot

import "context"

type IRepository interface {
	SaveMessage(ctx context.Context, chatId int64, message string) error
	GetMessage(ctx context.Context, chatId int64) (string, error)
}
