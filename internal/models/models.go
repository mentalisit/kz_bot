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
type Users struct {
	User1 Sborkz
	User2 Sborkz
	User3 Sborkz
	User4 Sborkz
}
type Sborkz struct {
	Id          int
	Corpname    string
	Name        string
	Mention     string
	Tip         string
	Dsmesid     string
	Tgmesid     int
	Wamesid     string
	Time        string
	Date        string
	Lvlkz       string
	Numkzn      int
	Numberkz    int
	Numberevent int
	Eventpoints int
	Active      int
	Timedown    int
}

type EmodjiUser struct {
	Id                            int
	Tip, Name, Em1, Em2, Em3, Em4 string
}
