package logger

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type discordWriter struct {
	webhook string
}

func NewDiscordWriter(webhook string) *discordWriter {
	return &discordWriter{webhook: webhook}
}

func (t *discordWriter) Write(p []byte) (n int, err error) {
	message := string(p)
	sendDiscordMessage(t.webhook, message)

	return len(p), nil
}

func sendDiscordMessage(webhookURL, text string) {
	message := map[string]interface{}{
		"content":    fmt.Sprintf("```%s```", text),
		"username":   "Logger",
		"avatar_url": "https://e7.pngegg.com/pngimages/836/966/png-clipart-go-programming-language-computer-programming-others-baltimore-web-application-thumbnail.png",
	}

	jsonBody, err := json.Marshal(message)
	if err != nil {
		fmt.Println(err)
		return
	}

	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
}
