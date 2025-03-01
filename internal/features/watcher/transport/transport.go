package transport

import (
	"bot-test/config"
	"bot-test/internal/features/watcher"
	"bot-test/internal/features/watcher/types"
	auth_service "bot-test/pkg/auth-service"
	"bot-test/pkg/models"
	"context"
	"errors"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"github.com/zelenin/go-tdlib/client"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	"path/filepath"
)

type Transport struct {
	logger  *otelzap.Logger
	service watcher.IService
	config  *config.Config

	authService auth_service.IAuthService

	listener *client.Listener
}

func (t *Transport) initWatcher(ctx context.Context, worker *models.Worker) error {
	tdlibParameters := &client.SetTdlibParametersRequest{
		UseTestDc:           false,
		DatabaseDirectory:   filepath.Join(".tdlib", "database"),
		FilesDirectory:      filepath.Join(".tdlib", "files"),
		UseFileDatabase:     false,
		UseChatInfoDatabase: false,
		UseMessageDatabase:  false,
		UseSecretChats:      false,
		ApiId:               t.config.TG.ApiId,
		ApiHash:             t.config.TG.ApiHash,
		SystemLanguageCode:  "en",
		DeviceModel:         "Server",
		SystemVersion:       "1.0.0",
		ApplicationVersion:  "1.0.0",
	}
	// TODO: Авторизация через бота (не забыть мютекс по овнеру)
	authorizer := client.ClientAuthorizer(tdlibParameters)
	go t.authService.Authorize(ctx, authorizer, worker)
	go client.CliInteractor(authorizer)

	tdlibClient, err := client.NewClient(
		authorizer,
		client.WithLogVerbosity(&client.SetLogVerbosityLevelRequest{
			NewVerbosityLevel: 0,
		}),
	)
	if err != nil {
		return err
	}

	listener := tdlibClient.GetListener()
	defer listener.Close()

	for update := range t.listener.Updates {
		if update.GetType() == client.TypeUpdateNewMessage {
			message := update.(*client.UpdateNewMessage)

			if message.Message.Content.MessageContentType() == client.TypeMessageText {
				content := message.Message.Content.(*client.MessageText)
				t.logger.Ctx(ctx).Info(
					"Recieved message from chat",
					zap.Int64("chatId", message.Message.ChatId),
				)

				sender, ok := message.Message.SenderId.(*client.MessageSenderUser)
				if !ok {
					t.logger.Ctx(ctx).Warn(
						"Error getting sender",
						zap.Int64("messageId", message.Message.Id),
					)
				}

				// TODO: Прокидывать фотку?
				err := t.service.ProcessMessage(ctx, &types.IncomingMessage{
					ChatId:    message.Message.ChatId,
					Text:      content.Text.Text,
					UserId:    sender.UserId,
					MessageId: message.Message.Id,
				})

				if err != nil {
					t.logger.Ctx(ctx).Warn(
						"Error processing message",
						zap.Int64("messageId", message.Message.Id),
						zap.Int64("chatId", message.Message.ChatId),
					)
				}
			}

		}
	}
	return nil
}

func (t *Transport) Start(ctx context.Context) {
	workers, err := t.service.GetWorkers(ctx)
	if err != nil {
		// TODO: runner with err
	}

	g, gctx := errgroup.WithContext(ctx)

	for _, w := range workers {
		g.Go(func() error {
			return t.initWatcher(gctx, w)
		})
	}

	err = g.Wait()
	if err != nil {
		// TODO: runner with err
	}
}

func (t *Transport) Stop() error {
	if t.listener == nil {
		return errors.New("no listener")
	}

	t.listener.Close()
	return nil
}

func NewTransport(config *config.Config, logger *otelzap.Logger) watcher.ITransport {
	return &Transport{
		config: config,
		logger: logger,
	}
}
