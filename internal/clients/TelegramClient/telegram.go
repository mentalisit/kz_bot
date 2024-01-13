package TelegramClient

import (
	"fmt"
	tgbotapi "github.com/samuelemusiani/telegram-bot-api"
	"kz_bot/pkg/logger"
	"strconv"

	"kz_bot/internal/config"
	"kz_bot/internal/models"
	"kz_bot/internal/storage"
	"kz_bot/pkg/clientTelegram"
	"kz_bot/pkg/utils"
)

type Telegram struct {
	ChanRsMessage     chan models.InMessage
	ChanBridgeMessage chan models.BridgeMessage
	t                 *tgbotapi.BotAPI
	log               *logger.Logger
	storage           *storage.Storage
	debug             bool
	bridgeConfig      map[string]models.BridgeConfig
	corpConfigRS      map[string]models.CorporationConfig
}

func NewTelegram(log *logger.Logger, st *storage.Storage, cfg *config.ConfigBot) *Telegram {
	client, err := clientTelegram.NewTelegram(log, cfg)
	if err != nil {
		return nil
	}

	tg := &Telegram{
		ChanRsMessage:     make(chan models.InMessage, 10),
		ChanBridgeMessage: make(chan models.BridgeMessage, 20),
		t:                 client,
		log:               log,
		storage:           st,
		debug:             cfg.IsDebug,
		bridgeConfig:      st.BridgeConfigs,
		corpConfigRS:      st.CorpConfigRS,
	}

	go tg.update()

	return tg
}
func (t *Telegram) update() {
	ut := tgbotapi.NewUpdate(0)
	ut.Timeout = 60
	//получаем обновления от телеграм
	updates := t.t.GetUpdatesChan(ut)
	for update := range updates {
		if update.CallbackQuery != nil {
			t.callback(update.CallbackQuery) //нажатия в чате
		} else if update.Message != nil {

			if update.Message.Chat.IsPrivate() { //если пишут боту в личку
				t.ifPrivatMesage(update.Message)
			} else if update.Message.IsCommand() {
				t.updatesComand(update.Message) //если сообщение является командой
			} else { //остальные сообщения
				t.logicMix(update.Message, false)
			}
		} else if update.EditedMessage != nil {
			t.logicMix(update.EditedMessage, true)
		} else if update.MyChatMember != nil {
			t.myChatMember(update.MyChatMember)

		} else if update.ChatMember != nil {
			t.chatMember(update.ChatMember)

		} else {
			fmt.Printf(" else update %+v", update)
		}
	}
}
func (t *Telegram) ifPrivatMesage(m *tgbotapi.Message) {
	if m.From.UserName == "Mentalisit" && m.Text == "/update" {
		utils.UpdateRun()
	} else {
		t.SendChannel(strconv.FormatInt(m.Chat.ID, 10), "сорян это в разработке \n"+
			"я еще не решил как тут сделать"+
			"Присылай идеи для работы с ботом мне @Mentalisit ")
	}
}
