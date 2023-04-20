package utils

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func UpdateRun() {
	executable, err := os.Executable()
	if err != nil {
		fmt.Println(err)
		return
	}
	cmd := exec.Command("cmd", "/c", "start", "/D", filepath.Dir(executable), "bot_kz_updater.exe")
	err = cmd.Start()
	if err != nil {
		fmt.Println("Error starting process:", err)
		os.Exit(1)
	}
	fmt.Println("Process started successfully.")
}
