package HadesClient

import (
	"fmt"
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
	go h.reloadConsoleClient()
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

func (h *Hades) CheckMember(member, corporation string, mId int64) {
	for _, aMember := range h.member {
		if aMember.UserName == member && (aMember.CorpName == corporation || aMember.CorpName == "1") {
			a := models.MessageHadesClient{
				Sender:      member,
				ChannelType: 0,
				Corporation: corporation,
				Command:     "access",
				MessageId:   mId,
			}
			h.toGame <- a
			fmt.Printf("%+v", a)
		}
	}
}
func (h *Hades) CheckMemberRang(member, corporation string, playerId int64) {
	for _, aMember := range h.member {
		if aMember.UserName == member && aMember.CorpName == corporation && aMember.Rang != 0 {
			a := models.MessageHadesClient{
				Sender:      member,
				ChannelType: aMember.Rang,
				Corporation: corporation,
				Command:     "rang",
				MessageId:   playerId,
			}
			h.toGame <- a
			fmt.Printf("%+v", a)
		}
	}
}
