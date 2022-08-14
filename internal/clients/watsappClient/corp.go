package watsappClient

import (
	"fmt"
	"strings"
)

func (w *Watsapp) AccesChatWA(text, chatid string) {
	res := strings.HasPrefix(text, ".")
	if res == true && text == ".add" {
		w.accessAddChannelWA(chatid)
	} else if res == true && text == ".del" {
		w.accessDelChannelWa(chatid)
	}
}

func (w *Watsapp) accessAddChannelWA(chatid string) { // внесение в дб и добавление в масив
	ok, _ := w.CorpConfig.CheckChannelConfigWA(chatid)
	if ok {
		go func() {
			send, err := w.Send(chatid, "Я уже могу работать на вашем канале\n"+
				"повторная активация не требуется.\nнапиши Справка")
			if err != nil {
				return
			}
			w.DeleteMessage(chatid, send)
		}()
	} else {
		chatName := w.ChatName(chatid)
		fmt.Println("новая активация корпорации ", chatName)
		w.dbase.CorpConfig.AddWaCorpConfig(chatName, chatid)
		w.Send(chatid, "Спасибо за активацию. Если что пиши Справка")
	}
}
func (w *Watsapp) accessDelChannelWa(chatid string) { //удаление с бд и масива для блокировки
	ok, config := w.CorpConfig.CheckChannelConfigWA(chatid)
	if !ok {
		w.Send(chatid, "ваш канал и так не подключен к логике бота ")
	} else {
		w.dbase.CorpConfig.DeleteWaChannel(chatid)
		fmt.Println("отключение корпорации ", config.CorpName)
		w.CorpConfig.ReloadConfig()
		w.dbase.CorpConfig.ReadBotCorpConfig()
		go w.Send(chatid, "вы отключили мои возможности")
	}
}
