package corpsConfig

import (
	"kz_bot/internal/models"
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
	AddCorp(CorpName string, DsChannel string, TgChannel int64, WaChannel string, DelMesComplite int, mesiddshelp string, mesidtghelp int, guildid string)
	ReloadConfig()
}

func (c CorpConfig) ReloadConfig() {
	*P = *New()
}
func (c CorpConfig) AddCorp(CorpName string, DsChannel string, TgChannel int64, WaChannel string, DelMesComplite int, mesiddshelp string, mesidtghelp int, guildid string) {
	corpConfig := models.BotConfig{
		Type:      0xff,
		CorpName:  CorpName,
		DsChannel: DsChannel,
		TgChannel: TgChannel,
		WaChannel: WaChannel,
		Config: models.Configs{
			DelMesComplite: DelMesComplite,
			MesidDsHelp:    mesiddshelp,
			MesidTgHelp:    mesidtghelp,
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
