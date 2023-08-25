package models

//	type Message struct {
//		Text        string
//		Sender      string
//		Avatar      string
//		ChannelType int
//		Corporation string
//		Command     string
//		Messager    string
//	}
type MessageHadesClient struct {
	Text                  string
	Sender                string
	Avatar                string
	ChannelType           int
	Corporation           string
	Command               string
	Messager              string
	MessageId             int64
	SolarSystemId         int64
	MatchedStarSupernovas int
}

//type Corporation struct {
//	Id        int
//	Corp      string
//	DsChat    string
//	DsChatWS1 string
//	DsChatWS2 string
//	GuildId   string
//	TgChat    string
//	TgChatWS1 string
//}

type MessageHades struct {
	Text        string
	Sender      string
	Avatar      string
	ChannelType int
	Corporation string
	Command     string
	Messager    string
	MessageId   string
}
type MessageHadesDs struct {
	MessageId string
}
type MessageHadesTg struct {
	MessageId int
}
type CorporationHadesClient struct {
	Corp      string
	DsChat    string
	DsChatWS1 string
	DsChatWS2 string
	GuildId   string
	TgChat    string
	TgChatWS1 string
	//ThreadID  int
}
type GameMessageId struct {
	MessageId int64  `bson:"MessageId"`
	CorpName  string `bson:"CorpName"`
}
type GameMessageIdWs1 struct {
	MessageId int64  `bson:"MessageId"`
	CorpName  string `bson:"CorpName"`
	StarId    int64  `bson:"StarId"`
}
type AllianceMember struct {
	CorpName string `bson:"CorpName"`
	UserName string `bson:"UserName"`
	Rang     int    `bson:"Rang"`
}
