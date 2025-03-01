package transport

import (
	"bot-test/internal/features/watcher"
	"bot-test/internal/features/watcher/types"
	"context"
	"errors"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"github.com/zelenin/go-tdlib/client"
	"go.uber.org/zap"
)

type Transport struct {
	tdlib   *client.Client
	logger  *otelzap.Logger
	service watcher.IService

	listener *client.Listener
}

func (t *Transport) Start(ctx context.Context) {
	t.listener = t.tdlib.GetListener()

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
}

func (t *Transport) Stop() error {
	if t.listener == nil {
		return errors.New("no listener")
	}

	t.listener.Close()
	return nil
}

func NewTransport(tdlib *client.Client, logger *otelzap.Logger) watcher.ITransport {
	return &Transport{
		tdlib:  tdlib,
		logger: logger,
	}
}
