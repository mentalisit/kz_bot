package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"kz_bot/config"
	"kz_bot/internal/bot"
	"kz_bot/internal/clients"
	"kz_bot/internal/dbase"
	"kz_bot/internal/logger"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	fmt.Println("ЗАПУСК БОТА")
	err := Run()
	if err != nil {
		fmt.Println("Ошибка запуска бота", err)
		time.Sleep(1 * time.Second)
		panic(err.Error())
	}
}

func Run() (err error) {
	//читаем конфигурацию с ENV
	cfg, err := config.InitConfig()
	if err != nil {
		return err
	}

	//создаем логгер в телегу
	log := logger.NewLoggerTG(cfg.LogToken, cfg.LogChatId)

	//Если запуск на резервном сервере то блокируем выполнение
	config.Reserv(log)

	err = runLogicBot(cfg, log)
	if err != nil {
		return err
	}

	return err
}
func runLogicBot(cfg config.ConfigBot, log *logrus.Logger) error {
	log.Println("🚀  загрузка  🚀")
	//подключаюсь к базе ланных
	db, errd := dbase.NewDb(cfg, log)
	if errd != nil {
		return errd
	}

	//читаю конфиг корпораций
	db.CorpConfig.ReadBotCorpConfig()
	//запускаю месенджеры
	cl := clients.NewClient(cfg, db, log)

	//запускаю основную логику бота
	go bot.NewBot(*cl, db, log, cfg.Debug).InitBot()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	fmt.Println("остановка")
	db.Shutdown()
	return nil
}
