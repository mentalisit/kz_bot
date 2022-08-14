package main

import (
	"fmt"
	"kz_bot/config"
	"kz_bot/internal/bot"
	"kz_bot/internal/clients"
	"kz_bot/internal/dbase"
	"kz_bot/internal/logger"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	fmt.Println("ЗАПУСК БОТА")
	err := Run()
	if err != nil {
		fmt.Println("Ошибка запуска бота", err)
		panic(err.Error())
	}
}
func Run() error {
	//читаем конфигурацию с ENV
	cfg, err := config.InitConfig()
	if err != nil {
		return err
	}

	//создаем логгер в телегу
	log := logger.NewLoggerTG(cfg.LogToken, cfg.LogChatId)
	log.Println("🚀  загрузка  🚀")

	debug := true //нужно переделать

	//подключаюсь к базе ланных
	db, errd := dbase.NewDb(cfg, log, debug)
	if errd != nil {
		return errd
	}

	//читаю конфиг корпораций
	db.CorpConfig.ReadBotCorpConfig()
	//запускаю месенджеры
	cl := clients.NewClient(cfg, db, log, debug)

	//запускаю основную логику бота
	go bot.NewBot(*cl, db, log, debug).InitBot()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	fmt.Println("остановка")
	db.Shutdown()
	return nil
}
