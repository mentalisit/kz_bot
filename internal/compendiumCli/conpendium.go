package compendiumCli

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type TechnicalData struct {
	Map   map[string]Item `json:"map"`
	Array []Item          `json:"array"`
}

type Item struct {
	Type  string `json:"type"`
	Level int    `json:"level"`
	Ws    int    `json:"ws"`
}

func GetUserId(userID string) (genesis, enrich, rsextender int) {
	apiKey := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiI1ODI4ODIxMzc4NDIxMjI3NzMiLCJndWlsZElkIjoiNzE2NzcxNTc5Mjc4OTE3NzAyIiwiaWF0IjoxNzA2MjM3MzY0LCJleHAiOjE3Mzc3OTQ5NjQsInN1YiI6ImFwaSJ9.Wsf-2U8GDGaCNpxafRIUABIKO3zLyYKvPYWzxtbK-LE"

	// Формирование URL-адреса
	url := fmt.Sprintf("https://bot.hs-compendium.com/compendium/api/tech?token=%s&userid=%s", apiKey, userID)

	// Выполнение GET-запроса
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Ошибка при выполнении запроса:", err)
		return
	}
	defer response.Body.Close()

	// Проверка кода ответа
	if response.StatusCode != http.StatusOK {
		fmt.Printf("Неправильный статус код: %d\n", response.StatusCode)
		return
	}

	// Чтение тела ответа
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Ошибка при чтении ответа:", err)
		return
	}

	// Декодирование JSON-данных в структуру TechnicalData
	var technicalData TechnicalData
	err = json.Unmarshal(body, &technicalData)
	if err != nil {
		fmt.Println("Ошибка при декодировании JSON:", err)
		return
	}
	genesis = technicalData.Map["genesis"].Level
	enrich = technicalData.Map["enrich"].Level
	rsextender = technicalData.Map["rsextender"].Level
	return
}
