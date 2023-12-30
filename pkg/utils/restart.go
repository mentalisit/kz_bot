package utils

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func RestartRun() {
	executable, err := os.Executable()
	if err != nil {
		fmt.Println(err)
		return
	}
	cmd := exec.Command("cmd", "/c", "start", "/D", filepath.Dir(executable), "kz.bat")
	err = cmd.Start()
	if err != nil {
		fmt.Println("Error starting process:", err)
		os.Exit(1)
	}
	fmt.Println("Process started successfully.")
}
func RestorePanic() {
	if r := recover(); r != nil {
		fmt.Println("Возникла паника:", r)
		RestartRun()
		os.Exit(1)
	}
}
