package TelegramClient

import (
	"bytes"
	tgbotapi "github.com/matterbridge/telegram-bot-api/v6"
	"io"
	"kz_bot/internal/models"
	"net/http"
	"net/url"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"
)

func (t *Telegram) SendEmded(lvlkz string, chatid string, text string) int {
	a := strings.SplitN(chatid, "/", 2)
	chatId, err := strconv.ParseInt(a[0], 10, 64)
	if err != nil {
		t.log.ErrorErr(err)
	}
	ThreadID := 0
	if len(a) > 1 {
		ThreadID, err = strconv.Atoi(a[1])
		if err != nil {
			t.log.ErrorErr(err)
		}
	}
	var keyboardQueue = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(lvlkz+"+", lvlkz+"+"),
			tgbotapi.NewInlineKeyboardButtonData(lvlkz+"-", lvlkz+"-"),
			tgbotapi.NewInlineKeyboardButtonData(lvlkz+"++", lvlkz+"++"),
			tgbotapi.NewInlineKeyboardButtonData(lvlkz+"+30", lvlkz+"+++"),
		),
	)
	msg := tgbotapi.NewMessage(chatId, text)
	msg.MessageThreadID = ThreadID
	msg.ReplyMarkup = keyboardQueue
	message, _ := t.t.Send(msg)

	return message.MessageID

}
func (t *Telegram) SendEmbedTime(chatid string, text string) int {
	a := strings.SplitN(chatid, "/", 2)
	chatId, err := strconv.ParseInt(a[0], 10, 64)
	if err != nil {
		t.log.ErrorErr(err)
	}
	ThreadID := 0
	if len(a) > 1 {
		ThreadID, err = strconv.Atoi(a[1])
		if err != nil {
			t.log.ErrorErr(err)
		}
	}

	var keyboardQueue = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("+", "+"),
			tgbotapi.NewInlineKeyboardButtonData("-", "-"),
		),
	)
	msg := tgbotapi.NewMessage(chatId, text)
	msg.MessageThreadID = ThreadID
	msg.ReplyMarkup = keyboardQueue
	message, _ := t.t.Send(msg)

	return message.MessageID
}

// отправка сообщения в телегу
func (t *Telegram) SendChannel(chatid string, text string) int {
	a := strings.SplitN(chatid, "/", 2)
	chatId, err := strconv.ParseInt(a[0], 10, 64)
	if err != nil {
		t.log.ErrorErr(err)
	}
	ThreadID := 0
	if len(a) > 1 {
		ThreadID, err = strconv.Atoi(a[1])
		if err != nil {
			t.log.ErrorErr(err)
		}
	}
	m := tgbotapi.NewMessage(chatId, text)
	m.MessageThreadID = ThreadID
	tMessage, _ := t.t.Send(m)
	return tMessage.MessageID
}

func (t *Telegram) SendChannelDelSecond(chatid string, text string, second int) {
	a := strings.SplitN(chatid, "/", 2)
	chatId, err := strconv.ParseInt(a[0], 10, 64)
	if err != nil {
		t.log.ErrorErr(err)
	}
	ThreadID := 0
	if len(a) > 1 {
		ThreadID, err = strconv.Atoi(a[1])
		if err != nil {
			t.log.ErrorErr(err)
		}
	}
	m := tgbotapi.NewMessage(chatId, text)
	m.MessageThreadID = ThreadID
	tMessage, err1 := t.t.Send(m)
	if err1 != nil {
		t.log.Error(err1.Error())
	}
	if second <= 60 {
		go func() {
			time.Sleep(time.Duration(second) * time.Second)
			t.DelMessage(chatid, tMessage.MessageID)
		}()
	} else {
		t.storage.TimeDeleteMessage.TimerInsert(models.Timer{
			Tgmesid:  strconv.Itoa(tMessage.MessageID),
			Tgchatid: chatid,
			Timed:    second,
		})
	}
}

