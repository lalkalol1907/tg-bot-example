package chat_watcher

import "bot-test/internal/registry/ex"

func Watcher() {
	d := ex.NewDiContainer[Components]()

	c := d.Provide(
		WithTdlib,
		WithLockService,
		WithTransport,
		WithService,
	)

	d.Run(c.In.Transport.Start)
}
