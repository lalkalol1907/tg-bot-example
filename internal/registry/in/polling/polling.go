package polling

import (
	"bot-test/internal/registry/ex"
)

func Polling() {
	d := ex.NewDiContainer[Components]()

	c := d.Provide(
		WithBot,
		WithPollTransport,
	)

	d.Run(c.In.BotTransport.Run)
}
