package dbase

import "kz_bot/internal/models"

type CorpConfig interface {
	ReadBotCorpConfig()                               //Чтение из бд конфигураций корпораций при запуске бота
	DeleteTgChannel(chatid int64)                     //отключение бота от чата в телеграм
	DeleteDsChannel(chatid string)                    //отключение бота от чата в дискорд
	AddTgCorpConfig(chatName string, chatid int64)    //добавление чата телеграм в конфиг корпораций
	AddDsCorpConfig(chatName, chatid, guildid string) //добавление чата дискорд в конфиг корпораций
}
type Top interface {
	TopTemp() string //временный топ
	TopTempEvent() string
	TopAll(CorpName string) bool
	TopAllEvent(CorpName string, numberevent int) bool
	TopLevel(CorpName, lvlkz string) bool                    //топ по уровню
	TopEventLevel(CorpName, lvlkz string, numEvent int) bool //топ по уровню во время ивента
	TopAllDay(CorpName string, oldDate string) bool
	TopLevelDay(CorpName, lvlkz string, oldDate string) bool
}
type Event interface {
	NumActiveEvent(CorpName string) (event1 int)    //номер активного ивента
	NumDeactivEvent(CorpName string) (event0 int)   //номер предыдущего ивента
	UpdateActiveEvent0(CorpName string, event1 int) //отключение активного ивента
	EventStartInsert(CorpName string)               //включение ивента
	CountEventNames(CorpName, name string, numberkz, numEvent int) (countEventNames int)
	CountEventsPoints(CorpName string, numberkz, numberEvent int) int
	UpdatePoints(CorpName string, numberkz, points, event1 int) int
	ReadNamesMessage(CorpName string, numberkz, numberEvent int) (nd, nt models.Names, t models.Sborkz)
}
type Subscribe interface {
	CheckSubscribe(name, lvlkz string, TgChannel int64, tipPing int) int                //проверка активной подписки
	Subscribe(name, nameMention, lvlkz string, tipPing int, TgChannel int64)            //подписка
	Unsubscribe(name, lvlkz string, TgChannel int64, tipPing int)                       //отписка
	SubscPing(nameMention, lvlkz, CorpName string, tipPing int, TgChannel int64) string //чтение для пинга игроков в телеграм
}
type Emoji interface {
	EmReadUsers(name, tip string) models.EmodjiUser    //чтение эмоджи игрока с бд
	EmUpdateEmodji(name, tip, slot, emo string) string //обновление эмоджи игрока
	EmInsertEmpty(tip, name string)                    // внесение имени для эмоджи
}
type DbInterface interface {
	CorpConfig
	Top
	Event
	Subscribe
	Emoji
	СountName(name, lvlkz, corpName string) int                                                                                 //проверка состоит ли игрок уже в очереди
	CountQueue(lvlkz, CorpName string) int                                                                                      //проверка сколько игроков в очереди
	CountNumberNameActive1(lvlkz, CorpName, name string) int                                                                    //проверка количество выполненых игр
	NumberQueueLvl(lvlkz, CorpName string) int                                                                                  //Номер катки по уровню
	ReadAll(lvlkz, CorpName string) (users models.Users)                                                                        //чтение игроков в очереди
	InsertQueue(dsmesid, wamesid, CorpName, name, nameMention, tip, lvlkz, timekz string, tgmesid, numkzN int)                  //внесение данных сбора
	MesidTgUpdate(mesidtg int, lvlkz string, corpname string)                                                                   //изменение ид сообщения в бд
	MesidDsUpdate(mesidds, lvlkz, corpname string)                                                                              //изменение ид сообщения в бд
	UpdateCompliteRS(lvlkz string, dsmesid string, tgmesid int, wamesid string, numberkz int, numberevent int, corpname string) //закрытие очереди кз
	CountNameQueue(name string) (countNames int)                                                                                //проверка игрока на наличие в очереди
	ElseTrue(name string) models.Sborkz                                                                                         //для выхода из очереди при другом старте
	DeleteQueue(name, lvlkz, CorpName string)                                                                                   //Если игрок покидает очередь
	UpdateMitutsQueue(name, CorpName string) models.Sborkz                                                                      //проверка хочет ли игрок продолжить время в очереди
	TimerInsert(dsmesid, dschatid string, tgmesid int, tgchatid int64, timed int)                                               //внесение ид сообщения в бд
	TimerDeleteMessage() []models.Timer                                                                                         //удаление из таймера
	P30Pl(lvlkz, CorpName, name string) int                                                                                     //+30 минут если в очереди
	UpdateTimedown(lvlkz, CorpName, name string)                                                                                //при нажатии плюса при остатке меньше 3х минут
	ReadMesIdDS(mesid string) (string, error)
	Queue(corpname string) []string
	AutoHelp() []models.BotConfig
	AutoHelpUpdateMesid(newMesidHelp, dschannel string)
	MinusMin() []models.Sborkz
	OneMinutsTimer() []string
	MessageUpdateMin(corpname string) ([]string, []int, []string)
	MessageupdateDS(dsmesid string, config models.BotConfig) models.InMessage
	MessageupdateTG(tgmesid int, config models.BotConfig) models.InMessage
}
