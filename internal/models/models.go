package models

var ChDs = make(chan InMessage, 10)
var ChTg = make(chan InMessage, 10)
var ChWa = make(chan InMessage, 10)

type InMessage struct {
	Mtext         string
	Tip           string
	Name          string
	NameMention   string
	Lvlkz, Timekz string
	Ds            struct {
		Mesid   string
		Nameid  string
		Guildid string
	}
	Tg struct {
		Mesid  int
		Nameid int64
	}
	Wa struct {
		Nameid string
	}
	Config BotConfig
	Option struct {
		Callback bool
		Edit     bool
		Update   bool
		Queue    bool
	}
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

type Timer struct {
	Id       int
	Dsmesid  string
	Dschatid string
	Tgmesid  int
	Tgchatid int64
	Timed    int
}

type Names struct {
	Name1 string
	Name2 string
	Name3 string
	Name4 string
}
