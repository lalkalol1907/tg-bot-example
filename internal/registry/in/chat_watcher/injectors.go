package chat_watcher

import (
	"bot-test/internal/features/watcher/service"
	"bot-test/internal/features/watcher/transport"
	"bot-test/internal/registry/ex"
	lock_service "bot-test/pkg/lock-service"
)

func WithTransport(c *ex.Components[Components]) (ex.CloseFunc, error) {
	c.In.Transport = transport.NewTransport(c.Config, c.Logger)
	return c.In.Transport.Stop, nil
}

func WithService(c *ex.Components[Components]) (ex.CloseFunc, error) {
	c.In.Service = service.NewService(c.In.LockService, c.Logger)
	return nil, nil
}

func WithLockService(c *ex.Components[Components]) (ex.CloseFunc, error) {
	c.In.LockService = lock_service.NewLockService(c.Redis)
	return nil, nil
}
