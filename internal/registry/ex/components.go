package ex

import (
	"bot-test/config"
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/lalkalol1907/tg-bot-stepper/stepper"
	"github.com/lalkalol1907/tg-bot-stepper/types"
	"github.com/redis/go-redis/v9"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
)

// Components Общие компоненты для всех делойментов
type Components[T interface{}] struct {
	Config *config.Config // Сам инжектится первым
	Http   *echo.Echo     // Сам инжектится для проб сразу после конфига, сам встает. Можно использовать и для запросов

	Logger   *otelzap.Logger
	Redis    *redis.Client
	DB       *sqlx.DB
	BotCache types.Cache

	BotStepper *stepper.Stepper

	In T
}

// Общие инжекторы для всех деплойментов, T - компоненты деплоймента
func getBaseInjectors[T any]() []Injector[T] {
	return []Injector[T]{
		WithRedis[T],
		WithLogger[T],
		WithBotCache[T],
		WithBotStepper[T],
		WithDB[T],
	}
}

// Общие пробы для всех деплойментов
func getBaseProbes[T any](components *Components[T]) []Probe {
	return []Probe{
		func(ctx context.Context) error {
			return components.Redis.Ping(ctx).Err()
		},
	}
}
