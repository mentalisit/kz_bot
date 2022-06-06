package models

import (
	"sync"
)

var ChDs = make(chan InMessage, 10)
var ChTg = make(chan InMessage, 10)

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
	Mesid   string
	Nameid  string
	Guildid string
}

type Tg struct {
	Mesid  int
	Nameid int64
}

type Configs struct {
	DelMesComplite int
	MesidDsHelp    string
	MesidTgHelp    int
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

type TableConfig struct {
	Id             int
	Corpname       string
	Dschannel      string
	Tgchannel      int64
	Wachannel      string
	Mesiddshelp    string
	Mesidtghelp    int
	Delmescomplite int
	Guildid        string
}
