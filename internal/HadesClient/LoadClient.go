package HadesClient

import (
	"fmt"
	"github.com/mitchellh/go-ps"
	"kz_bot/internal/config"
	"os"
	"os/exec"
	"time"
)

func (h *Hades) reloadConsoleClient() {
	if config.Instance.BotMode == "dev" {
		return
	}
	var s []string
	for _, client := range h.corporation {
		s = append(s, client.Corp)
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
