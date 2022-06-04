package models

import (
	"github.com/bwmarrin/discordgo"
	"sync"
)

type InMessage struct {
	Mutex         sync.Mutex
	Mtext         string
	Tip           string
	Name          string
	NameMention   string
	Lvlkz, Timekz string
	Ds            Ds
	Tg            Tg
	Config        BotConfig
	Option        Option
}
type Option struct {
	Callback bool
	Edit     bool
	Update   bool
}
type Ds struct {
	Mesid       string
	Nameid      string
	Guildid     string
	Attachments *discordgo.MessageAttachment
}

type Tg struct {
	Mesid  int
	Nameid int64
}

type Configs struct {
	DelMesComplite int
	mesidDsHelp    string
	mesidTgHelp    int
	Primer         string
	Guildid        string
}
type BotConfig struct {
	Type      int
	CorpName  string
	DsChannel string
	TgChannel int64
	WaChannel string
	Config    Configs
}
