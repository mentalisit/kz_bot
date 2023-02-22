package TelegramClient

import "kz_bot/internal/models"

func (t *Telegram) Help(Channel int64) {
	t.SendChannelDelSecond(Channel, "Справка\n"+models.HhelpText+"\n/help", 180)
}

// команда хелп
func (t *Telegram) help(chatid int64, mesid int) {
	t.DelMessageSecond(chatid, mesid, 10)
	t.SendChannelDelSecond(chatid, models.Help, 60)
}

// очередь кз
func (t *Telegram) helpQueue(chatid int64, mesid int) {
	go t.DelMessageSecond(chatid, mesid, 10)
	t.SendChannelDelSecond(chatid, models.HelpQueue, 60)
}

// Уведомления
func (t *Telegram) helpNotification(chatid int64, mesid int) {
	go t.DelMessageSecond(chatid, mesid, 10)
	t.SendChannelDelSecond(chatid, models.HelpNotification, 60)
}

// Ивент кз
func (t *Telegram) helpEvent(chatid int64, mesid int) {
	go t.DelMessageSecond(chatid, mesid, 10)
	t.SendChannelDelSecond(chatid, models.HelpEvent, 60)
}

// Топ лист
func (t *Telegram) helpTop(chatid int64, mesid int) {
	go t.DelMessageSecond(chatid, mesid, 10)

	t.SendChannelDelSecond(chatid, models.HelpTop, 60)
}

// Работа с иконками
func (t *Telegram) helpIcon(chatid int64, mesid int) {
	go t.DelMessageSecond(chatid, mesid, 10)

	t.SendChannelDelSecond(chatid, models.HelpIcon, 60)
}
