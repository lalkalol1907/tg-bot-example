package ex

import (
	"bot-test/config"
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
)

// CloseFunc чистящая функция при завершении работы
type CloseFunc = func() error

type Injector[T any] func(c *Components[T]) (CloseFunc, error)
type Runnable = func(ctx context.Context)
type Probe = func(ctx context.Context) error

type DiContainer[T any] struct {
	// Компоненты, T - компоненты деплоймента
	Components *Components[T]
	// Очередь завкрытия при завершении, выполняется в обратном порядке
	closeQueue []CloseFunc

	// Исполняемая функция
	runnable Runnable
	// Функции проб
	probes []Probe
}

// Provide Функция инжекторов
func (d *DiContainer[T]) Provide(injectors ...Injector[T]) *Components[T] {
	c := new(Components[T])
	q := make([]CloseFunc, 0)

	// Парсим конфиг
	c.Config = config.NewConfig()
	if err := c.Config.Parse(); err != nil {
		panic(fmt.Errorf("error parsing config: %v", err))
	}

	// Создаем хттп сервер
	c.Http = echo.New()

	// Добавляем инжекторы деплоймента к базовым
	injectors = append(getBaseInjectors[T](), injectors...)

	// Инжектим по порядку
	for _, inj := range injectors {
		closing, err := inj(c)

		if err != nil {
			panic(fmt.Errorf("dependency error: %v", err))
		}

		// Если есть "чистилка", добавляем в массив
		if closing != nil {
			q = append(q, closing)
		}
	}

	d.Components = c
	d.closeQueue = append(q, func() error {
		return c.Http.Shutdown(context.Background())
	})

	return c
}

func (d *DiContainer[T]) AddProbes(probes ...Probe) {
	d.probes = append(getBaseProbes[T](d.Components), probes...)
}

func (d *DiContainer[T]) runHttpWithProbes() error {
	d.Components.Http.GET("/probe", func(c echo.Context) error {
		//
		return c.String(http.StatusOK, "")
	})
	return d.Components.Http.Start(":" + d.Components.Config.Http.Port)
}

func (d *DiContainer[T]) Run(r Runnable) {
	// Создаем контекст для грейсфула
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	// Слушаем его
	go d.initGracefulShutdown(ctx)

	// Запускаем аппку
	go func() {
		r(ctx)
		d.Components.Logger.Warn("application exited")
		cancel()
	}()

	// ЗЗапускаем хттп (в аппке не надо запускать его)
	if err := d.runHttpWithProbes(); err != nil {
		d.Components.Logger.Fatal("server closed", zap.Error(err))
	}
}

func (d *DiContainer[T]) initGracefulShutdown(ctx context.Context) {
	select {
	case <-ctx.Done():
		{
			d.Components.Logger.Info("graceful shutdown started")
			for i := len(d.closeQueue) - 1; i >= 0; i-- {
				if err := d.closeQueue[i](); err != nil {
					d.Components.Logger.Warn(fmt.Sprintf("error closing component: %s", err.Error()))
				}
			}

			os.Exit(0)
		}
	}
}

func NewDiContainer[T any]() *DiContainer[T] {
	return new(DiContainer[T])
}
