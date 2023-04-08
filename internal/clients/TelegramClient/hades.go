package TelegramClient

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"kz_bot/internal/models"
	"kz_bot/internal/storage/CorpsConfig/hades"
	"regexp"
)

func (t *Telegram) ifMessageForHades(m *tgbotapi.Message) {
	if t.ifComands(m) {
		return
	}
	if m.Text == "" || filterRsPl(m.Text) {
		return
	}
	ok, corp := hades.HadesStorage.AllianceChatTg(m.Chat.ID)
	if ok {
		if m.Text != "" {

			mes := models.Message{
				Text:        m.Text,
				Sender:      t.nameOrNick(m.From.UserName, m.From.FirstName),
				Avatar:      t.GetAvatar(m.From.ID),
				ChannelType: 0,
				Corporation: corp.Corp,
				Command:     "text",
				Messager:    "tg",
			}
			t.toGame <- mes
		}
	}
	okWs, corp := hades.HadesStorage.Ws1ChatTg(m.Chat.ID)
	if okWs {
		if m.Text != "" {
			mes := models.Message{
				Text:        m.Text,
				Sender:      t.nameOrNick(m.From.UserName, m.From.FirstName),
				Avatar:      t.GetAvatar(m.From.ID),
				ChannelType: 1,
				Corporation: corp.Corp,
				Command:     "text",
				Messager:    "tg",
			}
			t.toGame <- mes
		}
	}
}
func (t *Telegram) nameOrNick(UserName, FirstName string) (name string) {
	if UserName != "" {
		return UserName

	} else {
		return FirstName
	}
}

func (t *Telegram) GetAvatar(userid int64) string {
	userProfilePhotos, err := t.t.GetUserProfilePhotos(tgbotapi.UserProfilePhotosConfig{UserID: userid})
	if err != nil {
		t.log.Println("err GetAvatar " + err.Error())
		return "https://thumb.cloud.mail.ru/weblink/thumb/xw1/VLES/v7tqy1nXQ/telegram.png"
	}
	//t.log.Printf("size photo %d", len(userProfilePhotos.Photos))
	if len(userProfilePhotos.Photos) == 0 {
		return "https://thumb.cloud.mail.ru/weblink/thumb/xw1/VLES/v7tqy1nXQ/telegram.png"
	}
	fileconfig := tgbotapi.FileConfig{FileID: userProfilePhotos.Photos[0][0].FileID}
	file, err := t.t.GetFile(fileconfig)
	if err != nil {
		t.log.Println("err GetAvatar File " + err.Error())
		return ""
	}
	return "https://api.telegram.org/file/bot" + t.t.Token + "/" + file.FilePath
}
func filterRsPl(s string) bool {
	re := regexp.MustCompile(`^([3-9]|[1][0-2])[\+]$`)
	match := re.MatchString(s)
	return match
}
