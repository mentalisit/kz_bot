package models

type Message struct {
	Text        string
	Sender      string
	Avatar      string
	ChannelType int
	Corporation string
	Command     string
	Messager    string
}
type Corporation struct {
	Id        int
	Corp      string
	DsChat    string
	DsChatWS1 string
	DsChatWS2 string
	GuildId   string
	TgChat    int64
	TgChatWS1 int64
}
