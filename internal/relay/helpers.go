package relay

import (
	"regexp"
	"strings"
)

var messageTextAuthor [2]string

// проверка на повторное сообщение
func (r *Relay) checkingForIdenticalMessage() bool {
	if messageTextAuthor[0] == r.in.Text && messageTextAuthor[1] == r.in.Author {
		go r.client.Ds.DeleteMessage(r.in.Ds.ChatId, r.in.Ds.MesId)
		return true
	}
	messageTextAuthor[0] = r.in.Text
	messageTextAuthor[1] = r.in.Author
	return false
}

// запрещаем многие ссылки
func filterMessageLinks(input string) string {
	// Регулярное выражение для поиска ссылок
	re := regexp.MustCompile(`(https?://[^\s]+)`)
	// Список разрешенных ссылок
	allowedLinks := []string{
		"https://t.me/",
		"https://discord.com/channels/",
		"https://cdn.discordapp.com/attachments/",
		"https://discord.gg/",
		"https://userxinos.github.io/",
	}
	// Запрещенная ссылка
	forbiddenLink := "запрещенная ссылка"
	// Заменяем все ссылки, кроме разрешенных, на запрещенную ссылку
	output := re.ReplaceAllStringFunc(input, func(link string) string {
		for _, allowedLink := range allowedLinks {
			if strings.HasPrefix(link, allowedLink) {
				return link
			}
		}
		return forbiddenLink
	})
	return output
}

// поиск и замена @&rs на пинг
func (r *Relay) replaceTextMentionRsRole(input, guildId string) string {
	re := regexp.MustCompile(`@&rs([4-9]|1[0-2])`)
	output := re.ReplaceAllStringFunc(input, func(s string) string {
		return r.client.Ds.TextToRoleRsPing(s[2:], guildId)
	})
	return output
}
