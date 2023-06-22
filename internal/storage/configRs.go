package storage

import (
	"kz_bot/internal/models"
)

type Caache interface {
	ReloadConfig()
	//AddCorp(CorpName string, DsChannel string, TgChannel int64, WaChannel string, DelMesComplite int, mesiddshelp string, Country string, guildid string)

	CheckChannelConfigWA(chatid string) (channelGood bool, config models.CorporationConfig)
	CheckChannelConfigTG(chatid int64) (channelGood bool, config models.CorporationConfig)
	CheckCorpNameConfig(corpname string) (channelGood bool, config models.CorporationConfig)
	ReadAllChannel() (chatDS []string, chatTG []int64, chatWA []string)
}
type ConfigRs interface {
	InsertConfigRs(c models.CorporationConfig)
}
