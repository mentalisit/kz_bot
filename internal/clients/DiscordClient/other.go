package DiscordClient

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"kz_bot/pkg/translator"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func (d *Discord) CheckAdmin(nameid string, chatid string) bool {
	perms, err := d.s.UserChannelPermissions(nameid, chatid)
	if err != nil {
		d.log.ErrorErr(err)
	}
	if perms&discordgo.PermissionAdministrator != 0 {
		return true
	} else {
		return false
	}
}
func (d *Discord) RoleToIdPing(rolePing, guildid string) (string, error) {

	if guildid == "" {
		d.log.Panic("почему то нет гуилд ид")
	}
	g, err := d.s.Guild(guildid)
	if err != nil {
		ge, err1 := d.s.Guild(guildid)
		if err1 != nil {
			d.log.Error(err1.Error())
			return rolePing, err1
		}
		g = ge
	}
	exist, role := d.roleExists(g, rolePing)
	if !exist {
		//создаем роль и возврашаем пинг
		role = d.createRole(rolePing, guildid)
		return role.Mention(), nil
	} else {
		return role.Mention(), nil
	}
}
func (d *Discord) TextToRoleRsPing(rolePing, guildid string) string {

	if guildid == "" {
		d.log.Panic("почему то нет гуилд ид")
		panic("почему то нет гуилд ид")
	}
	g, err := d.s.Guild(guildid)
	if err != nil {
		d.log.ErrorErr(err)
	}
	exist, role := d.roleExists(g, rolePing)
	if !exist {
		return fmt.Sprintf("`роль %s не найдена в %s`", rolePing, g.Name)
	} else {
		return role.Mention()
	}
}
func (d *Discord) DMchannel(AuthorID string) (chatidDM string) {
	create, err := d.s.UserChannelCreate(AuthorID)
	if err != nil {
		return ""
	}
	chatidDM = create.ID
	return chatidDM
}
func (d *Discord) CleanChat(chatid, mesid, text string) {
	res := strings.HasPrefix(text, ".")
	if !res { //если нет префикса  то удалить через 3 минуты
		go d.DeleteMesageSecond(chatid, mesid, 180)
	}
}

// получаем есть ли роль и саму роль
func (d *Discord) roleExists(g *discordgo.Guild, nameRoles string) (bool, *discordgo.Role) {
	nameRoles = strings.ToLower(nameRoles)

	for _, role := range g.Roles {
		if role.Name == "@everyone" {
			continue
		}
		if strings.ToLower(role.Name) == nameRoles {
			return true, role
		}
	}
	return false, nil
}

func (d *Discord) GuildChatName(chatid, guildid string) string {
	g, err := d.s.Guild(guildid)
	if err != nil {
		d.log.ErrorErr(err)
	}
	chatName := g.Name
	channels, _ := d.s.GuildChannels(guildid)

	for _, r := range channels {
		if r.ID == chatid {
			chatName = chatName + "." + r.Name
			fmt.Println(chatName)
		}
	}
	return chatName
}

func (d *Discord) createRole(rolPing, guildid string) *discordgo.Role {
	t := true
	perm := int64(37080064)
	create, err := d.s.GuildRoleCreate(guildid, &discordgo.RoleParams{
		Name:        rolPing,
		Permissions: &perm,
		Mentionable: &t,
	})
	if err != nil {
		d.log.ErrorErr(err)
		return nil
	}
	return create
}

func (d *Discord) getLang(chatId, key string) string {
	_, conf := d.CheckChannelConfigDS(chatId)
	return d.storage.Words.GetWords(conf.Country, key)
}

func (d *Discord) CleanOldMessageChannel(chatId, lim string) {
	limit, _ := strconv.Atoi(lim)
	if limit == 0 {
		return
	}
	messages, err := d.s.ChannelMessages(chatId, limit, "", "", "")
	if err != nil {
		d.log.ErrorErr(err)
		return
	}
	for _, message := range messages {
		if message.WebhookID == "" {
			if !message.Author.Bot {
				d.DeleteMessage(chatId, message.ID)
				continue
			}
			if !strings.HasPrefix(message.Content, ".") {
				d.DeleteMessage(chatId, message.ID)
				continue
			}
		}
	}
}

func (d *Discord) avatar(m *discordgo.MessageCreate) bool {
	str, ok := strings.CutPrefix(m.Content, ". ")
	if ok {
		arg := strings.Split(strings.ToLower(str), " ")
		if len(arg) == 2 {
			if arg[0] == "ава" {
				mentionIds := userMentionRE.FindAllStringSubmatch(arg[1], -1)
				if len(mentionIds) > 0 {
					members, err := d.s.GuildMembers(m.GuildID, "", 999)
					if err != nil {
						d.log.ErrorErr(err)
					}
					for _, member := range members {
						if member.User.ID == mentionIds[0][1] {
							aname := m.Author.Username
							if m.Member.Nick != "" {
								aname = m.Member.Nick
							}
							name := member.User.Username
							if member.Nick != "" {
								name = member.Nick
							}
							em := &discordgo.MessageEmbed{
								Title: fmt.Sprintf("Аватар %s по запросу %s", name, aname),
								Color: 14232643,
								Image: &discordgo.MessageEmbedImage{
									URL: member.AvatarURL("2048"),
								},
								Author: nil,
							}
							embed, err := d.s.ChannelMessageSendEmbed(m.ChannelID, em)
							if err != nil {
								fmt.Println(err.Error())
								return false
							}
							go d.DeleteMesageSecond(m.ChannelID, embed.ID, 183)
							go d.DeleteMesageSecond(m.ChannelID, m.ID, 30)
							return true
						}
					}
				}
			}
		}
	}
	return false
}

func (d *Discord) getAuthorName(m *discordgo.MessageCreate) string {
	username := m.Author.Username
	if m.Member != nil && m.Member.Nick != "" {
		username = m.Member.Nick
	}
	return username
}

func (d *Discord) latinOrNot(m *discordgo.MessageCreate) {
	cyrillicPattern := regexp.MustCompile(`[а-яА-ЯґҐєЄіІїЇ]`)
	if len(m.Content) >= 1 {
		if !cyrillicPattern.MatchString(m.Content) {
			channel, err := d.s.Channel(m.ChannelID)
			if err != nil {
				return
			}
			gostPattern := regexp.MustCompile(`гост`)

			if gostPattern.MatchString(channel.Name) {
				//возможно нужно доп условие
				go func() {
					text2 := translator.TranslateAnswer(m.Content, "ru")
					mes := d.SendWebhook(text2, m.Author.Username, m.ChannelID, m.GuildID, m.Author.AvatarURL("128"))
					d.DeleteMesageSecond(m.ChannelID, mes, 90)
				}()
			}
		}
	}
}
func (d *Discord) transtale(m *discordgo.Message, lang string, r *discordgo.MessageReactionAdd) {
	text2 := translator.TranslateAnswer(m.Content, lang)
	go func() {
		time.Sleep(30 * time.Second)
		err := d.s.MessageReactionRemove(r.ChannelID, r.MessageID, r.Emoji.Name, r.UserID)
		if err != nil {
			fmt.Println("Ошибка удаления реакции", err)
		}
	}()
	mes := d.SendWebhook(text2, m.Author.Username, m.ChannelID, m.GuildID, m.Author.AvatarURL("128"))
	d.DeleteMesageSecond(m.ChannelID, mes, 90)
}
