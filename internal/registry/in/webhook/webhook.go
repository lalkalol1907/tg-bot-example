package webhook

import (
	"bot-test/internal/registry/ex"
)

func Webhook() {
	d := ex.NewDiContainer[Components]()

	c := d.Provide(
		WithRepository,
		WithService,
		WithBot,
		WithCommands,
		WithWebhookTransport,
	)

	d.Run(c.In.BotTransport.Run)
}
