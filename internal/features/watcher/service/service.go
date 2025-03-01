package service

import (
	"bot-test/internal/features/watcher"
	"bot-test/internal/features/watcher/types"
	lockservice "bot-test/pkg/lock-service"
	"bot-test/pkg/models"
	tagcheck "bot-test/pkg/tag-check"
	"context"
	"errors"
	"fmt"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.uber.org/zap"
)

type Service struct {
	lockService *lockservice.LockService
	logger      *otelzap.Logger

	chatRepository   watcher.IChatRepository
	tagRepository    watcher.ITagRepository
	workerRepository watcher.IWorkerRepository
}

func (s *Service) ProcessMessage(ctx context.Context, update *types.IncomingMessage) error {
	// проверяем, можем ли мы обрабатывать сообщение из этого чата
	chat, err := s.chatRepository.GetChatByChatIdWorkerId(ctx, update.ChatId, update.WorkerId)
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
	_, err = tagcheck.FindTags(update.Text, tags)
	if err != nil {
		s.logger.Ctx(ctx).Warn("error finding tags", zap.Error(err))
		return err
	}

	//
	// Пушим в кафку, если совпало
	// Отправляем пользаку, если совпало (тоже пушим в кафку)

	return nil
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
