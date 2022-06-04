package corpsConfig

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"kz_bot/internal/models"
	"strings"
)

var P = New()

func New() *Proxies {
	var arr Proxies
	return &arr
}

type Proxies []models.BotConfig
type CorpConfig struct{}
type ConfigCorp interface {
	CheckCorpNameConfig(corpname string) (channelGood bool, config models.BotConfig)
	CheckChannelConfigDS(chatid string) (channelGood bool, config models.BotConfig)
	CheckChannelConfigTG(chatid int64) (channelGood bool, config models.BotConfig)
}

func addCorp(CorpName string, DsChannel string, TgChannel int64, WaChannel string, DelMesComplite int, mesiddshelp string, mesidtghelp int, guildid string) {
	corpConfig := BotConfig{
		Type:      0xff,
		CorpName:  CorpName,
		DsChannel: DsChannel,
		TgChannel: TgChannel,
		WaChannel: WaChannel,
		Config: Configs{
			DelMesComplite: DelMesComplite,
			mesidDsHelp:    mesiddshelp,
			mesidTgHelp:    mesidtghelp,
			Primer:         "",
			Guildid:        guildid,
		},
	}
	*P = append(*P, corpConfig)
}

func (c CorpConfig) CheckChannelConfigDS(chatid string) (channelGood bool, config models.BotConfig) {
	if chatid != "" {
		for _, pp := range *P {
			if chatid == pp.DsChannel {
				channelGood = true
				config = pp
				break
			}
		}
	}
	return channelGood, config
}
func (c CorpConfig) CheckChannelConfigTG(chatid int64) (channelGood bool, config models.BotConfig) {
	if chatid != 0 {
		for _, pp := range *P {
			if chatid == pp.TgChannel {
				channelGood = true
				config = pp
				break
			}
		}
	}
	return channelGood, config
}
func (c CorpConfig) CheckCorpNameConfig(corpname string) (channelGood bool, config models.BotConfig) {
	if corpname != "" { // есть ли корпа
		for _, pp := range *P {
			if corpname == pp.CorpName {
				channelGood = true
				config = pp
				break
			}
		}
	}
	return channelGood, config
}

func readBotConfig() { // чтение с бд и добавление в масив
	db, err := database.DbConnection()
	if err != nil {
		logrus.Println(err)
	}
	results, err := db.Query("SELECT * FROM config")
	if err != nil {
		logrus.Println(err)
	}
	var t TableConfig
	for results.Next() {
		err = results.Scan(&t.id, &t.corpname, &t.dschannel, &t.tgchannel, &t.wachannel, &t.mesiddshelp, &t.mesidtghelp, &t.delmescomplite, &t.guildid)
		addCorp(t.corpname, t.dschannel, t.tgchannel, t.wachannel, t.delmescomplite, t.mesiddshelp, t.mesidtghelp, t.guildid)
	}
}
