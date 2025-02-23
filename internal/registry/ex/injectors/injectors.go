package injectors

import (
	"bot-test/internal/bot/cache"
	"bot-test/internal/bot/respository"
	"bot-test/internal/bot/service"
	"bot-test/internal/registry/ex"
	"context"
	"fmt"
	"github.com/lalkalol1907/tg-bot-stepper/stepper"
	"github.com/redis/go-redis/v9"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.uber.org/zap"
)

func WithLogger[T any](c *ex.Components[T]) (ex.CloseFunc, error) {
	c.Logger = otelzap.New(zap.NewExample())
	return nil, nil
}

func WithBotStepper[T any](c *ex.Components[T]) (ex.CloseFunc, error) {
	c.BotStepper = stepper.NewStepper(c.BotCache, c.Logger)
	return nil, nil
}

func WithRedis[T any](c *ex.Components[T]) (ex.CloseFunc, error) {
	c.Redis = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", c.Config.Redis.Host, c.Config.Redis.Port),
		Password: "",
		DB:       0,
	})

	return c.Redis.Close, c.Redis.Ping(context.Background()).Err()
}

func WithBotCache[T any](c *ex.Components[T]) (ex.CloseFunc, error) {
	c.BotCache = cache.NewBotCache(c.Redis)
	return nil, nil
}

func WithRepository[T any](c *ex.Components[T]) (ex.CloseFunc, error) {
	c.Repository = respository.NewRepository(c.Redis)
	return nil, nil
}

func WithService[T any](c *ex.Components[T]) (ex.CloseFunc, error) {
	c.Service = service.NewService(c.Repository)
	return nil, nil
}
