package DiscordClient

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"regexp"
	"strings"
)

//nujno sdelat lang

func (d *Discord) AccesChatDS(m *discordgo.MessageCreate) {
	res := strings.HasPrefix(m.Content, ".")
	if res == true && m.Content == ".add" {
		go d.DeleteMesageSecond(m.ChannelID, m.ID, 10)
		d.accessAddChannelDs(m.ChannelID, m.GuildID)
	} else if res == true && m.Content == ".del" {
		go d.DeleteMesageSecond(m.ChannelID, m.ID, 10)
		d.accessDelChannelDs(m.ChannelID, m.GuildID)
	}
	if res {
		if d.CleanOldMessage(m) {
			return
		}
		//d.accessAddGlobalDs(m)
	}
}

func (d *Discord) accessAddChannelDs(chatid, guildid string) { // внесение в дб и добавление в масив
	ok, _ := d.CheckChannelConfigDS(chatid)
	if ok {
		go d.SendChannelDelSecond(chatid, "Я уже могу работать на вашем канале\n"+
			"повторная активация не требуется.\nнапиши Справка", 30)
	} else {
		chatName := d.GuildChatName(chatid, guildid)
		d.log.Println("новая активация корпорации ", chatName)
		d.AddDsCorpConfig(chatName, chatid, guildid)
		go d.SendChannelDelSecond(chatid, "Спасибо за активацию.", 10)
		d.HelpChannelUpdate(chatid)
	}
}
func (d *Discord) accessDelChannelDs(chatid, guildid string) { //удаление с бд и масива для блокировки
	ok, config := d.CheckChannelConfigDS(chatid)
	d.DeleteMessage(chatid, config.MesidDsHelp)
	if !ok {
		go d.SendChannelDelSecond(chatid, "ваш канал и так не подключен к логике бота ", 60)
	} else {
		d.DeleteDs(chatid)
		d.log.Println("отключение корпорации ", d.GuildChatName(chatid, guildid))
		//d.storage.Cache.ReloadConfig()
		d.storage.CorpsConfig.ReadCorps()
		go d.SendChannelDelSecond(chatid, "вы отключили мои возможности", 60)
	}
}

//func (d *Discord) accessAddGlobalDs(m *discordgo.MessageCreate) {
//	str, ok := strings.CutPrefix(m.Content, ".")
//	if ok {
//		arr := strings.Split(str, " ")
//		if arr[0] == "AddGlobalChat" {
//			good, _ := d.storage.CacheGlobal.CheckChannelConfigDS(m.ChannelID)
//			if good {
//				d.SendChannelDelSecond(m.ChannelID, "Этот чат уже подключен", 10)
//			} else {
//				guild, _ := d.s.Guild(m.GuildID)
//				err := d.storage.CorpsConfig.AddGlobalDsCorp(context.Background(), guild.Name, m.ChannelID, m.GuildID)
//				if err != nil {
//					d.log.Println(err)
//					return
//				} else {
//					d.SendChannelDelSecond(m.ChannelID, "GlobalChat активирован", 10)
//					m.Content = fmt.Sprintf("%s присоеденилась к реле Rs_bot", guild.Name)
//					d.logicMixGlobal(m)
//				}
//
//			}
//		}
//		if len(arr) > 1 {
//			if arr[0] == "block" {
//				mentionIds := userMentionRE.FindAllStringSubmatch(arr[1], -1)
//				for _, match := range mentionIds {
//					user := d.getUserById(match[1], m.GuildID)
//					memory.BlackListNamesId = append(memory.BlackListNamesId, user.User.ID)
//					marshal, err := json.Marshal(memory.BlackListNamesId)
//					if err != nil {
//						return
//					}
//					d.storage.CorpsConfig.UpdateJsonBlackList(marshal)
//				}
//			} else if arr[0] == "unblock" {
//				var newList []string
//				mentionIds := userMentionRE.FindAllStringSubmatch(arr[1], -1)
//				for _, match := range mentionIds {
//					user := d.getUserById(match[1], m.GuildID)
//					if user.User.ID == m.Author.ID {
//						return
//					}
//					for _, s := range memory.BlackListNamesId {
//						if s != user.User.ID {
//							newList = append(newList, s)
//						}
//					}
//				}
//				if len(newList) > 0 {
//					memory.BlackListNamesId = newList
//					marshal, err := json.Marshal(newList)
//					if err != nil {
//						return
//					}
//					d.storage.CorpsConfig.UpdateJsonBlackList(marshal)
//				}
//			}
//		}
//	}
//}

func (d *Discord) CleanOldMessage(m *discordgo.MessageCreate) bool {
	re := regexp.MustCompile(`^\.очистка (\d{1,2}|100)`)
	matches := re.FindStringSubmatch(m.Content)
	if len(matches) > 0 {
		fmt.Println("limitMessage " + matches[1])
		d.CleanOldMessageChannel(m.ChannelID, matches[1])
		return true
	}
	return false
}
