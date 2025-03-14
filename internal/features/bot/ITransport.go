package bot

import "context"

type IPollTransport interface {
	IBaseTransport
	Stop() error
}

type IBaseTransport interface {
	Run(ctx context.Context)
}
