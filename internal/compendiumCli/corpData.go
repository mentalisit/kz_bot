package compendiumCli

import (
	"fmt"
	"kz_bot/internal/compendiumCli/Compendium"
	"kz_bot/internal/compendiumCli/module_types"
	"kz_bot/pkg/logger"
)

const (
	StorageKey = "hscompendium" //"identity.json"
)

var roleID = "718194885269782669"

func MainCorp(log *logger.Logger, code string) {
	code = "gnYY-seZK-ufig"
	c := Compendium.NewCompendium(log, StorageKey)

	err := c.Initialize()
	if err != nil {
		log.Error(err.Error())
		return
	}
	if c.Ident.Token == "" {
		c.Ident, err = c.CheckConnectCode(code)
		if err != nil {
			log.Error(err.Error())
			return
		}

		log.Info(fmt.Sprintf("c.Ident %+v", c.Ident))
		connect, err := c.Connect(c.Ident)
		if err != nil {
			log.Error(err.Error())
			return
		}
		c.Ident, err = c.Client.RefreshConnection(connect.Token)
		if err != nil {
			fmt.Println(err)
			return
		}
		c.WriteStorage()
		log.Info(fmt.Sprintf("connect %+v", connect))

	}
	for i, level := range c.SyncData.TechLevels {
		n := module_types.GetTechFromIndex(i)
		log.Info(fmt.Sprintf("%s %d", n, level.Level))
	}
	log.Info(fmt.Sprintf("%+v\n", "ok"))

}
