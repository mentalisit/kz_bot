package models

type BridgeMessage struct {
	Text          string
	Sender        string
	Tip           string
	ChatId        string
	MesId         string
	GuildId       string
	TimestampUnix int64
	FileUrl       string
	Avatar        string
	Reply         *BridgeMessageReply
	Config        *BridgeConfig
}

type BridgeMessageReply struct {
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
