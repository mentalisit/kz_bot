package TelegramClient

import (
	"fmt"
	"github.com/gofrs/uuid"
	tgbotapi "github.com/matterbridge/telegram-bot-api/v6"
	"kz_bot/internal/compendiumCli"
	"kz_bot/internal/config"
	"kz_bot/internal/models"
	"strconv"
	"strings"
)

const nickname = "Для того что бы БОТ мог Вас индентифицировать, создайте уникальный НикНей в настройках. Вы можете использовать a-z, 0-9 и символы подчеркивания. Минимальная длина - 5 символов."

func (t *Telegram) callback(cb *tgbotapi.CallbackQuery) {
	callback := tgbotapi.NewCallback(cb.ID, cb.Data)
	if _, err := t.t.Request(callback); err != nil {
		t.log.ErrorErr(err)
	}
	ChatId := strconv.FormatInt(cb.Message.Chat.ID, 10) + fmt.Sprintf("/%d", cb.Message.MessageThreadID)
	ok, config := t.CheckChannelConfigTG(ChatId)
	if ok {
		name := t.nameNick(cb.From.UserName, cb.From.FirstName, ChatId)
		in := models.InMessage{
			Mtext:       cb.Data,
			Tip:         "tg",
			Name:        name,
			NameMention: "@" + name,
			Tg: struct {
				Mesid int
			}{
				Mesid: cb.Message.MessageID,
			},
			Config: config,
			Option: models.Option{
				Reaction: true},
		}

		t.ChanRsMessage <- in
	}
}

func (t *Telegram) myChatMember(member *tgbotapi.ChatMemberUpdated) {
	ChatId := strconv.FormatInt(member.Chat.ID, 10) + "/0"
	if member.NewChatMember.Status == "member" {
		t.SendChannelDelSecond(ChatId, fmt.Sprintf("@%s мне нужны права админа для коректной работы", member.From.UserName), 60)
	} else if member.NewChatMember.Status == "administrator" {
		t.SendChannelDelSecond(ChatId, fmt.Sprintf("@%s спасибо ... я готов к работе \nАктивируй нужный режим бота,\n если сложности пиши мне @Mentalisit", member.From.UserName), 60)
	}
}

func (t *Telegram) chatMember(chMember *tgbotapi.ChatMemberUpdated) {
	if chMember.NewChatMember.IsMember {
		ChatId := strconv.FormatInt(chMember.Chat.ID, 10) + "/0"
		t.SendChannelDelSecond(ChatId,
			fmt.Sprintf("%s Добро пожаловать в наш чат ", chMember.NewChatMember.User.FirstName),
			60)
	}

}
func (t *Telegram) nameNick(UserName, FirstName string, chatid string) (name string) {
	if UserName != "" {
		name = UserName
	} else {
		name = FirstName
		go t.SendChannelDelSecond(chatid, nickname, 60)
	}
	return name
}

func (t *Telegram) handleDownload(message *tgbotapi.Message) (url, fileName string) {
	size := int64(0)
	switch {
	case message.Sticker != nil:
		url, _ = t.t.GetFileDirectURL(message.Sticker.FileID)
		size = int64(message.Sticker.FileSize)
	case message.Voice != nil:
		url, _ = t.t.GetFileDirectURL(message.Voice.FileID)
		size = message.Voice.FileSize
	case message.Video != nil:
		url, _ = t.t.GetFileDirectURL(message.Video.FileID)
		size = message.Video.FileSize
		fileName = message.Video.FileName
	case message.Audio != nil:
		url, _ = t.t.GetFileDirectURL(message.Audio.FileID)
		size = message.Audio.FileSize
		fileName = message.Audio.FileName
	case message.Document != nil:
		url, _ = t.t.GetFileDirectURL(message.Document.FileID)
		size = message.Document.FileSize
		fileName = message.Document.FileName
	case message.Photo != nil:
		photos := message.Photo
		size = int64(photos[len(photos)-1].FileSize)
		url, _ = t.t.GetFileDirectURL(photos[len(photos)-1].FileID)

	}
	if size > 25000000 {
		fmt.Println("big size")
		message.Text += " файл слишком большой для пересылки"
		return "", ""
	}
	var urlReplace = "https://api.telegram.org/file/bot" + config.Instance.Token.TokenTelegram
	url = strings.Replace(url, urlReplace, "http://mentalisit.sytes.net:4243", 1)
	fmt.Println(url)
	return url, fileName
}

func (t *Telegram) handlePoll(message *tgbotapi.Message) {
	if message.Poll != nil {
		text := "Запущен  ОПРОС  \n"
		text += message.Poll.Question
		text += "\nВарианты ответа:\n"
		for _, o := range message.Poll.Options {
			text += fmt.Sprintf(" %s\n", o.Text)
		}
		message.Text = text
	}
}

