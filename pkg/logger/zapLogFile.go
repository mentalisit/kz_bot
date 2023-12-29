package logger

import (
	"fmt"
	"os"
)

// createLogFile создает файл с логами и возвращает его FileWriteSyncer
func createLogFile(fileName string) *os.File {
	err := ensureLogDir()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic("Не удалось создать файл с логами: " + err.Error())
	}
	return file
}

// ensureLogDir создает директорию "log", если она не существует
func ensureLogDir() error {
	logDir := "log"
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		// Создаем директорию
		if err := os.Mkdir(logDir, os.ModePerm); err != nil {
			return fmt.Errorf("ошибка при создании директории: %v", err)
		}
		fmt.Printf("Директория %s успешно создана\n", logDir)
	}
	return nil
}
