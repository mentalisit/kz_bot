package storage

import "kz_bot/internall/storage/memory"

type Cache interface {
	ReloadConfig()
	AddCorp(CorpName string, DsChannel string, TgChannel int64, WaChannel string, DelMesComplite int, mesiddshelp string, Country string, guildid string)
	CheckChannelConfigDS(chatid string) (channelGood bool, config memory.CorpporationConfig)
	CheckChannelConfigWA(chatid string) (channelGood bool, config memory.CorpporationConfig)
	CheckChannelConfigTG(chatid int64) (channelGood bool, config memory.CorpporationConfig)
	CheckCorpNameConfig(corpname string) (channelGood bool, config memory.CorpporationConfig)
	ReadAllChannel() (chatDS []string, chatTG []int64, chatWA []string)
}
