package models

import (
	"kz_bot/internal/storage/memory"
	"time"
)

type InGlobalMessage struct {
	Content string
	Tip     string
	Name    string
	Ds      struct {
		Mesid   string
		Nameid  string
		Guildid string
		Avatar  string
		ChatId  string
		Reply   ReplyWebhookMessage
	}
	Tg struct {
		Mesid  int
		Nameid int64
	}
	Wa struct {
		Nameid string
		Mesid  string
	}
	Config memory.ConfigGlobal
}
type ReplyWebhookMessage struct {
	Text, Username, ChatId, GuildId, Avatar string
	Reply                                   struct {
		TimeMessage time.Time
		Text        string
		Avatar      string
		UserName    string
	}
}
