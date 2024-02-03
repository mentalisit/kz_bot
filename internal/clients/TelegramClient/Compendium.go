package TelegramClient

import (
	"fmt"
	tgbotapi "github.com/samuelemusiani/telegram-bot-api"
	"kz_bot/internal/compendiumCli"
	"kz_bot/pkg/imageGenerator"
	"strings"
)

func (t *Telegram) prefixCompendium(m *tgbotapi.Message, chatid string) bool {
	after, found := strings.CutPrefix(m.Text, "%")
	if found && (m.Chat.ID == -1002116077159 || m.Chat.ID == -1001556223093) { //HS UA Community and test room
		switch after {
		case "t i":
			{
				return t.techImage(chatid, m.From.UserName)
			}
		case "users":
			{
				return t.getUsersCompendium(chatid)
			}
		default:
			split := strings.Split(after, " ")
			if len(split) > 1 {
				if split[0] == "user" {
					username := split[1]
					return t.techImage(chatid, username)
				}
			}
		}
	}
	return false
}
func (t *Telegram) techImage(chatid string, UserName string) bool {
	fmt.Println("techImage")
	compendium, err := compendiumCli.GetCompendium(t.log, "5W9Z-FJgL-VKVW", "testkey")
	if err != nil {
		t.log.ErrorErr(err)
		return false
	}
	member, err := compendium.GetMember("", UserName)
	if err != nil {
		t.log.ErrorErr(err)
		return false
	}

	userPic := imageGenerator.GenerateUser(member.AvatarURL, t.getChatPhoto(chatid), UserName, t.ChatName(chatid), member.Tech)
	t.SendFilePic(chatid, "Вот картинка", userPic)
	return true
}
func (t *Telegram) getUsersCompendium(chatid string) bool {
	chatId, threadID := t.chat(chatid)
	compendium, err := compendiumCli.GetCompendium(t.log, "5W9Z-FJgL-VKVW", "testkey")
	if err != nil {
		t.log.ErrorErr(err)
		return false
	}
	members, err := compendium.GetRoleMembers("")
	if err != nil {
		t.log.ErrorErr(err)
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
	}
	return true
}
