package TelegramClient

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
	"kz_bot/internal/config"
	"kz_bot/internal/models"
	"kz_bot/internal/storage"
	"kz_bot/pkg/clientTelegram"
)

type Telegram struct {
	inbox   chan models.InMessage
	toGame  chan models.Message
	t       *tgbotapi.BotAPI
	log     *logrus.Logger
	storage *storage.Storage
	debug   bool
}

func NewTelegram(inbox chan models.InMessage, togame chan models.Message, log *logrus.Logger, st *storage.Storage, cfg *config.ConfigBot) *Telegram {
	client, err := clientTelegram.NewTelegram(log, cfg)
	if err != nil {
		return nil
	}

	tg := &Telegram{
		inbox:   inbox,
		toGame:  togame,
		t:       client,
		log:     log,
		storage: st,
		debug:   cfg.IsDebug,
	}

	go tg.update()

	return tg
}
func (t Telegram) update() {
	ut := tgbotapi.NewUpdate(0)
	ut.Timeout = 60
	//получаем обновления от телеграм
	updates := t.t.GetUpdatesChan(ut)
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
				t.logicMix(update.Message)
			}

		} else if update.MyChatMember != nil {
			t.myChatMember(update.MyChatMember)

		} else if update.ChatMember != nil {
			t.chatMember(update.ChatMember)

		} else {
			fmt.Println(1, update)
		}
	}
}