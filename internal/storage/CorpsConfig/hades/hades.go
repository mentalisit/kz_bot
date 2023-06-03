package hades

import (
	"context"
	"fmt"
	"github.com/mitchellh/go-ps"
	"github.com/sirupsen/logrus"
	"kz_bot/internal/config"
	"kz_bot/internal/models"
	"kz_bot/pkg/clientDB/postgresqlS"
	"os"
	"os/exec"
	"time"
)

var HadesStorage *Hades

type Hades struct {
	client postgresqlS.Client
	log    *logrus.Logger
}

func NewHades(client postgresqlS.Client, log *logrus.Logger) *Hades {
	h := &Hades{client: client, log: log}
	HadesStorage = h
	go h.ReadCorporation()
	return h
}
func (h *Hades) ReadCorporation() {

	read :=
		`SELECT * FROM kzbot.bridge`
	results, err := h.client.Query(context.Background(), read)
	if err != nil {
		h.log.Println("Ошибка чтения крнфигурации корпораций hades", err)
	}
	var clients []string
	fmt.Printf("hades: ")
	for results.Next() {
		var t models.Corporation
		err = results.Scan(&t.Id, &t.Corp, &t.DsChat, &t.DsChatWS1, &t.DsChatWS2,
			&t.GuildId, &t.TgChat, &t.TgChatWS1)
		corps = append(corps, t)
		clients = append(clients, t.Corp)
		fmt.Printf(" %s ,", t.Corp)
	}
	fmt.Println()
	go h.reloadConsoleClient(clients)

}
func (h *Hades) reloadConsoleClient(s []string) {
	if config.Instance.BotMode != "server" {
		return
	}
	procs, err := ps.Processes()
	if err != nil {
		h.log.Println("reloadConsoleClient " + err.Error())
		return
	}
	for _, proc := range procs {
		if proc.Executable() == "ConsoleClient.exe" {
			fmt.Println("Restarting hsBot")
			cmd := exec.Command("taskkill", "/F", "/IM", "ConsoleClient.exe")
			_ = cmd.Run()
		}
	}

	err = os.Chdir("hsbot")
	if err != nil {
		h.log.Println("reloadConsoleClient " + err.Error())
		return
	}
	for i := 0; i < len(s); i++ {
		time.Sleep(5 * time.Second)
		err = os.Chdir(s[i])
		cmd := exec.Command("cmd.exe", "/c", "start", "./ConsoleClient.exe")
		err = cmd.Run()
		time.Sleep(10 * time.Second)
		if err != nil {
			h.log.Println("run console client " + err.Error())
			return
		}
		err = os.Chdir("..")
		if err != nil {
			h.log.Println(err.Error())
		}
	}
}
