package service

import (
	bot2 "bot-test/internal/features/bot"
	"context"
)

type Service struct {
	repository bot2.IRepository
}

func (s *Service) AddGood(ctx context.Context, ownerId int64, name string) error {
	return s.repository.AddGood(ctx, ownerId, name)
}

func (s *Service) DeleteGood(ctx context.Context, goodId string) error {
	return s.repository.DeleteGood(ctx, goodId)
}

func NewService(repository bot2.IRepository) bot2.IService {
	return &Service{
		repository: repository,
	}
}
