package watsappClient

import (
	"kz_bot/internal/models"
)

func (w *Watsapp) LogicMIXwa(text, name, nameid, chatid string) {
	ok, config := w.CorpConfig.CheckChannelConfigWA(chatid)
	w.AccesChatWA(text, chatid)
	if ok {
		in := models.InMessage{
			Mtext:       text,
			Tip:         "wa",
			Name:        name,
			NameMention: name,
			Wa: struct {
				Nameid string
			}{
				Nameid: nameid},
			Config: config,
			Option: struct {
				Callback bool
				Edit     bool
				Update   bool
				Queue    bool
			}{},
		}
		models.ChWa <- in
	}

}
