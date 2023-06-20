package TelegramClient

import "kz_bot/internal/models"

func (t *Telegram) loadDbHades() {
	corp := t.storage.HadesClient.GetAllCorporationHades()
	for _, client := range corp {
		t.corporationHades[client.Corp] = client
	}
}
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
