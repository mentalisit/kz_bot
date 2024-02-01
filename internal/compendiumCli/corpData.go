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
	code = "bhfa-rQGy-nabt"
	c := Compendium.NewCompendium(log, StorageKey)

	err := c.Initialize()
	if err != nil {
		log.ErrorErr(err)
		return
	}
	if c.Ident.Token == "" {
		c.Ident, err = c.CheckConnectCode(code)
		if err != nil {
			log.ErrorErr(err)
			return
		}

		log.Info(fmt.Sprintf("c.Ident %+v", c.Ident))
		connect, err := c.Connect(c.Ident)
		if err != nil {
			log.ErrorErr(err)
			return
		}
		c.Ident, err = c.Client.RefreshConnection(connect.Token)
		if err != nil {
			log.ErrorErr(err)
			return
		}
		c.WriteStorage()
		log.InfoStruct("connect ", connect)

	}
	log.InfoStruct(" GetGuild ", c.GetGuild())
	log.Info(c.Ident.User.AvatarURL)
	//for i, level := range c.SyncData.TechLevels {
	//	n := module_types.GetTechFromIndex(i)
	//
	//	if n == "" {
	//		log.Info(fmt.Sprintf("%d %d", i, level.Level))
	//
	//	} else {
	//		log.Info(fmt.Sprintf("%s %d", n, level.Level))
	//	}
	//}
	//data, err := c.CorpData("1202029515556270120")
	//if err != nil {
	//	log.ErrorErr(err)
	//	return
	//}
	//
	//for i, Member := range data.Members {
	//	fmt.Println(i, Member.Name)
	//}
	log.Info(fmt.Sprintf("%+v\n", "ok"))
}
