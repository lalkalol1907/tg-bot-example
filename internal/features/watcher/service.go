package watcher

import (
	"bot-test/internal/features/watcher/types"
	"bot-test/pkg/models"
	"context"
)

type IService interface {
	ProcessMessage(ctx context.Context, update *types.IncomingMessage) error

	GetWorkers(ctx context.Context) ([]*models.Worker, error)
}
