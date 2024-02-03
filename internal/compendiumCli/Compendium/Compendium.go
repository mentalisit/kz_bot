package Compendium

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"kz_bot/pkg/logger"

	//"kz_bot/internal/compendiumCli"
	"kz_bot/internal/compendiumCli/bot_api"
	"kz_bot/internal/compendiumCli/module_types"
	"kz_bot/internal/models"
	"path/filepath"
	"time"
)

const RefreshInterval = 5 * time.Minute

func NewCompendium(log *logger.Logger, StorageKey string) *Compendium {
	return &Compendium{
		Client:     *bot_api.NewCompendiumApiClient(log), // Assuming you have NewCompendiumApiClient defined
		StorageKey: StorageKey,
		log:        log,
	}
}

func (c *Compendium) GetUser() *models.User {
	if c.Ident.User.Username != "" {
		return &c.Ident.User
	}
	return nil
}

func (c *Compendium) GetGuild() *models.Guild {
	if c.Ident.Guild.Name != "" {
		return &c.Ident.Guild
	}
	return nil
}

func (c *Compendium) GetTechLevels() map[int]models.TechLevel {
	if len(c.SyncData.TechLevels) > 0 {
		return c.SyncData.TechLevels
	}
	return nil
}

func (c *Compendium) Initialize() error {
	c.Ident = new(models.Identity)
	c.Ident = c.ReadStorage()
	if c.Ident != nil && c.Ident.Token != "" {
		if c.SyncData == nil || len(c.SyncData.TechLevels) == 0 {
			c.SyncData = &models.SyncData{}
			c.SyncUserData("get")
		} else {
			c.SyncUserData("sync")
		}
		// Emit "connected" event

	} else if c.Ident == nil {
		return errors.New("net token")
	}
	c.Ticker = time.NewTicker(RefreshInterval)
	go c.Tick()
	return nil
}

func (c *Compendium) Shutdown() {
	if c.Ticker != nil {
		c.Ticker.Stop()
	}
}

func (c *Compendium) CheckConnectCode(code string) (*models.Identity, error) {
	i, err := c.Client.CheckIdentity(code)
	iden := models.Identity{
		User:  i.User,
		Guild: i.Guild[0],
		Token: i.Token,
	}
	return &iden, err
}

func (c *Compendium) Connect(ident *models.Identity) (*models.Identity, error) {
	//c.ClearData()
	//c.log.Info(fmt.Sprintf("Ident %+v", ident))
	c.Ident, _ = c.Client.Connect(ident)
	// Emit "connected" event
	c.LastTokenRefresh = time.Now().Unix()
	c.WriteStorage()

	c.SyncUserData("get")
	return c.Ident, nil
}

func (c *Compendium) Logout() {
	// Emit "disconnected" event
	c.ClearData()
}

func (c *Compendium) CorpData(roleID string) (*models.CorpData, error) {
	if c.Ident.Token == "" {
		return nil, errors.New("not connected")
	}
	return c.Client.CorpData(c.Ident.Token, roleID)
}

func (c *Compendium) SetTechLevel(techID int, level int) error {
	if c.Ident.Token == "" {
		return errors.New("not connected")
	}
	if module_types.GetTechFromIndex(techID) == "" {
		return errors.New("Invalid tech id")
	}

	if len(c.SyncData.TechLevels) == 0 {
		c.SyncData = &models.SyncData{Ver: 1, InSync: 1, TechLevels: make(map[int]models.TechLevel)}
	}
	c.SyncData.TechLevels[techID] = models.TechLevel{Level: level, Ts: time.Now().Unix()}
	c.SyncUserData("sync")
	return nil
}

func (c *Compendium) WriteStorage() {
	if c.Ident.Token == "" {
		return
	}
	//c.log.Info(fmt.Sprintf(" WriteStorage c.Ident %+v\n", c.Ident))
	//c.log.Info(fmt.Sprintf(" WriteStorage c.Ident.Guild %+v\n", c.Ident.Guild))
	data := models.StorageData{
		Ident:        c.Ident,
		UserData:     c.SyncData,
		Refresh:      c.LastRefresh,
		TokenRefresh: c.LastTokenRefresh,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		c.log.ErrorErr(err)
		return
	}
	identPath := filepath.Join(".", "config", c.StorageKey)
	err = ioutil.WriteFile(identPath, jsonData, 0644)
	if err != nil {
		c.log.ErrorErr(err)
		return
	}

	return
}

func (c *Compendium) ReadStorage() *models.Identity {
	// Загрузка существующей идентификации
	identPath := filepath.Join(".", "config", c.StorageKey)
	identBytes, err := ioutil.ReadFile(identPath)
	if err != nil {
		c.log.Info(err.Error())
		return &models.Identity{}
	}
	var stored models.StorageData
	err = json.Unmarshal(identBytes, &stored)
	if err != nil {
		c.log.Info(errors.New("нет сохраненной сессии").Error())
		c.ClearData()
		return &models.Identity{}
	}
	c.Ident = stored.Ident
	c.SyncData = stored.UserData
	c.LastRefresh = stored.Refresh
	c.LastTokenRefresh = stored.TokenRefresh
	return stored.Ident
}

func (c *Compendium) ClearData() {
	// Clear local storage
	c.Ident = &models.Identity{}
	c.LastTokenRefresh = 0
	c.LastRefresh = 0
	c.SyncData = &models.SyncData{}
}

func (c *Compendium) SyncUserData(mode string) {
	if c.Ident.Token == "" || (mode != "get" && len(c.SyncData.TechLevels) == 0) {
		// Emit error event
		return
	}
	if len(c.SyncData.TechLevels) == 0 {

		c.SyncData = &models.SyncData{
			Ver:    1,
			InSync: 1,
		}
		//c.log.Info(string("len(c.SyncData.TechLevels)"))
	}

	sync, err := c.Client.Sync(c.Ident.Token, mode, c.SyncData.TechLevels)
	if err != nil {
		c.log.ErrorErr(err)
		return
	}
	c.SyncData = &sync

	c.LastRefresh = time.Now().Unix()
	c.WriteStorage()
	// Emit "sync" event
}

func (c *Compendium) Tick() {
	for range c.Ticker.C {
		if c.Ident.Token != "" {
			if time.Now().Unix()-c.LastTokenRefresh > int64(RefreshInterval) {
				// three months
				newIdent, err := c.Client.RefreshConnection(c.Ident.Token)
				if err == nil {
					c.Ident = newIdent
					c.LastTokenRefresh = time.Now().Unix()
					c.WriteStorage()
				} else {
					c.ClearData()
					// Emit "connectfailed" event
				}
			}
			if time.Now().Unix()-c.LastRefresh > int64(RefreshInterval) {
				c.SyncUserData("sync")
			}
		}
	}
}
