package storage

import (
	"kz_bot/internal/storage/memory"
)

type Cache interface {
	ReloadConfig()
	AddCorp(CorpName string, DsChannel string, TgChannel int64, WaChannel string, DelMesComplite int, mesiddshelp string, Country string, guildid string)
	CheckChannelConfigDS(chatid string) (channelGood bool, config memory.CorpporationConfig)
	CheckChannelConfigWA(chatid string) (channelGood bool, config memory.CorpporationConfig)
	CheckChannelConfigTG(chatid int64) (channelGood bool, config memory.CorpporationConfig)
	CheckCorpNameConfig(corpname string) (channelGood bool, config memory.CorpporationConfig)
	ReadAllChannel() (chatDS []string, chatTG []int64, chatWA []string)
}
type CacheGlobal interface {
	AddCorp(CorpName string, DsChannel string, TgChannel int64, WaChannel string, Country string, guildid string)
	CheckChannelConfigDS(chatid string) (channelGood bool, config memory.ConfigGlobal)
	CheckChannelConfigTG(chatid int64) (channelGood bool, config memory.ConfigGlobal)
	CheckChannelConfigWA(chatid string) (channelGood bool, config memory.ConfigGlobal)
}
