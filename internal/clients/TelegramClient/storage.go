package TelegramClient

import "kz_bot/internal/models"

// BridgeCheckChannelConfigTg bridge
func (t *Telegram) BridgeCheckChannelConfigTg(mId string) (bool, models.BridgeConfig) {
	for _, config := range t.bridgeConfig {
		for _, channelD := range config.ChannelTg {
			if channelD.ChannelId == mId {
				return true, config
			}
		}
	}
	return false, models.BridgeConfig{}
}
func (t *Telegram) CheckChannelCompendium(chatId64 int64) (bool, corpCompendium) {
	for _, config := range corp {
		if config.chatid == chatId64 {
			return true, config
		}
	}
	return false, corpCompendium{}
}

// CheckChannelConfigTG RsConfig
func (t *Telegram) CheckChannelConfigTG(chatid string) (channelGood bool, config models.CorporationConfig) {
	for _, corpporationConfig := range t.corpConfigRS {
		if corpporationConfig.TgChannel == chatid {
			return true, corpporationConfig
		}
	}
	return false, models.CorporationConfig{}
}

// AddTgCorpConfig add RsConfig
func (t *Telegram) AddTgCorpConfig(chatName string, chatid, lang string) {
	c := models.CorporationConfig{
		CorpName:  chatName,
		Country:   lang,
		TgChannel: chatid,
	}
	t.storage.ConfigRs.InsertConfigRs(c)
	t.corpConfigRS[c.CorpName] = c
	t.log.Info(chatName + " Добавлена в конфиг корпораций ")
}
