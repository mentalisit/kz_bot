package models

import "kz_bot/internal/storage/memory"

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
