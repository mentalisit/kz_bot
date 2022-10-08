package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"kz_bot/config"
	"kz_bot/internal/bot"
	"kz_bot/internal/clients"
	"kz_bot/internal/dbase"
	"kz_bot/internal/logger"
	"log"
	"net"
	"os"
	"os/exec"
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

	if cfg.BotMode == "reserve" {
		for {
			if runPing(cfg) {
				fmt.Println(time.Now().UTC(), " ожидание")
				time.Sleep(1 * time.Minute)
			} else {
				go func() {
					for {
						if runPing(cfg) {
							log.Println("Сервер доступен, нужно переключится ")
							ConnectRDP()
						}
						time.Sleep(1 * time.Minute)
						panic("dostupen")
					}
				}()
				err = runLogicBot(cfg, log)
			}
		}
	} else {
		err = runLogicBot(cfg, log)
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
func runPing(cfg config.ConfigBot) bool {
	timeout := time.Duration(1 * time.Second)
	var google, ping bool

	_, err1 := net.DialTimeout("tcp", "8.8.8.8", timeout)
	if err1 == nil {
		google = true
	}
	if google {
		_, err := net.DialTimeout("tcp", cfg.ServerAdrr, timeout)
		if err == nil {
			ping = true
		}
	}
	return ping
}
func ConnectRDP() {
	cmd := exec.Command("cmd.exe", "/C", "start", "D:\\config\\1.bat")
	cmd.Start()
	log.Printf("Пробуем перезапуск .")
}
