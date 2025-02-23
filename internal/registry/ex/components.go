package ex

import (
	"bot-test/config"
	"bot-test/internal/bot"
	"bot-test/internal/registry/ex/injectors"
	"context"
	"github.com/labstack/echo/v4"
	"github.com/lalkalol1907/tg-bot-stepper/stepper"
	"github.com/lalkalol1907/tg-bot-stepper/types"
	"github.com/redis/go-redis/v9"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
)

type Components[T interface{}] struct {
	Config *config.Config
	Http   *echo.Echo // Сам инжектится для проб сразу после конфига. Можно использовать и для запросов

	Redis      *redis.Client
	Logger     *otelzap.Logger
	Repository bot.Repository
	Service    bot.Service
	BotCache   types.Cache

	BotStepper *stepper.Stepper

	In T
}

func getBaseInjectors[T any]() []Injector[T] {
	return []Injector[T]{
		injectors.WithRedis[T],
		injectors.WithLogger[T],
		injectors.WithRepository[T],
		injectors.WithService[T],
		injectors.WithBotCache[T],
		injectors.WithBotStepper[T],
	}
}

func getBaseProbes[T any](components *Components[T]) []Probe {
	return []Probe{
		func(ctx context.Context) error {
			return components.Redis.Ping(ctx).Err()
		},
	}
}
