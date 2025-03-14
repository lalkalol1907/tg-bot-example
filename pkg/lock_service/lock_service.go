package lock_service

import (
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"time"
)

type LockService struct {
	redis *redis.Client
}

func (s *LockService) TryLockOnce(ctx context.Context, key string) error {
	return s.redis.SetNX(ctx, key, 1, time.Minute).Err()
}

func (s *LockService) SetLock(ctx context.Context, key string) error {
	for i := 0; i < 30; i++ {
		if err := s.TryLockOnce(ctx, key); err != nil {
			time.Sleep(time.Second)
			continue
		}
		return nil
	}

	return errors.New("error setting lock")
}

func (s *LockService) RemoveLock(ctx context.Context, key string) error {
	return s.redis.Del(ctx, key).Err()
}

func NewLockService(redis *redis.Client) *LockService {
	return &LockService{
		redis: redis,
	}
}
