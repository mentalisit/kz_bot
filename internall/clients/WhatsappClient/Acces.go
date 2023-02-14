package WhatsappClient

import (
	"context"
	"fmt"
	"strings"
	"time"
)

func (w *Whatsapp) accesChatWA(text, chatid string) {
	res := strings.HasPrefix(text, ".")
	if res == true && text == ".add" {
		w.accessAddChannelWA(chatid)
	} else if res == true && text == ".del" {
		w.accessDelChannelWa(chatid)
	}
}

func (w *Whatsapp) accessAddChannelWA(chatid string) { // внесение в дб и добавление в масив
	ok, _ := w.storage.Cache.CheckChannelConfigWA(chatid)
	if ok {
		go func() {
			mesId := w.SendText(chatid, "Я уже могу работать на вашем канале\n"+
				"повторная активация не требуется.\nнапиши Справка")

			w.DeleteMessageSecond(chatid, mesId, 30)
		}()
	} else {
		g, chatName := w.getGroupName(chatid)
		if g {
			fmt.Println("новая активация корпорации ", chatName)
			ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
			defer cancel()
			err := w.storage.CorpsConfig.AddWaCorpConfig(ctx, chatName, chatid)
			if err != nil {
				w.log.Println(err)
			}
			w.SendText(chatid, "Спасибо за активацию. Если что пиши Справка")
		}

	}
}
func (w *Whatsapp) accessDelChannelWa(chatid string) { //удаление с бд и масива для блокировки
	ok, config := w.storage.Cache.CheckChannelConfigWA(chatid)
	if !ok {
		w.SendText(chatid, "ваш канал и так не подключен к логике бота ")
	} else {
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
		err := w.storage.CorpsConfig.DeleteWa(ctx, chatid)
		if err != nil {
			w.log.Println(err)
		}
		fmt.Println("отключение корпорации ", config.CorpName)
		w.storage.Cache.ReloadConfig()
		w.storage.CorpsConfig.ReadCorps()
		go w.SendText(chatid, "вы отключили мои возможности")
	}
}
