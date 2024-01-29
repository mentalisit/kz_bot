package main

import (
	"fmt"
	"kz_bot/internal/BridgeChat"
	"kz_bot/internal/bot"
	"kz_bot/internal/clients"
	"kz_bot/internal/compendiumCli"
	"kz_bot/internal/config"
	"kz_bot/internal/storage"
	"kz_bot/pkg/logger"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	fmt.Println("Bot loading ")
	//defer utils.RestorePanic()

	err := RunNew()
	if err != nil {
		fmt.Println("Error loading bot", err)
		time.Sleep(10 * time.Second)
		panic(err.Error())
	}
}

func RunNew() error {
	//читаем конфигурацию с ENV
	cfg := config.InitConfig()

	//создаем логгер
	log := logger.LoggerZap(cfg.Logger.Token, cfg.Logger.ChatId)

	if cfg.BotMode == "dev" {
		log = logger.LoggerZapDEV()
		compendiumCli.MainCorp(log, "yYqY-DJ4c-LCce")
		os.Exit(1)
		//time.Sleep(5 * time.Second)
		time.Sleep(5 * time.Minute)
	}

	//Если запуск на резервном сервере то блокируем выполнение
	//config.Reserv(log)

	log.Info("🚀  загрузка  🚀 " + cfg.BotMode)

	//storage
	st := storage.NewStorage(log, cfg)

	//clients Discord, Telegram, //Whatsapp
	cl := clients.NewClients(log, st, cfg)
	go bot.NewBot(st, cl, log, cfg)
	//go HadesClient.NewHades(log, cl, st)
	go BridgeChat.NewBridge(log, cl, st)

	//ожидаем сигнала завершения
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	return nil
}
