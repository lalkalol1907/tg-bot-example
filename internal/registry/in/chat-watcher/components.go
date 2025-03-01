package chat_watcher

import (
	"bot-test/internal/features/watcher"
	lock_service "bot-test/pkg/lock-service"
	"github.com/zelenin/go-tdlib/client"
)

type Components struct {
	Tdlib *client.Client

	LockService *lock_service.LockService

	Transport watcher.ITransport
	Service   watcher.IService
}
