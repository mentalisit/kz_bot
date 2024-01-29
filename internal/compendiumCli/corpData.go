package compendiumCli

import (
	"fmt"
	"kz_bot/internal/compendiumCli/Compendium"
	"kz_bot/pkg/logger"
)

const (
	StorageKey = "hscompendium" //"identity.json"
)

var roleID = "718194885269782669"

func MainCorp(log *logger.Logger, code string) {
	c := Compendium.NewCompendium(log, StorageKey)

	err := c.Initialize()
	if err != nil {
		log.Error(err.Error())
		return
	}
	log.Info(fmt.Sprintf("c.Initialize %+v %+v", c.Ident.User.Username, c.Ident.Guild[0].ID))
	if c.Ident.Token == "" {
		c.Ident, err = c.CheckConnectCode(code)
		if err != nil {
			log.Error(err.Error())
			return
		}

		log.Info(fmt.Sprintf("c.Ident %+v", c.Ident.User.Username))
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
	log.Info(fmt.Sprintf("%+v\n", c.SyncData))
}
