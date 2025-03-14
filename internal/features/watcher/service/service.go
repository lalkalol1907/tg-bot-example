package service

import (
	"bot-test/config"
	"bot-test/internal/features/watcher"
	"bot-test/internal/features/watcher/types"
	lockservice "bot-test/pkg/lock-service"
	"bot-test/pkg/models"
	tagcheck "bot-test/pkg/tag-check"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/segmentio/kafka-go"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"github.com/zelenin/go-tdlib/client"
	"go.uber.org/zap"
)

type Service struct {
	lockService *lockservice.LockService
	logger      *otelzap.Logger
	config      *config.Config

	chatRepository   watcher.IChatRepository
	tagRepository    watcher.ITagRepository
	workerRepository watcher.IWorkerRepository

	tdlib *client.Client

	kafkaConnection *kafka.Conn
}

func (s *Service) ProcessMessage(ctx context.Context, update *types.IncomingMessage) error {
	// проверяем, можем ли мы обрабатывать сообщение из этого чата
	chat, err := s.chatRepository.GetChatById(ctx, update.ChatId, update.WorkerId)
	if err != nil {
		s.logger.Ctx(ctx).Warn("error getting chat", zap.Error(err))
		return err
	}

	if chat == nil {
		s.logger.Ctx(ctx).Warn("no permissions to this chat", zap.Int64("chatId", update.ChatId))
		return errors.New("worker has no permissions for this chat")
	}

	lockKey := fmt.Sprintf("%d:%d", update.ChatId, update.MessageId)

	if err := s.lockService.TryLockOnce(ctx, lockKey); err != nil {
		s.logger.Warn("error setting lock", zap.Error(err))
		return err
	}

	defer s.lockService.RemoveLock(ctx, lockKey)

	// Достаем из базы теги для этого ownerId
	tags, err := s.tagRepository.GetTagsByOwnerId(ctx, chat.OwnerId)
	if err != nil {
		s.logger.Ctx(ctx).Warn("error getting tags", zap.Error(err))
		return err
	}

	// Прогоняем на совпадения
	res, err := tagcheck.FindTags(update.Text, tags)
	if err != nil {
		s.logger.Ctx(ctx).Warn("error finding tags", zap.Error(err))
		return err
	}

	val := &types.Collision{
		Result:  res,
		UserId:  update.UserId,
		ChatId:  update.ChatId,
		OwnerId: chat.OwnerId,
	}

	kfkMsg, err := json.Marshal(val)
	if err != nil {
		s.logger.Ctx(ctx).Warn("error marshalling message to kafka", zap.Error(err))
		return err
	}

	if _, err := s.kafkaConnection.WriteMessages(kafka.Message{
		Topic: s.config.Kafka.ProducerTopics.NewMessage,
		Key:   []byte(fmt.Sprintf("%d", update.MessageId)),
		Value: kfkMsg,
	}); err != nil {
		s.logger.Ctx(ctx).Warn("error writing to kafka", zap.Error(err))
		return err
	}

	_, err = s.tdlib.SendMessage(&client.SendMessageRequest{
		ChatId: update.ChatId, // TODO: Тут чат с пользаком начать
		InputMessageContent: &client.InputMessageText{
			Text: &client.FormattedText{
				Text: "Мы тут тебе айфончик завезли", // TODO: Текст получать из БД
			},
		},
	})

	return err
}

func (s *Service) GetWorkers(ctx context.Context) ([]*models.Worker, error) {
	return s.workerRepository.GetWorkers(ctx)
}

func NewService(
	lockService *lockservice.LockService,
	logger *otelzap.Logger,
) watcher.IService {
	return &Service{
		lockService: lockService,
		logger:      logger,
	}
}
