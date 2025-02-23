package service

import (
	"bot-test/internal/bot"
	"context"
)

type Service struct {
	repository bot.Repository
}

func (s *Service) SaveMessage(ctx context.Context, chatId int64, message string) error {
	return s.repository.SaveMessage(ctx, chatId, message)
}

func (s *Service) GetMessage(ctx context.Context, chatId int64) (string, error) {
	return s.repository.GetMessage(ctx, chatId)
}

func NewService(repository bot.Repository) bot.Service {
	return &Service{
		repository: repository,
	}
}
