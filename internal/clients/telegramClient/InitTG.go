package telegramClient

import (
	"database/sql"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
	corpsConfig "kz_bot/internal/clients/corpConfig"
	"kz_bot/internal/dbase/dbaseMysql"
)

type Telegram struct {
	t tgbotapi.BotAPI
	corpsConfig.CorpConfig
	dbase dbaseMysql.Db
	log   *logrus.Logger
}

func (t *Telegram) InitTG(tokent string, db *sql.DB, log *logrus.Logger) {
	t.dbase.Db = db
	t.log = log
	//подключение к телеграм
	TgBot, Err := tgbotapi.NewBotAPI(tokent)
	if Err != nil {
		t.log.Panic("ошибка подключения к телеграм ", Err)
	}
	TgBot.Debug = false
	t.log.Printf("Бот TELEGRAM загружен  %s\n", TgBot.Self.UserName)
	ut := tgbotapi.NewUpdate(0)
	ut.Timeout = 60
	go func() { //получаем обновления от телеграм
		updates := TgBot.GetUpdatesChan(ut)
		for update := range updates {
			if update.CallbackQuery != nil {
				t.callback(update.CallbackQuery) //нажатия в чате
			} else if update.Message != nil {
				if update.Message.Chat.IsPrivate() { //если пишут боту в личку
					t.SendChannel(update.Message.Chat.ID, "сорян это в разработке \n"+
						"я еще не решил как тут сделать"+
						"Присылай идеи для работы с ботом мне @mentalisit ")
				} else if update.Message.IsCommand() {
					t.updatesComand(update.Message) //если сообщение является командой

				} else { //остальные сообщения
					t.logicMixTelegram(update.Message)
				}

			} else if update.MyChatMember != nil {
				t.myChatMember(update.MyChatMember)

			} else if update.ChatMember != nil {
				t.chatMember(update.ChatMember)

			} else {
				fmt.Println(1, update)
			}
		}
	}()

	t.t = *TgBot

}
