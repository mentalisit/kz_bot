package TelegramClient

import (
	"fmt"
	tgbotapi "github.com/matterbridge/telegram-bot-api/v6"
	"kz_bot/internal/compendiumCli"
	"kz_bot/pkg/imageGenerator"
	"strings"
)

func (t *Telegram) prefixCompendium(m *tgbotapi.Message, chatid string) bool {
	after, found := strings.CutPrefix(m.Text, "%")
	_, c := t.CheckChannelCompendium(m.Chat.ID)
	if found && (m.Chat.ID == -1002116077159 || m.Chat.ID == -1001194014201 || m.Chat.ID == -1001556223093 || m.Chat.ID == -1001873861265) { //HS UA Community,UAGC, test room,bz
		switch after {
		case "t i":
			return t.techImage(chatid, m.From.UserName, c.storage)
		case "users":
			return t.getUsersCompendium(chatid, c.storage)
		default:
			split := strings.Split(after, " ")
			if len(split) > 1 {
				if split[0] == "user" {
					username := split[1]
					return t.techImage(chatid, username, c.storage)
				}
			}
		}
	} else if found {
		t.log.Info(fmt.Sprintf("Запрос с чата %s %s", t.ChatName(chatid), chatid))
		t.SendChannel(chatid, "Для подключения режима компендиум в этот чат обратитесь к @mentalisit")
		_, _ = t.t.Send(tgbotapi.NewMessage(m.From.ID, "Не стесняйся мне интересны новые вызовы :) пиши @mentalisit "))
	}
	return false
}
func (t *Telegram) techImage(chatid string, UserName string, storage string) bool {
	t.ChatTyping(chatid)
	compendium, err := compendiumCli.GetCompendium(t.log, "", storage)
	if err != nil {
		t.log.ErrorErr(err)
		t.SendChannel(chatid, fmt.Sprintf("Произошол сбой нуждается в дороботке "))
		compendium.Shutdown()
		return false
	}
	member, err := compendium.GetMember("", UserName)
	if err != nil {
		t.log.Info(fmt.Sprintf("Игрок под ником %s не найден запрос с %s", UserName, t.ChatName(chatid)))
		t.SendChannel(chatid, fmt.Sprintf("Игрок под ником %s не найден", UserName))
		compendium.Shutdown()
		return false
	}

	t.SendChannelDelSecond(chatid, "Выполняется генерация картинки", 5)
	t.ChatTyping(chatid)

	userPic := imageGenerator.GenerateUser(member.AvatarURL, t.getChatPhoto(chatid), UserName, t.ChatName(chatid), member.Tech)
	t.SendFilePic(chatid, "Вот картинка", userPic)
	return true
}
func (t *Telegram) getUsersCompendium(chatid string, storage string) bool {
	t.ChatTyping(chatid)
	chatId, threadID := t.chat(chatid)
	compendium, err := compendiumCli.GetCompendium(t.log, "", storage)
	if err != nil {
		t.log.ErrorErr(err)
		t.SendChannel(chatid, fmt.Sprintf("Произошол сбой нуждается в дороботке "))
		compendium.Shutdown()
		return false
	}
	members, err := compendium.GetRoleMembers("")
	if err != nil {
		t.log.ErrorErr(err)
		compendium.Shutdown()
		return false
	}
	text := ""
	for _, member := range members {
		text += member.Name + "\n"
	}
	mes := tgbotapi.NewMessage(chatId, text)
	mes.MessageThreadID = threadID
	_, err1 := t.t.Send(mes)
	if err1 != nil {
		t.log.Error(err1.Error())
		compendium.Shutdown()
	}
	return true
}
