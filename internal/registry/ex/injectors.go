package ex

import (
	"bot-test/internal/features/bot/cache"
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/lalkalol1907/tg-bot-stepper/stepper"
	"github.com/redis/go-redis/v9"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.uber.org/zap"

	_ "github.com/lib/pq"
)

func WithLogger[T any](c *Components[T]) (CloseFunc, error) {
	c.Logger = otelzap.New(zap.NewExample())
	return nil, nil
}

func WithBotStepper[T any](c *Components[T]) (CloseFunc, error) {
	c.BotStepper = stepper.NewStepper(c.BotCache, c.Logger)
	return nil, nil
}

func WithRedis[T any](c *Components[T]) (CloseFunc, error) {
	c.Redis = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", c.Config.Redis.Host, c.Config.Redis.Port),
		Password: "",
		DB:       0,
	})

	return c.Redis.Close, c.Redis.Ping(context.Background()).Err()
}

func WithBotCache[T any](c *Components[T]) (CloseFunc, error) {
	c.BotCache = cache.NewBotCache(c.Redis, c.Config.Redis.CachePrefix)
	return nil, nil
}

func WithDB[T any](c *Components[T]) (CloseFunc, error) {
	db, err := sqlx.Connect("postgres",
		fmt.Sprintf(
			"user=%s dbname=%s password=%s host=%s port=%s sslmode=%s",
			c.Config.DB.Username,
			c.Config.DB.DbName,
			c.Config.DB.Password,
			c.Config.DB.Host,
			c.Config.DB.Port,
			c.Config.DB.SslMode,
		),
	)

	if err != nil {
		return nil, err
	}

	c.DB = db
	return db.Close, nil
}
