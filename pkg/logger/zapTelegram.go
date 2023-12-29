package logger

import (
	"bytes"
	"fmt"
	"net/http"
	"strconv"
)

type telegramWriter struct {
	botToken string
	chatID   string
}

func NewTelegramWriter(botToken string, chatID int64) *telegramWriter {
	return &telegramWriter{
		botToken: botToken,
		chatID:   strconv.FormatInt(chatID, 10),
	}
}

func (t *telegramWriter) Write(p []byte) (n int, err error) {
	message := string(p)
	err = sendTelegramMessage(t.botToken, t.chatID, message)
	if err != nil {
		return 0, err
	}
	return len(p), nil
}

func sendTelegramMessage(token, chatID, message string) error {
	apiURL := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", token)

	params := map[string]string{
		"chat_id": chatID,
		"text":    message,
	}

	body := new(bytes.Buffer)
	for key, value := range params {
		body.WriteString(fmt.Sprintf("%s=%s&", key, value))
	}
	bodyStr := body.String()

	resp, err := http.Post(apiURL, "application/x-www-form-urlencoded", bytes.NewBufferString(bodyStr))
	if err != nil {
		return fmt.Errorf("ошибка при отправке запроса: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("неправильный статус ответа: %s", resp.Status)
	}

	return nil
}
