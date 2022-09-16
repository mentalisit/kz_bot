package main

import (
	"fmt"
	"kz_bot/config"
	"kz_bot/internal/bot"
	"kz_bot/internal/clients"
	"kz_bot/internal/dbase"
	"kz_bot/internal/logger"
	"net"
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
	if cfg.BotMode == "reserve" {
		for {
			ping := runPing(cfg)
			if ping {
				fmt.Println(time.Now().UTC(), " ожидание")
				time.Sleep(1 * time.Minute)
			} else if !ping {
				go func() {
					for {
						if runPing(cfg) {
							panic("Server ready")
						}
						time.Sleep(1 * time.Minute)
					}
				}()
				err = runLogicBot(cfg)
			}
		}
	} else {
		err = runLogicBot(cfg)
	}

	return err
}
func runLogicBot(cfg config.ConfigBot) error {
	//создаем логгер в телегу
	log := logger.NewLoggerTG(cfg.LogToken, cfg.LogChatId)
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
func runPing(cfg config.ConfigBot) (run bool) {
	timeout := time.Duration(1 * time.Second)
	_, err := net.DialTimeout("tcp", cfg.ServerAdrr, timeout)
	if err != nil {
		return false
	} else {
		return true
	}
}
