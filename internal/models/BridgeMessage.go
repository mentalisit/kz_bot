package models

type BridgeMessage struct {
	Text    string
	Sender  string
	Tip     string
	Ds      *BridgeMessageDs
	Tg      *BridgeMessageTg
	FileUrl string
	Config  *BridgeConfig
}
type BridgeMessageDs struct {
	ChatId        string
	MesId         string
	Avatar        string
	GuildId       string
	TimestampUnix int64
	Reply         *ReplyDs
}
type ReplyDs struct {
	TimeMessage int64
	Text        string
	Avatar      string
	UserName    string
}
type BridgeMessageTg struct {
	ChatId        string
	MesId         int
	Avatar        string
	TimestampUnix int64
	GroupName     string
	Reply         *ReplyTg
}
type ReplyTg struct {
	TimeMessage int64
	Text        string
	Avatar      string
	UserName    string
}

type BridgeConfig struct {
	Id                int
	NameRelay         string
	HostRelay         string
	Role              []string
	ChannelDs         []BridgeConfigDs
	ChannelTg         []BridgeConfigTg
	ForbiddenPrefixes []string
}
type BridgeConfigDs struct {
	ChannelId       string
	GuildId         string
	CorpChannelName string
	AliasName       string
	MappingRoles    map[string]string
}
type BridgeConfigTg struct {
	ChannelId string
	//ThreadID        int
	CorpChannelName string
	AliasName       string
	MappingRoles    map[string]string
}
