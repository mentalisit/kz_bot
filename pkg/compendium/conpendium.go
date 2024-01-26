package compendium

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

//{Type:rs Level:9 Ws:0}
//{Type:shipmentrelay Level:9 Ws:0}
//{Type:corplevel Level:17 Ws:0}
//{Type:transp Level:5 Ws:0}
//{Type:miner Level:5 Ws:0}
//{Type:bs Level:5 Ws:0}
//{Type:cargobay Level:10 Ws:0}
//{Type:computer Level:7 Ws:0}
//{Type:remoterepair Level:2 Ws:0}
//{Type:rush Level:5 Ws:0}
//{Type:stealth Level:7 Ws:0}
//{Type:tradeburst Level:7 Ws:0}
//{Type:shipdrone Level:8 Ws:0}
//{Type:rsextender Level:8 Ws:0} {Type:relicdrone Level:12 Ws:0} {Type:dispatch Level:7 Ws:0} {Type:miningboost Level:9 Ws:0} {Type:hydroreplicator Level:10 Ws:0} {Type:artifactboost Level:7 Ws:0} {Type:remote Level:8 Ws:0} {Type:genesis Level:11 Ws:0} {Type:enrich Level:7 Ws:0} {Type:crunch Level:1 Ws:0} {Type:laser Level:9 Ws:0} {Type:mass Level:11 Ws:0} {Type:battery Level:11 Ws:0} {Type:dual Level:4 Ws:0} {Type:barrage Level:8 Ws:0} {Type:alpha Level:2 Ws:0} {Type:delta Level:6 Ws:0} {Type:passive Level:10 Ws:0} {Type:omega Level:9 Ws:0} {Type:blast Level:8 Ws:0} {Type:mirror Level:1 Ws:0} {Type:area Level:5 Ws:0} {Type:emp Level:7 Ws:0} {Type:solitude Level:10 Ws:0} {Type:fortify Level:3 Ws:0} {Type:teleport Level:8 Ws:0} {Type:damageamplifier Level:9 Ws:0} {Type:destiny Level:7 Ws:0} {Type:barrier Level:8 Ws:0} {Type:vengeance Level:7 Ws:0} {Type:deltarocket Level:3 Ws:0} {Type:leap Level:7 Ws:0} {Type:bond Level:2 Ws:0} {Type:suspend Level:1 Ws:0} {Type:omegarocket Level:1 Ws:0} {Type:decoydrone Level:7 Ws:0} {Type:repairdrone Level:4 Ws:0} {Type:rocketdrone Level:9 Ws:0} {Type:chainrayturret Level:1 Ws:0}
