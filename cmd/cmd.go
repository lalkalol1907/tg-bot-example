package cmd

import (
	watcher "bot-test/internal/registry/in/chat-watcher"
	"bot-test/internal/registry/in/polling"
	"bot-test/internal/registry/in/webhook"
	"flag"
)

var commands = map[string]func(){
	"polling": polling.Polling,
	"webhook": webhook.Webhook,
	"watcher": watcher.Watcher,
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
