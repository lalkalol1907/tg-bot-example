package bot

import "context"

type Service interface {
	SaveMessage(ctx context.Context, chatId int64, message string) error
	GetMessage(ctx context.Context, chatId int64) (string, error)
}
