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
		panic(err.Error())
	}

	//db := dbasePostgres.Db{}
	//db.InitPostrges()

	//tf := &telegraf.Telegraf{}
	//tf.InitTelegraf(log)
}
func Run() error {
	cfg, err := config.InitConfig()
	if err != nil {
		return err
	}
	log := logger.NewLoggerTG(cfg.LogToken, cfg.LogChatId)
	log.Println("🚀  загрузка  🚀")

	db, errd := dbase.NewDb(cfg, log)
	if errd != nil {
		return errd
	}

	db.CorpConfig.ReadBotCorpConfig()
	cl := clients.NewClient(cfg, db, log)

	go bot.NewBot(*cl, db, log).InitBot()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	fmt.Println("остановка")
	db.Shutdown()
	return nil
}
