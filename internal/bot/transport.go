package bot

import "context"

type Transport interface {
	Run(ctx context.Context)
	Stop() error
}
