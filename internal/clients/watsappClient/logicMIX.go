package watsappClient

import (
	"kz_bot/internal/models"
)

func (w *Watsapp) LogicMIXwa(text, name, nameid, chatid, mesid string) {
	ok, config := w.CorpConfig.CheckChannelConfigWA(chatid)
	w.AccesChatWA(text, chatid)
	if ok {
		in := models.InMessage{
			Mtext:       text,
			Tip:         "wa",
			Name:        name,
			NameMention: "name",
			Wa: struct {
				Nameid string
				Mesid  string
			}{
				Nameid: nameid,
				Mesid:  mesid},
			Config: config,
			Option: models.Option{InClient: true},
		}
		models.ChWa <- in
	}
}
