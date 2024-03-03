package bot_api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/mentalisit/logger"
	"io/ioutil"
	"kz_bot/internal/models"
	"net/http"
)

// CompendiumApiClient represents a client for the Compendium API
type CompendiumApiClient struct {
	URL string
	log *logger.Logger
}

func NewCompendiumApiClient(log *logger.Logger) *CompendiumApiClient {
	return &CompendiumApiClient{
		URL: "https://bot.hs-compendium.com/compendium",
		log: log,
	}
}

// CheckIdentity validates the code and returns a token and identity
func (c *CompendiumApiClient) CheckIdentity(code string) (*models.IdentityGET, error) {
	apiURL := "https://bot.hs-compendium.com/compendium/applink/identities?ver=2&code=1"

	// Подготовка запроса
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return &models.IdentityGET{}, err
	}

	// Установка параметров запроса
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Authorization", code)

	// Отправка запроса
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return &models.IdentityGET{}, err
	}
	defer resp.Body.Close()

	// Проверка успешного ответа
	if resp.StatusCode < 200 || resp.StatusCode >= 500 {
		return &models.IdentityGET{}, errors.New("Server Error")
	}

	// Чтение тела ответа
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &models.IdentityGET{}, err
	}

	if len(body) < 50 {
		return &models.IdentityGET{}, errors.New("Invalid User Id")
	}

	// Декодирование JSON-строки в структуру Identity
	var identity models.IdentityGET
	err = json.Unmarshal(body, &identity)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return &models.IdentityGET{}, err
	}
	return &identity, nil
}

// Connect gets a connection token and identity
func (c *CompendiumApiClient) Connect(identity *models.Identity) (*models.Identity, error) {
	url := fmt.Sprintf("%s/applink/connect", c.URL)
	c.log.Info("guild_id " + identity.Guild.ID)
	data := map[string]interface{}{
		"guild_id": identity.Guild.ID,
	}

	payload, err := json.Marshal(data)
	if err != nil {
		return &models.Identity{}, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return &models.Identity{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", identity.Token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return &models.Identity{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &models.Identity{}, err
	}
	//c.log.Info(string(body))
	if resp.StatusCode < 200 || resp.StatusCode >= 500 {
		return &models.Identity{}, errors.New("Server Error")
	}

	var obj models.Identity
	if err := json.Unmarshal(body, &obj); err != nil {
		return &models.Identity{}, err
	}
	c.log.Info(fmt.Sprintf("string(body) %+v\n", string(body)))
	c.log.Info(fmt.Sprintf("obj %+v\n", obj.Guild))
	return &obj, nil
	//if resp.StatusCode >= 400 {
	//	return &models.Identity{}, errors.New(obj["error"].(string))
	//}
	//return &models.Identity{
	//	User: obj["user"].(models.User),
	//	Guild: []models.Guild{models.Guild{
	//		ID:   obj["guild"].(models.Guild).ID,
	//		Name: obj["guild"].(models.Guild).Name,
	//		URL:  fmt.Sprintf("https://cdn.discordapp.com/avatars/%s/%s.png", obj["user"].(models.User).ID, obj["user"].(models.User).Avatar),
	//	}},
	//	Token: obj["token"].(string),
	//}, nil
}

// RefreshConnection refreshes the connection token
func (c *CompendiumApiClient) RefreshConnection(token string) (*models.Identity, error) {
	endpoint := fmt.Sprintf("%s/applink/refresh", c.URL)
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 500 {
		return nil, errors.New("Server Error")
	}

	var obj models.Identity
	if err := json.NewDecoder(resp.Body).Decode(&obj); err != nil {
		return nil, err
	}

	obj.Guild.URL = fmt.Sprintf("https://cdn.discordapp.com/avatars/%s/%s.png", obj.User.ID, obj.User.Avatar)
	return &obj, nil
}

// CorpData retrieves various data for all members in the corp
func (c *CompendiumApiClient) CorpData(token string, roleID string) (*models.CorpData, error) {
	endpoint := fmt.Sprintf("%s/cmd/corpdata?roleId=%s", c.URL, roleID)
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 500 {
		return nil, errors.New("Server Error")
	}

	var obj models.CorpData
	if err := json.NewDecoder(resp.Body).Decode(&obj); err != nil {
		return nil, err
	}

	return &obj, nil
}

// Sync synchronizes tech levels with the bot
func (c *CompendiumApiClient) Sync(token string, mode string, currentTech map[int]models.TechLevel) (models.SyncData, error) {
	if mode != "get" && mode != "set" && mode != "sync" {
		return models.SyncData{}, fmt.Errorf("Invalid sync mode %s", mode)
	}

	if mode == "get" {
		currentTech = map[int]models.TechLevel{}
	}

	url := fmt.Sprintf("%s/cmd/syncTech/%s", c.URL, mode)

	data := map[string]interface{}{
		"ver":        1,
		"techLevels": currentTech,
	}

	payload, err := json.Marshal(data)
	if err != nil {
		return models.SyncData{}, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return models.SyncData{}, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return models.SyncData{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return models.SyncData{}, err
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 500 {
		return models.SyncData{}, errors.New("Server Error")
	}

	var obj models.SyncData
	if err := json.Unmarshal(body, &obj); err != nil {
		return models.SyncData{}, err
	}

	if resp.StatusCode >= 400 {
		return models.SyncData{}, errors.New("obj.Error")
	}

	return obj, nil
}
