package Tg

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func (t *Telegram) InitTG(tokent string) {
	//подключение к телеграм
	TgBot, Err := tgbotapi.NewBotAPI(tokent)
	if Err != nil {
		log.Panic("ошибка подключения к телеграм ", Err)
	}
	TgBot.Debug = false
	fmt.Printf("Бот TELEGRAM загружен  %s\n", TgBot.Self.UserName)
	ut := tgbotapi.NewUpdate(0)
	ut.Timeout = 60
	go func() { //получаем обновления от телеграм
		updates := TgBot.GetUpdatesChan(ut)
		for update := range updates {
			if update.CallbackQuery != nil {
				//callback(update.CallbackQuery) //нажатия в чате
			} else if update.Message != nil {
				if update.Message.Chat.IsPrivate() { //если пишут боту в личку
					//tgSendChannel(update.Message.Chat.ID, "сорян это в разработке ")
				} else if update.Message.IsCommand() {
					t.updatesComand(update.Message) //если сообщение является командой
				} else { //остальные сообщения
					fmt.Println("test", update.Message.Text)
					t.logicMixTelegram(update.Message)
				}

			} else if update.MyChatMember != nil {
				//myChatMember(update.MyChatMember)

			} else if update.EditedMessage != nil {
				log.Println("Измененный текст в телеге ", update.EditedMessage.Text)
			} else {
				fmt.Println(1, update)
			}
		}
	}()

	t.t = *TgBot
}
