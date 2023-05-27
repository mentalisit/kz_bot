package models

import "time"

type RelayMessage struct {
	Text   string
	Tip    string
	Author string
	Ds     *RelayMessageDs
	Tg     *RelayMessageTg
	Config *RelayConfig
}
type RelayMessageDs struct {
	ChatId        string
	MesId         string
	Avatar        string
	GuildId       string
	TimestampUnix int64
	Reply         struct {
		TimeMessage time.Time
		Text        string
		Avatar      string
		UserName    string
	}
}
type RelayMessageTg struct {
	ChatId        int64
	MesId         int
	Avatar        string
	GuildId       int64
	TimestampUnix int64
	GroupName     string
	Reply         struct {
		TimeMessage int64
		Text        string
		Avatar      string
		UserName    string
	}
}
type RelayConfig struct {
	Id         int
	RelayName  string
	RelayAlias string
	GuildName  string
	DsChannel  string
	TgChannel  int64
	WaChannel  string
	GuildId    string
	Country    string
	Prefix     string
}

type RelayMessageMemory struct {
	Timestamp int64
	RelayName string
	MessageDs []struct {
		MessageId string
		ChatId    string
	}
	MessageTg []struct {
		MessageId int
		ChatId    int64
	}
}
