package bot

import (
	"fmt"
	"time"
)

const hhelpText = "В боте используются только кирилица \n" +
	"Встать в очередь: [4-11]+  или\n" +
	" [4-11]+[указать время ожидания в минутах]\n" +
	"(уровень кз)+(время ожидания)\n" +
	" 9+  встать в очередь на КЗ 9ур.\n" +
	" 9+60  встать на КЗ 9ур, время ожидания не более 60 минут.\n" +
	"Покинуть очередь: [4-11] -\n" +
	" 9- выйти из очереди КЗ 9ур.\n" +
	"Посмотреть список активных очередей: о[4-11]\n" +
	" о9 вывод очередь для вашей Кз\n" +
	"Получить роль кз: + [5-11]\n" +
	" +9 получить роль КЗ 9ур.\n" +
	" -9 снять роль "

func (b *Bot) hhelp() {
	b.iftipdelete()
	if b.in.Tip == "ds" {
		m := b.Ds.SendEmbedText(b.in.Config.DsChannel, "Справка",
			fmt.Sprintf("ВНИМАНИЕ БОТ УДАЛЯЕТ СООБЩЕНИЯ \n ОТ ПОЛЬЗОВАТЕЛЕЙ ЧЕРЕЗ 3 МИНУТЫ \n\n"+hhelpText))
		b.Ds.DeleteMesageSecond(b.in.Config.DsChannel, m.ID, 180)
	} else if b.in.Tip == "tg" {
		b.Tg.SendChannelDelSecond(b.in.Config.TgChannel, "Справка\n"+hhelpText+"\n/help", 180)
	}
}

func (b *Bot) autohelp() {
	tm := time.Now()
	mtime := tm.Format("15:04")
	if mtime == "12:00" {
		a := b.Db.AutoHelp()
		for _, s := range a {
			if s.Config.MesidDsHelp != "" {
				go b.Ds.DeleteMessage(s.DsChannel, s.Config.MesidDsHelp)
				b.helpChannelUpdate(s.DsChannel)
			} else {
				b.helpChannelUpdate(s.DsChannel)
			}
		}
	}
}

func (b *Bot) helpChannelUpdate(dschannel string) {
	newMesidHelp := b.hhelp1(dschannel)
	b.Db.AutoHelpUpdateMesid(newMesidHelp, dschannel)

}
func (b *Bot) hhelp1(chatid string) string {
	mes := b.Ds.Send(chatid, fmt.Sprintf("Справка \n"+
		"ВНИМАНИЕ БОТ УДАЛЯЕТ СООБЩЕНИЯ \n ОТ ПОЛЬЗОВАТЕЛЕЙ ЧЕРЕЗ 3 МИНУТЫ \n\n"+
		hhelpText))
	return mes
}
