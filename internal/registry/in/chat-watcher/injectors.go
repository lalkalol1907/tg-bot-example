package chat_watcher

import (
	"bot-test/internal/features/watcher/service"
	"bot-test/internal/features/watcher/transport"
	"bot-test/internal/registry/ex"
	lock_service "bot-test/pkg/lock-service"
	"github.com/zelenin/go-tdlib/client"
	"path/filepath"
)

func WithTdlib(c *ex.Components[Components]) (ex.CloseFunc, error) {
	tdlibParameters := &client.SetTdlibParametersRequest{
		UseTestDc:           false,
		DatabaseDirectory:   filepath.Join(".tdlib", "database"),
		FilesDirectory:      filepath.Join(".tdlib", "files"),
		UseFileDatabase:     false,
		UseChatInfoDatabase: false,
		UseMessageDatabase:  false,
		UseSecretChats:      false,
		ApiId:               c.Config.TG.ApiId,
		ApiHash:             c.Config.TG.ApiHash,
		SystemLanguageCode:  "en",
		DeviceModel:         "Server",
		SystemVersion:       "1.0.0",
		ApplicationVersion:  "1.0.0",
	}
	// client authorizer
	authorizer := client.ClientAuthorizer(tdlibParameters)
	go client.CliInteractor(authorizer)

	tdlibClient, err := client.NewClient(
		authorizer,
		client.WithLogVerbosity(&client.SetLogVerbosityLevelRequest{
			NewVerbosityLevel: 0,
		}),
	)

	if err != nil {
		return nil, err
	}

	c.In.Tdlib = tdlibClient

	return func() error {
		_, err := tdlibClient.Close()
		return err
	}, nil
}

func WithTransport(c *ex.Components[Components]) (ex.CloseFunc, error) {
	c.In.Transport = transport.NewTransport(c.In.Tdlib, c.Logger)
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
