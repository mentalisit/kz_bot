package config

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"log"
	"net"
	"os/exec"
	"time"
)

func Reserv(log *logrus.Logger) {
	if cfg.BotMode == "reserve" {
		for {
			if checkPing() {
				fmt.Println(time.Now().Format("15:04 02-01-2006"), " ожидание")
				time.Sleep(1 * time.Minute)
			} else {
				go func() {
					for {
						if checkPing() {
							log.Println("Сервер доступен, пвникую ")
							connectRDP()
							time.Sleep(1 * time.Minute)
							panic("dostupen")
						}
					}
				}()
				break
			}
		}
	}
}
func checkPing() bool {
	if ping("google.com:80") && ping(cfg.ServerAdrr) {
		return true
	} else {
		time.Sleep(5 * time.Second)
		if ping("google.com:80") && ping(cfg.ServerAdrr) {
			return true
		} else if ping("google.com:80") && !ping(cfg.ServerAdrr) {
			return false
		} else if !ping("google.com:80") && !ping(cfg.ServerAdrr) {
			time.Sleep(10 * time.Second)
			panic("нет интернета")
		}
	}
	return false
}
func ping(address string) bool {
	timeout := time.Duration(3 * time.Second)

	_, err := net.DialTimeout("tcp", address, timeout)
	if err == nil {
		return true
	}
	return false
}

func connectRDP() {
	cmd := exec.Command("cmd.exe", "/C", "start", "D:\\config\\1.bat")
	cmd.Start()
	log.Printf("Пробуем перезапуск .")
}
