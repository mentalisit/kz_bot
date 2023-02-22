package models

import (
	"kz_bot/internal/storage/memory"
)

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
		Avatar  string
	}
	Tg struct {
		Mesid  int
		Nameid int64
	}
	Wa struct {
		Nameid string
		Mesid  string
	}
	Config memory.CorpporationConfig
	Option Option
}

type Option struct {
	Reaction bool
	InClient bool
	Queue    bool
	Pl30     bool
	MinusMin bool
	Edit     bool
	Update   bool
	Elsetrue bool
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
type Names struct {
	Name1 string
	Name2 string
	Name3 string
	Name4 string
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
