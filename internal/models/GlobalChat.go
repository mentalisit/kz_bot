package models

import (
	"kz_bot/internal/storage/memory"
	"time"
)

type InGlobalMessage struct {
	Content string
	Tip     string
	Name    string
	Ds      InGlobalMessageDs
	Tg      struct {
		Mesid  int
		Nameid int64
	}
	Wa struct {
		Nameid string
		Mesid  string
	}
	Config memory.ConfigGlobal
}
type InGlobalMessageDs struct {
	MesId         string
	NameId        string
	Text          string
	Username      string
	ChatId        string
	GuildId       string
	Avatar        string
	TimestampUnix int64
	Reply         struct {
		TimeMessage time.Time
		Text        string
		Avatar      string
		UserName    string
	}
}
type MessageMemory struct {
	Timestamp int64
	Message   []struct {
		MessageId string
		ChatId    string
	}
}
