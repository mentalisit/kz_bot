package compendium

import (
	"fmt"
	"kz_bot/internal/models"
)

func GetCompendiumStruct() string {
	tokenCompendium := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiI1ODI4ODIxMzc4NDIxMjI3NzMiLCJndWlsZElkIjoiNjczODYwOTU2Njk4NDQzODA2IiwiaWF0IjoxNjgxNDY0NzM0LCJleHAiOjE3MTMwMjIzMzQsInN1YiI6ImFwaSJ9.8JekfxwHxezraNTeT-i1ediDIapFkh2ELtMCULEV9Ek"
	UserId := "542249167641247775"
	list := GetCompendiumData(tokenCompendium, UserId)
	list = ifType(list)
	var s string
	s = fmt.Sprintf("%-15s%-10s%-5s\n", "Технологии", "Уровень", "Очки БЗ")
	for _, i := range list.Array {
		if i.Level != 0 {
			s = s + fmt.Sprintf("%-15s%-10d%-5d\n", i.Type, i.Level, i.Ws)
		}
	}
	fmt.Println(s)
	return s
}
func ifType(list models.JsonCompendium) models.JsonCompendium {
	var newArray models.JsonCompendium
	for _, array := range list.Array {
		var newstruct models.Array
		switch array.Type {
		case "bs":
			newstruct.Type = "Линкор"
		case "rs":
			newstruct.Type = "КЗ"
		case "corplevel":
			newstruct.Type = "Корпорация"
		case "transp":
			newstruct.Type = "Транспорт"
		case "miner":
			newstruct.Type = "Майнер"
		case "computer":
			newstruct.Type = "Гр.компьютер"
		case "shipmentrelay":
			newstruct.Type = "Переправка гр."
		case "relicdrone":
			newstruct.Type = "ДронРеликов"
		case "battery":
			newstruct.Type = "Батарея"

		default:
			newstruct.Type = array.Type
		}
		newstruct.Ws = array.Ws
		newstruct.Level = array.Level
		newArray.Array = append(newArray.Array, newstruct)
	}
	return newArray
}
