package utils

import (
	"fmt"
	"os"
	"os/exec"
)

func UpdateRun() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		os.Exit(1)
	}

	cmd := exec.Command("cmd", "/c", "start", "/D", dir, "bot_kz_updater.exe")
	err = cmd.Start()
	if err != nil {
		fmt.Println("Error starting process:", err)
		os.Exit(1)
	}
	fmt.Println("Process started successfully.")
}
