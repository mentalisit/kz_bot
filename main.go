package main

import (
	"fmt"
	"kz_bot/internal/BridgeChat"
	"kz_bot/internal/HadesClient"
	"kz_bot/internal/bot"
	"kz_bot/internal/clients"
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

	//создаем логгер в телегу
	log := logger.NewLoggerTG(cfg.Logger.Token, cfg.Logger.ChatId)

	if cfg.BotMode == "dev" {
		fmt.Println("Develop Running")
		//test func
		st := storage.NewStorage(log, cfg)
		cl := clients.NewClients(log, st, cfg)

		HadesClient.NewHades(log, cl, st)
		fmt.Scanln()
		//time.Sleep(1 * time.Minute)
		return nil
	}

	//Если запуск на резервном сервере то блокируем выполнение
	config.Reserv(log)

	log.Println("🚀  загрузка  🚀 " + cfg.BotMode)

	//storage
	st := storage.NewStorage(log, cfg)

	//clients Discord, Telegram, //Whatsapp
	cl := clients.NewClients(log, st, cfg)
	go bot.NewBot(st, cl, log, cfg)

	go BridgeChat.NewBridge(log, cl, st)

	//eti udalit
	//go GlobalChat.NewChat(st, cl, log)
	//go relay.NewRelay(log, st, cl)

	//ожидаем сигнала завершения
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	return nil
}

//func (h *Hades) reloadConsoleClient(s []string) {
//
//	procs, err := ps.Processes()
//	if err != nil {
//		h.log.Println("reloadConsoleClient " + err.Error())
//		return
//	}
//	for _, proc := range procs {
//		if proc.Executable() == "ConsoleClient.exe" {
//			fmt.Println("Restarting hsBot")
//			cmd := exec.Command("taskkill", "/F", "/IM", "ConsoleClient.exe")
//			_ = cmd.Run()
//		}
//	}
//
//	err = os.Chdir("hsbot")
//	if err != nil {
//		h.log.Println("reloadConsoleClient " + err.Error())
//		return
//	}
//	for i := 0; i < len(s); i++ {
//		time.Sleep(5 * time.Second)
//		err = os.Chdir(s[i])
//		cmd := exec.Command("cmd.exe", "/c", "start", "./ConsoleClient.exe")
//		err = cmd.Run()
//		time.Sleep(10 * time.Second)
//		if err != nil {
//			h.log.Println("run console client " + err.Error())
//			return
//		}
//		err = os.Chdir("..")
//		if err != nil {
//			h.log.Println(err.Error())
//		}
//	}
//}
