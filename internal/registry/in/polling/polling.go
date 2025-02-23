package polling

import (
	"bot-test/internal/registry/ex"
	"bot-test/internal/registry/in/polling/injectors"
)

func Polling() {
	d := ex.NewDiContainer[Components]()

	c := d.Provide(
		injectors.WithBot,
		injectors.WithPollTransport,
	)

	d.Run(c.In.BotTransport.Run)
}
