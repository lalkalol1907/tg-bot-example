package respository

import (
	"bot-test/internal/features/bot"
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

const prefix = "bot:message"

type Repository struct {
	redisClient *redis.Client
}

func (r *Repository) SaveMessage(ctx context.Context, chatId int64, message string) error {
	return r.redisClient.Set(ctx, fmt.Sprintf("%s:%d", prefix, chatId), message, 0).Err()
}

func (r *Repository) GetMessage(ctx context.Context, chatId int64) (string, error) {
	return r.redisClient.Get(ctx, fmt.Sprintf("%s:%d", prefix, chatId)).Result()
}

func NewRepository(redisClient *redis.Client) bot.IRepository {
	return &Repository{
		redisClient: redisClient,
	}
}
