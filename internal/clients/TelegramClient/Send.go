package TelegramClient

import (
	"fmt"
	tgbotapi "github.com/musianisamuele/telegram-bot-api"
	"io"
	"kz_bot/internal/models"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func (t *Telegram) SendEmded(lvlkz string, chatid string, text string) int {
	a := strings.SplitN(chatid, "/", 2)
	chatId, err := strconv.ParseInt(a[0], 10, 64)
	if err != nil {
		t.log.Println(err)
	}
	ThreadID := 0
	if len(a) > 1 {
		ThreadID, err = strconv.Atoi(a[1])
		if err != nil {
			t.log.Println(err)
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
		t.log.Println(err)
	}
	ThreadID := 0
	if len(a) > 1 {
		ThreadID, err = strconv.Atoi(a[1])
		if err != nil {
			t.log.Println(err)
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
		t.log.Println(err)
	}
	ThreadID := 0
	if len(a) > 1 {
		ThreadID, err = strconv.Atoi(a[1])
		if err != nil {
			t.log.Println(err)
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
		t.log.Println(err)
	}
	ThreadID := 0
	if len(a) > 1 {
		ThreadID, err = strconv.Atoi(a[1])
		if err != nil {
			t.log.Println(err)
		}
	}
	m := tgbotapi.NewMessage(chatId, text)
	m.MessageThreadID = ThreadID
	tMessage, err1 := t.t.Send(m)
	if err1 != nil {
		t.log.Println(err)
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
func (t *Telegram) SendFileFromURL(chatid string, fileURL string) int {
	fileURL = strings.TrimSpace(fileURL)
	a := strings.SplitN(chatid, "/", 2)
	chatId, err := strconv.ParseInt(a[0], 10, 64)
	if err != nil {
		t.log.Println(err)
	}
	ThreadID := 0
	if len(a) > 1 {
		ThreadID, err = strconv.Atoi(a[1])
		if err != nil {
			t.log.Println(err)
		}
	}

	parsedURL, err := url.Parse(fileURL)
	if err != nil {
		fmt.Println(err)
		//return 0
	}

	// Используем path.Base для получения последней части URL, которая представляет собой имя файла
	fileName := path.Base(parsedURL.Path)
	parsedURL.RawQuery = ""

	fmt.Println(parsedURL.String())

	// Получаем содержимое файла по URL
	response, err := http.Get(parsedURL.String())
	if err != nil {

	}
	defer response.Body.Close()

	// Открываем временный файл для сохранения содержимого
	tempFile, err := os.CreateTemp("", fileName)
	if err != nil {
		fmt.Println("151")
		return 0
	}
	defer tempFile.Close()
	fmt.Println("153")
	// Копируем содержимое файла из ответа во временный файл
	_, err = io.Copy(tempFile, response.Body)
	if err != nil {
		return 0
	}
	fmt.Println("159")
	// Создаем объект InputFile для отправки файла
	document := &tgbotapi.FileBytes{
		Name:  fileName,
		Bytes: getFileBytes(tempFile),
	}

	// Создаем сообщение с файлом
	msg := tgbotapi.NewDocument(chatId, *document)
	msg.MessageThreadID = ThreadID
	msg.Caption = "Текст к файлу (если нужен)"
	fmt.Println("170")
	// Отправляем сообщение с файлом
	m, err := t.t.Send(&msg)
	if err != nil {
		return 0
	}

	fmt.Println("Файл успешно отправлен в Telegram.")
	return m.MessageID
}

// getFileBytes считывает байты из файла
func getFileBytes(file *os.File) []byte {
	fileInfo, _ := file.Stat()
	size := fileInfo.Size()
	bytes := make([]byte, size)

	_, err := file.Read(bytes)
	if err != nil {
		fmt.Println(err)
	}

	return bytes
}
func (t *Telegram) SendPhoto(chatID string, photoURL, text string) {
	a := strings.SplitN(chatID, "/", 2)
	chatId, err := strconv.ParseInt(a[0], 10, 64)
	if err != nil {
		t.log.Println(err)
	}
	ThreadID := 0
	if len(a) > 1 {
		ThreadID, err = strconv.Atoi(a[1])
		if err != nil {
			t.log.Println(err)
		}
	}
	format := filepath.Ext(photoURL)
	if format != ".jpg" || format != ".jpe" || format != ".png" {
		return
	}

	// Получаем содержимое фотографии по URL
	response, err := http.Get(photoURL)
	if err != nil {
		t.log.Println(err)
	}
	defer response.Body.Close()

	fileName := filepath.Base(photoURL)

	// Создаем временный файл для сохранения фотографии
	tempFile, err := os.Create(fileName)
	if err != nil {
		t.log.Println(err)
	}

	// Копируем содержимое фотографии из ответа HTTP во временный файл
	_, err = io.Copy(tempFile, response.Body)
	if err != nil {
		t.log.Println(err)
	}
	tempFile.Close()

	// Создаем объект сообщения с фотографией
	msg := tgbotapi.NewPhoto(chatId, tgbotapi.FilePath(fileName))
	msg.Caption = text
	msg.MessageThreadID = ThreadID

	_, err = t.t.Send(msg)
	if err != nil {
		t.log.Println(err)
		return
	}
	err = os.Remove(fileName)
	if err != nil {
		t.log.Println(err)
		return
	}
}

//	func (t *Telegram) UploadFile(msg *config.Message, chatid int64, threadid int, parentID int) (string, error) {
//		var media []interface{}
//		for _, f := range msg.Extra["file"] {
//			fi := f.(config.FileInfo)
//			file := tgbotapi.FileBytes{
//				Name:  fi.Name,
//				Bytes: *fi.Data,
//			}
//
//			if b.GetString("MessageFormat") == HTMLFormat {
//				fi.Comment = makeHTML(html.EscapeString(fi.Comment))
//			}
//
//			switch filepath.Ext(fi.Name) {
//			case ".jpg", ".jpe", ".png":
//				pc := tgbotapi.NewInputMediaPhoto(file)
//				if fi.Comment != "" {
//					pc.Caption, pc.ParseMode = TGGetParseMode(b, msg.Username, fi.Comment)
//				}
//				media = append(media, pc)
//			case ".mp4", ".m4v":
//				vc := tgbotapi.NewInputMediaVideo(file)
//				if fi.Comment != "" {
//					vc.Caption, vc.ParseMode = TGGetParseMode(b, msg.Username, fi.Comment)
//				}
//				media = append(media, vc)
//			case ".mp3", ".oga":
//				ac := tgbotapi.NewInputMediaAudio(file)
//				if fi.Comment != "" {
//					ac.Caption, ac.ParseMode = TGGetParseMode(b, msg.Username, fi.Comment)
//				}
//				media = append(media, ac)
//			case ".ogg":
//				voc := tgbotapi.NewVoice(chatid, file)
//				voc.Caption, voc.ParseMode = TGGetParseMode(b, msg.Username, fi.Comment)
//				voc.ReplyToMessageID = parentID
//				res, err := b.c.Send(voc)
//				if err != nil {
//					return "", err
//				}
//				return strconv.Itoa(res.MessageID), nil
//			default:
//				dc := tgbotapi.NewInputMediaDocument(file)
//				if fi.Comment != "" {
//					dc.Caption, dc.ParseMode = TGGetParseMode(b, msg.Username, fi.Comment)
//				}
//				media = append(media, dc)
//			}
//		}
//
//		return b.sendMediaFiles(msg, chatid, threadid, parentID, media)
//	}
