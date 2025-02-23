package cmd

import (
	"bot-test/internal/registry/in/polling"
	"bot-test/internal/registry/in/webhook"
	"flag"
)

var commands = map[string]func(){
	"polling": polling.BotPolling,
	"webhook": webhook.BotWebhook,
}

func Run() {
	// Парсим флаги, находим деплоймент, который нам нужен
	flag.Parse()
	name := flag.Arg(0)

	// Если его нет, выходим
	deployment, ok := commands[name]
	if !ok {
		// print commands
		return
	}

	// Запускаем
	deployment()
}
