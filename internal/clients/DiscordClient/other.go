package DiscordClient

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strings"
)

func (d *Discord) CheckAdmin(nameid string, chatid string) bool {
	perms, err := d.s.UserChannelPermissions(nameid, chatid)
	if err != nil {
		d.log.Println("ошибка проверки админ ли ", err)
	}
	if perms&discordgo.PermissionAdministrator != 0 {
		return true
	} else {
		return false
	}
}
func (d *Discord) RoleToIdPing(rolePing, guildid string) string {

	if guildid == "" {
		d.log.Panic("почему то нет гуилд ид")
		panic("почему то нет гуилд ид")
	}
	g, err := d.s.Guild(guildid)
	if err != nil {
		d.log.Println("ошибка получении гильдии при получении роли", err)
	}
	exist, role := d.roleExists(g, rolePing)
	if !exist {
		//создаем роль и возврашаем пинг
		role = d.createRole(rolePing, guildid)
		return role.Mention()
	} else {
		return role.Mention()
	}
}
func (d *Discord) BotName() string { //получаем имя бота
	u, err := d.s.User("@me")
	if err != nil {
		d.log.Println("Ошибка получения имени бота", err)
	}
	return u.Username
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
		go d.DeleteMesageSecond(chatid, mesid, 179)
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

func (d *Discord) dsChatName(chatid, guildid string) string {
	g, err := d.s.Guild(guildid)
	if err != nil {
		d.log.Println("Ошибка проверка имени канала ", err)
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
		d.log.Println("ошибка создании новой роли ", err)
		return nil
	}
	return create
}

func (d *Discord) getLang(chatId, key string) string {
	_, conf := d.storage.Cache.CheckChannelConfigDS(chatId)
	return d.storage.Words.GetWords(conf.Country, key)
}
