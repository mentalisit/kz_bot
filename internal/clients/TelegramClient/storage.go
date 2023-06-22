package TelegramClient

import "kz_bot/internal/models"

// BridgeCheckChannelConfigTg bridge
func (t *Telegram) BridgeCheckChannelConfigTg(mId int64) (bool, models.BridgeConfig) {
	for _, config := range t.bridgeConfig {
		for _, channelD := range config.ChannelTg {
			if channelD.ChannelId == mId {
				return true, config
			}
		}
	}
	return false, models.BridgeConfig{}
}

// CheckChannelConfigTG RsConfig
func (t *Telegram) CheckChannelConfigTG(chatid int64) (channelGood bool, config models.CorporationConfig) {
	for _, corpporationConfig := range t.corpConfigRS {
		if corpporationConfig.TgChannel == chatid {
			return true, corpporationConfig
		}
	}
	return false, models.CorporationConfig{}
}

// AddDsCorpConfig add RsConfig
func (t *Telegram) AddDsCorpConfig(chatName string, chatid int64) {
	c := models.CorporationConfig{
		CorpName:  chatName,
		Country:   "ru",
		TgChannel: chatid,
	}
	t.storage.ConfigRs.InsertConfigRs(c)
	t.corpConfigRS[c.CorpName] = c
	t.log.Println(chatName, "Добавлена в конфиг корпораций ")
}

// hadesClient
func (t *Telegram) getCorpHadesAlliance(ChatId int64) models.CorporationHadesClient {
	for _, client := range t.corporationHades {
		if client.TgChat == ChatId {
			return client
		}
	}
	return models.CorporationHadesClient{}
}
func (t *Telegram) getCorpHadesWs1(ChatId int64) models.CorporationHadesClient {
	for _, client := range t.corporationHades {
		if client.TgChatWS1 == ChatId {
			return client
		}
	}
	return models.CorporationHadesClient{}
}
