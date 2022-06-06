package telegramClient

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"time"
)

//команда хелп
func (t Telegram) help(chatid int64, mesid int) {
	go func() {
		time.Sleep(10 * time.Second)
		t.t.Request(tgbotapi.DeleteMessageConfig(tgbotapi.NewDeleteMessage(chatid, mesid)))
	}()
	text := "Команды:\n" +
		"Очередь КЗ - /helpqueue\n" +
		"Уведомления - /helpnotification\n" +
		"Событие КЗ - /helpevent\n" +
		"ТОП лист - /helptop\n" +
		"Работа с иконками - /helpicon"
	t.SendChannelDelSecond(chatid, text, 60)

}

//очередь кз
func (t Telegram) helpQueue(chatid int64, mesid int) {
	go t.DelMessageSecond(chatid, mesid, 10)
	text := "Очередь на КЗ:\n" +
		"Встать в очередь:\n" +
		"[4-11]+ / [4-11]+[указать время ожидания в минутах]\n" +
		"9+ - встать в очередь на КЗ 9ур.\n" +
		"9+60 - встать в очередь на КЗ 9ур, время ожидания не более 60 минут.\n\n" +
		"Покинуть очередь:\n" +
		"[4-11]-\n" +
		"9-  -выйти из очереди на КЗ 9ур.\n\n" +

		"вывод очереди  о[x]  где х уровень нужной кз\n" +
		"о9  -вывод очереди для кз9"
	t.SendChannelDelSecond(chatid, text, 60)
}

//Уведомления
func (t Telegram) helpNotification(chatid int64, mesid int) {
	go t.DelMessageSecond(chatid, mesid, 10)
	text := "Уведомления:\n" +
		"	Подписаться на уведомления о начале очереди: +[4-11]\n" +
		"+10 -подписаться на уведомления о начале очереди на КЗ 10ур.\n\n" +

		"	Подписаться на уведомление, если в очереди 3 человека: ++[4-11]\n" +
		"++10 -подписаться на уведомления о наличии 3х человек в очереди на КЗ 10ур.\n\n" +

		"	Отключить уведомления о начале сбора: -[5-11]\n" +
		"-9 -отключить уведомления о начале сборе на КЗ 9ур.\n\n" +

		"	Отключить уведомления 3/4 в очереди: --[5-11]\n" +
		"--9 -отключить уведомления о наличии 3х человек в очереди на КЗ 9ур."
	t.SendChannelDelSecond(chatid, text, 60)
}

//Ивент кз
func (t Telegram) helpEvent(chatid int64, mesid int) {
	go t.DelMessageSecond(chatid, mesid, 10)
	text := "Режим Событий для КЗ:\n" +
		"Включить режим ( для админов ) - **Ивент старт** ,\n" +
		"Выключить режим ( для админов ) - Ивент стоп .\n\n" +
		//"Во время активного События доступна команда х+сам, где х - уровень КЗ.\n\n" +
		"Внести очки в базу - К (номер катки) (количество набраных очков)\n\n"
	//"Посмотреть все завершенные События - <b>rs event</b>\n" +
	//"Посмотреть ТОП по прошлым События - rs top event n, где n-номер События.\n" +
	//"Посмотреть ТОП по определенным КЗ, определенного События - <b>rs top x event n</b>, где <b>x</b> - уровень КЗ, <b>n</b> - номер События.\n"
	t.SendChannelDelSecond(chatid, text, 60)
}

//Топ лист
func (t Telegram) helpTop(chatid int64, mesid int) {
	go t.DelMessageSecond(chatid, mesid, 10)
	text := "ТОП - лист:\n\n" +
		"Топ - вывод ТОП-листа за все время ( в период активного Ивента - вывод ТОП-листа с начала старта),\n\n" +
		//"Топ д - вывод ТОП-листа за текущие сутки,\n\n"+
		//"Топ н - вывод ТОП-листа за последние 7 дней,\n\n"+
		"Топ [x] - вывод ТОП-листа на КЗ определенного уровня, где x - уровень КЗ." +
		"Пример \nТоп 9"
	t.SendChannelDelSecond(chatid, text, 60)
}

//Работа с иконками
func (t Telegram) helpIcon(chatid int64, mesid int) {
	go t.DelMessageSecond(chatid, mesid, 10)
	text := "Работа с иконками:\n" +
		"Функционал БОТа позволяет рядом с ником в очереди на КЗ, установить иконки ( пользователю долступно два слота ). Допустимы только html-коды иконок. Пример можно взять тут - https://unicode-table.com/ru/emoji/\n" +
		"Эмоджи [x] [y] - установка иконки в x -слот с кодом y\n" +
		"Эмоджи [x]- удаление иконки из слота x"
	t.SendChannelDelSecond(chatid, text, 60)
}