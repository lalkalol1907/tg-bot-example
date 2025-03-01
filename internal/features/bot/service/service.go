package service

import (
	bot2 "bot-test/internal/features/bot"
	"context"
)

type Service struct {
	repository bot2.IRepository
}

func (s *Service) SaveMessage(ctx context.Context, chatId int64, message string) error {
	return s.repository.SaveMessage(ctx, chatId, message)
}

func (s *Service) GetMessage(ctx context.Context, chatId int64) (string, error) {
	return s.repository.GetMessage(ctx, chatId)
}

func NewService(repository bot2.IRepository) bot2.IService {
	return &Service{
		repository: repository,
	}
}
