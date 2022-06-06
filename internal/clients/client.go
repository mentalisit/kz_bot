package clients

import "github.com/bwmarrin/discordgo"

type Discord interface {
	Send(chatid, text string) string
	SendChannelDelSecond(chatid, text string, second int)
	SendComplexContent(chatid, text string) string
	EditComplex(dsmesid, dschatid string, Embeds *discordgo.MessageEmbed)
	DeleteMesageSecond(chatid, mesid string, second int)
	DeleteMessage(chatid, mesid string)
	RoleToIdPing(rolePing, guildid string) string
	AddEnojiRsQueue(chatid, mesid string)
	CheckAdmin(nameid string, chatid string) bool
	BotName() string
}
type Telega interface {
	SendChannel(chatid int64, text string) int
	SendChannelDelSecond(chatid int64, text string, second int)
	SendEmded(lvlkz string, chatid int64, text string) int
	SendEmbedTime(chatid int64, text string) int
	EditText(chatid int64, editMesId int, textEdit string)
	EditMessageTextKey(chatid int64, editMesId int, textEdit string, lvlkz string)
	DelMessage(chatid int64, idSendMessage int)
	DelMessageSecond(chatid int64, idSendMessage int, second int)
	CheckAdminTg(chatid int64, name string) bool
	RemoveDuplicateElementInt(mesididid []int) []int
	ChatName(chatid int64) string
	BotName() string
}
