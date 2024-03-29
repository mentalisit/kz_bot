package words

//

func (w *Words) setWordsRu() {
	w.ru["1"] = "один"
	w.ru["2"] = "два"
	w.ru["tiUjeVocheredi"] = " ты уже в очереди"
	w.ru["zapustilOchered"] = " запустил очередь "
	w.ru["ocheredKz"] = "Очередь кз"
	w.ru["min."] = "мин."
	w.ru["prinuditelniStart"] = "принудительный старт"
	w.ru["SborNaKz"] = "Сбор на кз"
	w.ru["kz"] = "кз"
	w.ru["tiUjePodpisanNaKz"] = "ты уже подписан на кз"
	w.ru["dlyaDobavleniyaVochered"] = "для добавления в очередь напиши"
	w.ru["viPodpisalisNaPing"] = "вы подписались на пинг кз"
	w.ru["tiNePodpisanNaPingKz"] = "ты не подписан на пинг кз"
	w.ru["otpisalsyaOtPingaKz"] = "отписался от пинга кз"
	w.ru["prisoedenilsyKocheredi"] = " присоединился к очереди"
	w.ru["nujenEsheOdinDlyFulki"] = "нужен еще один для фулки"
	w.ru["sformirovana"] = "сформирована"
	w.ru["Vigru"] = "В ИГРУ"
	w.ru["zakrilOcheredKz"] = " закрыл очередь кз"
	w.ru["tiNeVOcheredi"] = " ты не в очереди"
	w.ru["pokinulOchered"] = " покинул очередь"
	w.ru["bilaUdalena"] = "была удалена"
	w.ru["dly ustanovki"] = "	Для установки эмоджи пиши текст \nЭмоджи пробел (номер ячейки1-4) пробел эмоджи \n пример \nЭмоджи 1 🚀\n	Ваши слоты"
	w.ru["vashiEmodji"] = "Ваши эмоджи\n"
	w.ru["dly iventa"] = "для ивента"
	w.ru["iventZapushen"] = "Ивент запущен. После каждого похода на КЗ, один из участников КЗ вносит полученные очки в базу командой К (номер катки) (количество набраных очков)"
	w.ru["rejimIventaUje"] = "Режим ивента уже активирован."
	w.ru["zapuskIostanovka"] = "Запуск | Оcтановка Ивента доступен Администратору канала."
	w.ru["IventOstanovlen"] = "Ивент остановлен."
	w.ru["iventItakAktiven"] = "Ивент и так не активен. Нечего останавливать "
	w.ru["dannieKzUjeVneseni"] = "данные о кз уже внесены "
	w.ru["ochki vnesen"] = "Очки внесены в базу"
	w.ru["dobavlenieOchkovNevozmojno"] = "добавление очков невозможно. Вы не являетесь участником КЗ под номером"
	w.ru["iventNeZapushen"] = "Ивент не запущен."
	w.ru["iventIgra"] = "Ивент игра"
	w.ru["vneseno"] = "Внесено"
	w.ru["VremyaPochtiVishlo"] = " время почти вышло...\nДля продления времени ожидания на 30м жми +\nДля выхода из очереди жми -"
	w.ru["ranovatoPlysik"] = "рановато плюсик жмешь, ты в очереди на кз"
	w.ru["budeshEshe"] = "будешь еще"
	w.ru["vremyaObnovleno"] = " время обновлено "
	w.ru["ranovatoMinus"] = "рановато минус жмешь, ты в очереди на кз"
	w.ru["pusta"] = " пуста " //очередь кз пуста
	w.ru["netAktivnuh"] = "Нет активных очередей "
	w.ru["prinuditelniStartDostupen"] = "Принудительный старт доступен участникам очереди."
	w.ru["bilaZapushenaNe"] = "была запущена не полной"
	w.ru["maksimalnoeVremya"] = "максимальное время в очереди ограничено на 180 минут\n твое время"
	w.ru["vremyaObnovleno"] = " время обновлено +30"
	w.ru["ScanDB"] = "Сканирую базу данных"
	w.ru["noHistory"] = " История не найдена "
	w.ru["formlist"] = "Формирую список "
	w.ru["topUchastnikov"] = "ТОП Участников"
	w.ru["iventa"] = "ивента"
	w.ru["teperViPodpisani"] = "Теперь вы подписаны на"
	w.ru["ViUjePodpisan"] = "Вы уже подписаны на"
	w.ru["oshibkaNedostatochno"] = "ошибка: недостаточно прав для выдачи роли "
	w.ru["viNePodpisani"] = "Вы не подписаны на роль"
	w.ru["netTakoiRoli"] = "нет такой роли на сервере"
	w.ru["ViOtpisalis"] = "Вы отписались от роли"
	w.ru["OshibkaNedostatochnadlyaS"] = "ошибка: недостаточно прав для снятия роли  "
	w.ru["jelaushieNa"] = "Желающие на"
	w.ru["DlyaDobavleniya"] = "для добавления в очередь"
	w.ru["DlyaVihodaIz"] = "для выхода из очереди"
	w.ru["DannieObnovleni"] = "Данные обновлены"
	w.ru["hhelpText"] = "Стать в очередь: [4-11]+  или\n " +
		"[4-11]+[указать время ожидания в минутах]\n" +
		"(уровень кз)+(время ожидания)\n" +
		" 9+  встать в очередь на КЗ 9ур.\n" +
		" 9+60  встать на КЗ 9ур, время ожидания не более 60 минут.\n" +
		"Покинуть очередь: [4-11] -\n 9- выйти из очереди КЗ 9ур.\n" +
		"Посмотреть список активных очередей: о[4-11]\n" +
		" о9 вывод очередь для вашей Кз\n" +
		"Получить роль кз: + [5-11]\n +9 получить роль КЗ 9ур.\n -9 снять роль \n" +
		"Для Тёмных красных звезд\n Для старта очереди\n9*\nДля получения роли \n+d9"
	w.ru["spravka"] = "Справка"
	w.ru["botUdalyaet"] = "ВНИМАНИЕ БОТ УДАЛЯЕТ СООБЩЕНИЯ \n ОТ ПОЛЬЗОВАТЕЛЕЙ ЧЕРЕЗ 3 МИНУТЫ"
	w.ru["accessAlready"] = "Я уже могу работать на вашем канале\nповторная активация не требуется.\nнапиши Справка"
	w.ru["accessTY"] = "Спасибо за активацию."
	w.ru["accessYourChannel"] = "ваш канал и так не подключен к логике бота "
	w.ru["YouDisabledMyFeatures"] = "вы отключили мои возможности"
	w.ru["dkz"] = "ткз"
	w.ru["ocheredTKz"] = "Очередь ткз"
	w.ru["zakrilOcheredTKz"] = " закрыл очередь ткз"
	w.ru["vashLanguage"] = "Вы переключили меня на Русский язык"

}
