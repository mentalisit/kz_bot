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
	ReadConfigRs() []models.CorporationConfig
	DeleteConfigRs(c models.CorporationConfig)
	AutoHelpUpdateMesid(c models.CorporationConfig)
	AutoHelp() []models.CorporationConfig
}

func (s *Storage) DeleteConfigRs(c models.CorporationConfig) {
	s.ConfigRs.DeleteConfigRs(c)
	var a map[string]models.CorporationConfig
	a = make(map[string]models.CorporationConfig)
	b := s.ConfigRs.ReadConfigRs()
	for _, config := range b {
		a[config.CorpName] = config
	}
	s.CorpConfigRS = a
}
