package compendium

import (
	"encoding/json"
	"kz_bot/internal/models"
	"net/http"
)

func GetCompendiumData(tokenCompendium, UserId string) models.JsonCompendium {
	url := "https://bot.hs-compendium.com/compendium/api/tech?token=" + tokenCompendium + "&userid=" + UserId

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		// Обработка ошибки
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		// Обработка ошибки
	}

	defer resp.Body.Close()

	var welcome models.JsonCompendium
	err = json.NewDecoder(resp.Body).Decode(&welcome)
	if err != nil {
		// Обработка ошибки
	}
	return welcome
}
