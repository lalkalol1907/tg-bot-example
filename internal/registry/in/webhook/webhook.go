package webhook

import (
	"bot-test/internal/registry/ex"
	"bot-test/internal/registry/in/webhook/injectors"
)

func Webhook() {
	d := ex.NewDiContainer[Components]()

	c := d.Provide(
		injectors.WithBot,
		injectors.WithWebhookTransport,
	)

	d.Run(c.In.BotTransport.Run)
}