type corpCompendium struct {
	name    string
	storage string
	chatid  int64
}

var corp []corpCompendium

func (t *Telegram) handleInlineQuery(m *tgbotapi.InlineQuery) {
	t.log.Info(fmt.Sprintf("handleInlineQuery от %s text:%s\n", m.From.UserName, m.Query))
	if m.Query == "u" {
		fromString, err := uuid.DefaultGenerator.NewV1()
		if err != nil {
			t.log.ErrorErr(err)
			return
		}
		anArticle := tgbotapi.NewInlineQueryResultArticleMarkdownV2(
			fromString.String(),
			"Users",
			"",
		)
		anArticle.InputMessageContent = tgbotapi.InputTextMessageContent{
			Text:      "%users",
			ParseMode: tgbotapi.ModeMarkdownV2,
		}

		_, _ = t.t.Request(tgbotapi.InlineConfig{
			InlineQueryID: m.ID,
			Results:       []interface{}{anArticle},
		})

	} else if m.Query == "" {
		var res []interface{}
		for _, c := range corp {
			fromString, err := uuid.DefaultGenerator.NewV1()
			if err != nil {
				t.log.ErrorErr(err)
				return
			}
			compendium, err := compendiumCli.GetCompendium(t.log, "", c.storage)
			if err != nil {
				t.log.ErrorErr(err)
				compendium.Shutdown()
				return
			}
			members, err := compendium.GetRoleMembers("")
			if err != nil {
				t.log.ErrorErr(err)
				compendium.Shutdown()
				return
			}
			compendium.Shutdown()
			text := "Пользователи  дискорд " + c.name + "\n\n"
			for _, member := range members {
				text += member.Name + "\n"
			}

			anArticle := tgbotapi.NewInlineQueryResultArticle(fromString.String(), c.name, "")
			anArticle.InputMessageContent = tgbotapi.InputTextMessageContent{
				Text:      text,
				ParseMode: tgbotapi.ModeHTML,
			}

			res = append(res, anArticle)
		}

		_, err := t.t.Request(tgbotapi.InlineConfig{
			InlineQueryID: m.ID,
			Results:       res,
			CacheTime:     0,
		})
		if err != nil {
			t.log.ErrorErr(err)
		}

	} else if m.Query == "user" {
		compendium, err := compendiumCli.GetCompendium(t.log, "", "HS UA Community")
		if err != nil {
			t.log.ErrorErr(err)
			compendium.Shutdown()
			return
		}
		members, err := compendium.GetRoleMembers("")
		if err != nil {
			t.log.ErrorErr(err)
			compendium.Shutdown()
			return
		}
		compendium.Shutdown()

		var art []interface{}

		for _, member := range members {
			fromString, err := uuid.DefaultGenerator.NewV1()
			if err != nil {
				t.log.ErrorErr(err)
				return
			}
			anArticle := tgbotapi.NewInlineQueryResultArticleHTML(
				fromString.String(),
				member.Name,
				member.Name,
			)
			anArticle.InputMessageContent = tgbotapi.InputTextMessageContent{
				Text:      "%user " + member.Name,
				ParseMode: tgbotapi.ModeHTML,
			}
			art = append(art, anArticle)
		}

		_, err = t.t.Request(tgbotapi.InlineConfig{
			InlineQueryID: m.ID,
			Results:       art,
		})
		if err != nil {
			t.log.ErrorErr(err)
			return
		}
	}
}

//func (t *Telegram) handleChosenInlineResult(m *tgbotapi.ChosenInlineResult) {
//	t.log.Info(fmt.Sprintf("handleChosenInlineResult  %+v \nFrom: %+v\n", m, m.From))
//	if m.ResultID == "Corp1" {
//		// Создаем меню с именами
//		var members = []string{"Alice", "Bob", "Charlie", "David"}
//		menu := make([][]tgbotapi.InlineKeyboardButton, 0)
//		row := make([]tgbotapi.InlineKeyboardButton, 0)
//		for _, member := range members {
//			btn := tgbotapi.NewInlineKeyboardButtonData(member, member)
//			row = append(row, btn)
//		}
//		menu = append(menu, row)
//		chatid, err := strconv.ParseInt(m.InlineMessageID, 10, 64)
//		if err != nil {
//			t.log.ErrorErr(err)
//		}
//		msg := tgbotapi.NewMessage(chatid, "Выберите имя:")
//
//		msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(menu...)
//
//		if _, err := t.t.Send(msg); err != nil {
//			log.Println("Error sending menu:", err)
//		}
//	}
//}
