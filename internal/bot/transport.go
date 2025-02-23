package bot

import "context"

type Transport interface {
	WebhookTransport
	Stop() error
}

type WebhookTransport interface {
	Run(ctx context.Context)
}