func (t *Telegram) SendChannelAsync(chatid string, text string, resultChannel chan<- models.MessageTg, wg *sync.WaitGroup) {
	defer wg.Done()
	a := strings.SplitN(chatid, "/", 2)
	chatId, err := strconv.ParseInt(a[0], 10, 64)
	if err != nil {
		t.log.ErrorErr(err)
	}
	ThreadID := 0
	if len(a) > 1 {
		ThreadID, err = strconv.Atoi(a[1])
		if err != nil {
			t.log.ErrorErr(err)
		}
	}
	m := tgbotapi.NewMessage(chatId, text)
	m.MessageThreadID = ThreadID
	tMessage, _ := t.t.Send(m)
	messageData := models.MessageTg{
		MessageId: strconv.Itoa(tMessage.MessageID),
		ChatId:    chatid,
	}
	resultChannel <- messageData
}
func (t *Telegram) SendFileFromURLAsync(chatid, text string, fileURL string, resultChannel chan<- models.MessageTg, wg *sync.WaitGroup) {
	defer wg.Done()
	fileURL = strings.TrimSpace(fileURL)
	a := strings.SplitN(chatid, "/", 2)
	chatId, err := strconv.ParseInt(a[0], 10, 64)
	if err != nil {
		t.log.ErrorErr(err)
	}
	ThreadID := 0
	if len(a) > 1 {
		ThreadID, err = strconv.Atoi(a[1])
		if err != nil {
			t.log.ErrorErr(err)
		}
	}

	parsedURL, err := url.Parse(fileURL)
	if err != nil {
		t.log.ErrorErr(err)
		return
	}

	// Используем path.Base для получения последней части URL, которая представляет собой имя файла
	fileName := path.Base(parsedURL.Path)
	parsedURL.RawQuery = ""
	fileURL = parsedURL.String()

	// Скачиваем файл по URL
	resp, err := http.Get(fileURL)
	if err != nil {
		t.log.ErrorErr(err)
		return
	}
	defer resp.Body.Close()

	// Читаем содержимое файла
	buffer := new(bytes.Buffer)
	_, err = io.Copy(buffer, resp.Body)
	if err != nil {
		t.log.ErrorErr(err)
		return
	}
	var media []interface{}

	file := tgbotapi.FileBytes{
		Name:  fileName,
		Bytes: buffer.Bytes(),
	}

	switch filepath.Ext(fileName) {

	case ".jpg", ".jpe", ".png":
		pc := tgbotapi.NewInputMediaPhoto(file)
		if text != "" {
			pc.Caption = text
		}
		media = append(media, pc)
	case ".mp4", ".m4v":
		vc := tgbotapi.NewInputMediaVideo(file)
		if text != "" {
			vc.Caption = text
		}
		media = append(media, vc)
	case ".mp3", ".oga":
		ac := tgbotapi.NewInputMediaAudio(file)
		if text != "" {
			ac.Caption = text
		}
		media = append(media, ac)
	default:
		dc := tgbotapi.NewInputMediaDocument(file)
		if text != "" {
			dc.Caption = text
		}
		media = append(media, dc)
	}

	if len(media) == 0 {
		return
	}
	mg := tgbotapi.MediaGroupConfig{
		BaseChat: tgbotapi.BaseChat{
			ChatID:          chatId,
			MessageThreadID: ThreadID,
			//ChannelUsername:  msg.Username,
			//ReplyToMessageID: parentID,
		},
		Media: media,
	}
	m, err := t.t.SendMediaGroup(mg)
	if err != nil {
		t.log.ErrorErr(err)
		return
	}

	messageData := models.MessageTg{
		MessageId: strconv.Itoa(m[0].MessageID),
		ChatId:    chatid,
	}
	resultChannel <- messageData
}
func (t *Telegram) SendFilePic(chatId string, text string, f *bytes.Reader) {
	chatid, threadID := t.chat(chatId)
	// Читаем содержимое файла
	buffer := new(bytes.Buffer)
	_, err := io.Copy(buffer, f)
	if err != nil {
		t.log.ErrorErr(err)
		return
	}
	var media []interface{}

	file := tgbotapi.FileBytes{
		//Name:  fileName,
		Bytes: buffer.Bytes(),
	}

	switch filepath.Ext("fileName.png") {

	case ".jpg", ".jpe", ".png":
		pc := tgbotapi.NewInputMediaPhoto(file)
		if text != "" {
			pc.Caption = text
		}
		media = append(media, pc)
	default:
		dc := tgbotapi.NewInputMediaDocument(file)
		if text != "" {
			dc.Caption = text
		}
		media = append(media, dc)
	}

	if len(media) == 0 {
		return
	}
	mg := tgbotapi.MediaGroupConfig{
		BaseChat: tgbotapi.BaseChat{
			ChatID:          chatid,
			MessageThreadID: threadID,
		},
		Media: media,
	}
	_, err = t.t.SendMediaGroup(mg)
	if err != nil {
		t.log.ErrorErr(err)
		return
	}
}
