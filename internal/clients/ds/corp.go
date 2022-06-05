package ds

/*
import (
	"github.com/bwmarrin/discordgo"
	corpsConfig "kz_bot/internal/clients/corpConfig"
	"strings"
)

func (d Ds) AccesChatDS(m *discordgo.MessageCreate) {
	res := strings.HasPrefix(m.Content, ".")
	if res == true && m.Content == ".add" {
		go d.DeleteMesageSecond(m.ChannelID, m.ID, 10)
		d.accessAddChannelDs(m.ChannelID, m.GuildID)
	} else if res == true && m.Content == ".del" {
		go d.DeleteMesageSecond(m.ChannelID, m.ID, 10)
		accessDelChannelDs(m.ChannelID)
	}
}

func (d Ds) accessAddChannelDs(chatid, guildid string) { // внесение в дб и добавление в масив
	c := corpsConfig.CorpConfig{}
	ok, _ := c.CheckChannelConfigDS(chatid)
	if ok {
		go d.SendChannelDelSecond(chatid, "Я уже могу работать на вашем канале\n"+
			"повторная активация не требуется.\nнапиши Справка1", 30)
	} else {
		chatName := dsChatName(guildid)
		insertConfig := `INSERT INTO config (corpname,dschannel,tgchannel,wachannel,mesiddshelp,mesidtghelp,delmescomplite,guildid) VALUES (?,?,?,?,?,?,?,?)`
		statement, err := db.Prepare(insertConfig)
		if err != nil {
			logrus.Println(err)
		}
		_, err = statement.Exec(chatName, chatid, 0, "", "", 0, 0, guildid)
		if err != nil {
			logrus.Println(err.Error())
		}
		//db.Close()
		addCorp(chatName, chatid, 0, "", 1, "", 0, guildid)
		go dsSendChannelDel1m(chatid, "Спасибо за активацию.\nпиши Справка1")
	}
}
func accessDelChannelDs(chatid string) { //удаление с бд и масива для блокировки
	ok, _ := checkChannelConfigDS(chatid)
	if !ok {
		go dsSendChannelDel1m(chatid, "ваш канал и так не подключен к логике бота ")
	} else {
		_, err := db.Exec("delete from config where dschannel = ? ", chatid)
		if err != nil {
			logrus.Println(err)
		}
		*P = *New()
		readBotConfig()
		go dsSendChannelDel1m(chatid, "вы отключили мои возможности")
	}
}


*/
