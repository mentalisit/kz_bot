package utils

import (
	"fmt"
	"os/exec"
	"strings"
	"time"
)

func PanicAndDown() {
	killProcces("cmd.exe")
	killProcces("kz_bot.exe")
	//runBot()

}

func killProcces(procces string) {
	cmd := exec.Command("tasklist")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		return
	}
	if string(output) != "" {
		for _, line := range strings.Split(string(output), "\n") {
			if strings.Contains(line, procces) {
				fields := strings.Fields(line)
				pid := fields[1]
				cmd := exec.Command("taskkill", "/F", "/PID", pid)
				if err := cmd.Run(); err != nil {
					fmt.Println("Ошибка при завершении процесса", fields[0], "с PID", pid, ":", err)
				} else {
					fmt.Println("Процесс", fields[0], "с PID", pid, "был завершен")
					time.Sleep(1 * time.Second)
				}
			}
		}
	}
}
