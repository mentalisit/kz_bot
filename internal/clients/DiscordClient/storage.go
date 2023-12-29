package DiscordClient

import (
	"kz_bot/internal/models"
)

// BridgeCheckChannelConfigDS bridge
func (d *Discord) BridgeCheckChannelConfigDS(ChatId string) (bool, models.BridgeConfig) {
	for _, config := range d.bridgeConfig {
		for _, channelD := range config.ChannelDs {
			if channelD.ChannelId == ChatId {
				return true, config
			}
		}
	}
	return false, models.BridgeConfig{}
}

// CheckChannelConfigDS RsConfig
func (d *Discord) CheckChannelConfigDS(chatid string) (channelGood bool, config models.CorporationConfig) {
	for _, corpporationConfig := range d.corpConfigRS {
		if corpporationConfig.DsChannel == chatid {
			return true, corpporationConfig
		}
	}
	return false, models.CorporationConfig{}
}

// AddDsCorpConfig add RsConfig
func (d *Discord) AddDsCorpConfig(chatName, chatid, guildid, lang string) {
	c := models.CorporationConfig{
		CorpName:  chatName,
		DsChannel: chatid,
		Country:   lang,
		Guildid:   guildid,
	}
	d.storage.ConfigRs.InsertConfigRs(c)
	d.corpConfigRS[c.CorpName] = c
	d.log.Info(chatName + " Добавлена в конфиг корпораций ")
	d.HelpChannelUpdate(c)
}

// hadesClient
//func (d *Discord) getCorpHadesAlliance(ChatId string) models.CorporationHadesClient {
//	for _, client := range d.corporationHades {
//		if client.DsChat == ChatId {
//			return client
//		}
//	}
//	return models.CorporationHadesClient{}
//}
//func (d *Discord) getCorpHadesWs1(ChatId string) models.CorporationHadesClient {
//	for _, client := range d.corporationHades {
//		if client.DsChatWS1 == ChatId {
//			return client
//		}
//	}
//	return models.CorporationHadesClient{}
//}
