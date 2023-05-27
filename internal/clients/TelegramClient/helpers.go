package TelegramClient

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"regexp"
)

func (t *Telegram) nameOrNick(UserName, FirstName string) (name string) {
	if UserName != "" {
		name = UserName

	} else {
		name = FirstName
	}
	name = replaceGameName(name)

	return name
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
func (t *Telegram) GetPic(fileid string) string {
	fileconfig := tgbotapi.FileConfig{FileID: fileid}
	file, err := t.t.GetFile(fileconfig)
	if err != nil {
		t.log.Println("err GetPic File " + err.Error())
		return ""
	}
	return "https://api.telegram.org/file/bot" + t.t.Token + "/" + file.FilePath
}

func filterRsPl(s string) bool {
	re := regexp.MustCompile(`^([3-9]|[1][0-2])[\+]$`)
	match := re.MatchString(s)
	return match
}
func replaceGameName(s string) string {
	type list struct {
		nameGame     string
		nameTelegram string
	}
	userList := []list{
		{nameGame: "Колхоз", nameTelegram: "andvs"},
		{nameGame: "Ivan", nameTelegram: "Ivan_Belskiy"},
		{nameGame: "Vovkasotka", nameTelegram: "HexagonChip"},
		{nameGame: "Джон Джонович", nameTelegram: "i_kebab"},
		{nameGame: "Encounter", nameTelegram: "Encounter1793"},
		{nameGame: "Angel", nameTelegram: "Angel_12346"},
		{nameGame: "Менталисит", nameTelegram: "Mentalisit"},
	}
	for _, l := range userList {
		if l.nameTelegram == s {
			return l.nameGame
		}
	}
	return s
}
