package main

import (
	"flag"
	"fmt"
	"kz_bot/internal/bot"
	"kz_bot/internal/clients"
	"kz_bot/internal/config"
	"kz_bot/internal/hades"
	"kz_bot/internal/hades/ReservCopyPaste"
	"kz_bot/internal/storage"
	"kz_bot/pkg/logger"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var flagVersion = flag.Bool("version", false, "show version")

func main() {
	data, err := os.ReadFile("version.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	flag.Parse()
	if *flagVersion {

		fmt.Println(string(data))
		return
	}
	fmt.Println("Bot loading " + string(data))

	err = RunNew()
	if err != nil {
		fmt.Println("Error loading bot", err)
		time.Sleep(10 * time.Second)
		panic(err.Error())

	}
}

func RunNew() error {
	//читаем конфигурацию с ENV
	cfg := config.InitConfig()

	//создаем логгер в телегу
	log := logger.NewLoggerTG(cfg.Logger.Token, cfg.Logger.ChatId)

	if cfg.BotMode == "dev" {
		fmt.Println("Develop Running")
		//test func
		return nil
	}

	//Если запуск на резервном сервере то блокируем выполнение
	config.Reserv(log)
	//Если нет пинга то загружаем бекап и запускаемся
	if cfg.BotMode == "reserve" {
		ReservCopyPaste.LoadBackup()
	}

	log.Println("🚀  загрузка  🚀 " + cfg.BotMode)

	//storage
	st := storage.NewStorage(log, cfg)

	//clients Discord, Telegram, //Whatsapp
	cl := clients.NewClients(log, st, cfg)
	go hades.NewHades(cl)
	go bot.NewBot(st, cl, log, cfg)

	//ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	//defer cancel()

	//ожидаем сигнала завершения
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	return nil
}
