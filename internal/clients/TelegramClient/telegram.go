package TelegramClient

import (
	"fmt"
	tgbotapi "github.com/matterbridge/telegram-bot-api/v6"
	"github.com/mentalisit/logger"
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
	a1 := corpCompendium{
		name:    "HS UA Community",
		storage: "HS UA Community",
		chatid:  -1002116077159,
	}
	corp = append(corp, a1)
	a2 := corpCompendium{
		name:    "UAGC",
		storage: "UAGC",
		chatid:  -1001194014201,
	}
	corp = append(corp, a2)
	a3 := corpCompendium{
		name:    "test3",
		storage: "HS UA Community",
		chatid:  -1001556223093,
	}
	corp = append(corp, a3)
	go tg.update()

	return tg
}
func (t *Telegram) update() {
	ut := tgbotapi.NewUpdate(0)
	ut.Timeout = 60
	//получаем обновления от телеграм
	updates := t.t.GetUpdatesChan(ut)
	for update := range updates {
		if update.InlineQuery != nil {
			t.handleInlineQuery(update.InlineQuery)
		} else if update.ChosenInlineResult != nil {
			//go t.handleChosenInlineResult(update.ChosenInlineResult)
		} else if update.CallbackQuery != nil {
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
			go func() {
				if update.Poll != nil {
					t.log.InfoStruct("pool ", update.Poll)
				} else if update.EditedChannelPost != nil {

				} else if update.ChannelPost != nil {

				} else {
					t.log.Info(fmt.Sprintf(" else update: %+v \n", update))
				}
			}()
		}
	}
}
func (t *Telegram) ifPrivatMesage(m *tgbotapi.Message) {
	if m.From.UserName == "mentalisit" && m.Text == "/update" {
		utils.UpdateRun()
	} else {
		t.SendChannel(strconv.FormatInt(m.Chat.ID, 10), "сорян это в разработке \n"+
			"я еще не решил как тут сделать"+
			"Присылай идеи для работы с ботом мне @Mentalisit ")
	}
}
