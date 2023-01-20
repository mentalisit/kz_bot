package discordClient

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

func (d *Discord) Help(Channel string) {
	m := d.SendEmbedText(Channel, "Справка",
		fmt.Sprintf("ВНИМАНИЕ БОТ УДАЛЯЕТ СООБЩЕНИЯ \n ОТ ПОЛЬЗОВАТЕЛЕЙ ЧЕРЕЗ 3 МИНУТЫ \n\n"+hhelpText))
	d.DeleteMesageSecond(Channel, m.ID, 180)
}

func (d *Discord) Autohelpds() {
	tm := time.Now()
	mtime := tm.Format("15:04")
	if mtime == "12:00" {
		a := d.dbase.CorpConfig.AutoHelp()
		for _, s := range a {
			if s.DsChannel != "" {
				if s.Config.MesidDsHelp != "" {
					go d.DeleteMessage(s.DsChannel, s.Config.MesidDsHelp)
					d.HelpChannelUpdate(s.DsChannel)
				} else {
					d.HelpChannelUpdate(s.DsChannel)
				}
			}

		}
		time.Sleep(time.Minute)
	}
}

func (d *Discord) HelpChannelUpdate(dschannel string) {
	newMesidHelp := d.hhelp1(dschannel)
	d.dbase.CorpConfig.AutoHelpUpdateMesid(newMesidHelp, dschannel)
}
func (d *Discord) hhelp1(chatid string) string {
	mes := d.SendEmbedText(chatid, "Справка", fmt.Sprintf(" \n"+
		"ВНИМАНИЕ БОТ УДАЛЯЕТ СООБЩЕНИЯ \n ОТ ПОЛЬЗОВАТЕЛЕЙ ЧЕРЕЗ 3 МИНУТЫ \n\n"+
		hhelpText))
	return mes.ID
}
