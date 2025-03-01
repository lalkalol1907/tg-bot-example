package watcher

import (
	"bot-test/pkg/models"
	"context"
)

type IWorkerRepository interface {
	GetWorkers(ctx context.Context) ([]*models.Worker, error)
}
