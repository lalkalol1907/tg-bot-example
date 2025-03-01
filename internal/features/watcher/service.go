package watcher

import (
	"bot-test/internal/features/watcher/types"
	"context"
)

type IService interface {
	ProcessMessage(ctx context.Context, update *types.IncomingMessage) error
}
