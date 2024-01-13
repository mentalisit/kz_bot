package TelegramClient

import (
	tgbotapi "github.com/samuelemusiani/telegram-bot-api"
)

func (t *Telegram) nameOrNick(UserName, FirstName string) (name string) {
	if UserName != "" {
		name = UserName

	} else {
		name = FirstName
	}
	return name
}

func (t *Telegram) GetAvatar(userid int64) string {
	userProfilePhotos, err := t.t.GetUserProfilePhotos(tgbotapi.UserProfilePhotosConfig{UserID: userid})
	if err != nil {
		return "https://thumb.cloud.mail.ru/weblink/thumb/xw1/VLES/v7tqy1nXQ/telegram.png"
	}
	if len(userProfilePhotos.Photos) == 0 {
		return "https://thumb.cloud.mail.ru/weblink/thumb/xw1/VLES/v7tqy1nXQ/telegram.png"
	}
	fileconfig := tgbotapi.FileConfig{FileID: userProfilePhotos.Photos[0][0].FileID}
	file, err := t.t.GetFile(fileconfig)
	if err != nil {
		t.log.Error(err.Error())
		return ""
	}
	return "https://api.telegram.org/file/bot" + t.t.Token + "/" + file.FilePath
}
