package HadesClient

import (
	"kz_bot/internal/models"
)

func (h *Hades) loadDB() {
	corporation := h.storage.HadesClient.GetAllCorporationHades()
	for _, client := range corporation {
		h.corporation[client.Corp] = client
	}

	mID := h.storage.HadesClient.GetAllGameMesId()
	for _, id := range mID {
		h.idMessage[id.CorpName] = id.MessageId
	}

	mIDw := h.storage.HadesClient.GetAllGameWs1MesId()
	for _, id := range mIDw {
		h.idMessageWs1[id.StarId] = id.MessageId
	}
	member := h.storage.HadesClient.GetAllMember()
	for _, allianceMember := range member {
		h.member[allianceMember.UserName] = allianceMember
	}
}

func (h *Hades) getChatIdAlliance() (mId int64) {
	mId = h.idMessage[h.in.Corporation]
	if mId == 0 {
		mId = h.storage.HadesClient.GetCorpMesId(h.in.Corporation)
	}
	return mId
}
func (h *Hades) getConfig(Corporation string) (corp models.CorporationHadesClient) {
	if h.corporation[Corporation].Corp != "" {
		corp = h.corporation[Corporation]
	} else {
		if Corporation != "" {
			corp = h.storage.HadesClient.GetCorporation(Corporation)
		}
	}
	return corp
}
