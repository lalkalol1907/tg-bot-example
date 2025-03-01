package cache

import (
	"context"
	"errors"
	"fmt"
	"github.com/lalkalol1907/tg-bot-stepper/types"
	"github.com/redis/go-redis/v9"
	"strings"
)

type BotCache struct {
	redisClient *redis.Client
	prefix      string
}

func (c *BotCache) formatKey(values ...string) string {
	if c.prefix != "" {
		values = append([]string{c.prefix}, values...)
	}

	return strings.Join(values, ":")
}

func (c *BotCache) Set(ctx context.Context, chatId int64, featureName string, stepName string) error {
	key := c.formatKey(fmt.Sprintf("%d", chatId))

	value := fmt.Sprintf("%s:%s", featureName, stepName)

	if err := c.redisClient.Set(ctx, key, value, 0).Err(); err != nil {
		return fmt.Errorf("error setting key %s: %v", key, err)
	}

	return nil
}

func (c *BotCache) Get(ctx context.Context, chatId int64) (*string, *string, error) {
	key := c.formatKey(fmt.Sprintf("%d", chatId))

	val, err := c.redisClient.Get(ctx, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, nil, nil
		}
		return nil, nil, fmt.Errorf("error getting key %s: %v", key, err)
	}

	split := strings.Split(val, ":")

	return &split[0], &split[1], err
}

func (c *BotCache) Del(ctx context.Context, chatId int64) error {
	key := c.formatKey(fmt.Sprintf("%d", chatId))

	if err := c.redisClient.Del(ctx, key).Err(); err != nil {
		return fmt.Errorf("error deleting key %s: %v", key, err)
	}

	return nil
}

func NewBotCache(redisClient *redis.Client) types.Cache {
	return &BotCache{
		redisClient: redisClient,
	}
}
