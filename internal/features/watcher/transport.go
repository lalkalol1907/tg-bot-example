package watcher

import "context"

type ITransport interface {
	Start(ctx context.Context)
	Stop() error
}
