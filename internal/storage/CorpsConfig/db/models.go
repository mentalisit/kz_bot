package db

type ConfigCorp struct {
	Id             int
	CorpName       string
	DsChannel      string
	TgChannel      int64
	WaChannel      string
	MesidDsHelp    string
	MesidTgHelp    int
	DelMesComplite int
	GuildId        string
	Country        string
}
