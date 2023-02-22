package main

import (
	"fmt"
	"kz_bot/internal/bot"
	"kz_bot/internal/clients"
	config2 "kz_bot/internal/config"
	"kz_bot/internal/storage"
	"kz_bot/pkg/logger"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	fmt.Println("Bot loading")
	err := RunNew()
	if err != nil {
		fmt.Println("Error loading bot", err)
		time.Sleep(10 * time.Second)
		panic(err.Error())

	}
}

//func Run() (err error) {
//	//читаем конфигурацию с ENV
//	cfg, err := config.InitConfig()
//	if err != nil {
//		return err
//	}
//
//	//создаем логгер в телегу
//	log := logger.NewLoggerTG(cfg.LogToken, cfg.LogChatId)
//
//	//Если запуск на резервном сервере то блокируем выполнение
//	config.Reserv(log)
//
//	err = runLogicBot(cfg, log)
//	if err != nil {
//		return err
//	}
//
//	return err
//}
//func runLogicBot(cfg config.ConfigBot, log *logrus.Logger) error {
//	log.Println("🚀  загрузка  🚀 " + cfg.BotMode)
//
//	//подключаюсь к базе ланных
//	db, errd := dbase.NewDb(cfg, log)
//	if errd != nil {
//		return errd
//	}
//
//	//читаю конфиг корпораций
//	db.CorpConfig.ReadBotCorpConfig()
//
//	//запускаю месенджеры
//	cl := cliennt.NewClient(cfg, db, log)
//
//	//запускаю основную логику бота
//	go bot.NewBot(*cl, db, log, cfg.Debug).InitBot()
//
//	quit := make(chan os.Signal, 1)
//	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
//	<-quit
//	fmt.Println("остановка")
//	db.Shutdown()
//	return nil
//}

func RunNew() error {
	//читаем конфигурацию с ENV
	cfg := config2.InitConfig()

	//создаем логгер в телегу
	log := logger.NewLoggerTG(cfg.Logger.Token, cfg.Logger.ChatId)

	//Если запуск на резервном сервере то блокируем выполнение
	config2.Reserv(log)

	log.Println("🚀  загрузка  🚀 " + cfg.BotMode)

	//storage
	st := storage.NewStorage(log, cfg)

	//clients Discord, Telegram, //Whatsapp
	cl := clients.NewClients(log, st, cfg)

	go bot.NewBot(st, cl, log, cfg)

	//ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	//defer cancel()

	//ожидаем сигнала завершения
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	return nil
}
