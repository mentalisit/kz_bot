package hades

import "kz_bot/internal/models"

var corps []models.Corporation

func (s *Hades) AllianceName(name string) (bool, models.Corporation) {
	for i := range corps {
		if corps[i].Corp == name {
			return true, corps[i]
		}
	}
	return false, models.Corporation{}
}
func (s *Hades) AllianceChat(chatId string) (bool, models.Corporation) {
	for i := range corps {
		if corps[i].DsChat == chatId {
			return true, corps[i]
		}
	}
	return false, models.Corporation{}
}
func (s *Hades) Ws1Chat(chatId string) (bool, models.Corporation) {
	for i := range corps {
		if corps[i].DsChatWS1 == chatId {
			return true, corps[i]
		}
	}
	return false, models.Corporation{}
}
func (s *Hades) Ws2Chat(chatId string) (bool, models.Corporation) {
	for i := range corps {
		if corps[i].DsChatWS2 == chatId {
			return true, corps[i]
		}
	}
	return false, models.Corporation{}
}
func (s *Hades) AllianceChatTg(chatId int64) (bool, models.Corporation) {
	for i := range corps {
		if corps[i].TgChat == chatId {
			return true, corps[i]
		}
	}
	return false, models.Corporation{}
}
func (s *Hades) Ws1ChatTg(chatId int64) (bool, models.Corporation) {
	for i := range corps {
		if corps[i].TgChatWS1 == chatId {
			return true, corps[i]
		}
	}
	return false, models.Corporation{}
}
