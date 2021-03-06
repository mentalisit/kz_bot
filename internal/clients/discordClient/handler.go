package discordClient

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strings"
)

func (d *Discord) cleanChat(m *discordgo.MessageCreate) {
	res := strings.HasPrefix(m.Content, ".")
	if !res { //если нет префикса  то удалить через 3 минуты
		go d.DeleteMesageSecond(m.ChannelID, m.ID, 180)
	}
	if len(m.Attachments) > 0 { //если что-то   то удалить через 3 минуты
		for _, attach := range m.Attachments {
			d.DeleteMesageSecond(m.ChannelID, attach.ID, 180)
		}
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
	g, err := d.d.Guild(guildid)
	if err != nil {
		d.log.Println("Ошибка проверка имени канала ", err)
	}
	chatName := g.Name
	channels, _ := d.d.GuildChannels(guildid)

	for _, r := range channels {
		if r.ID == chatid {
			chatName = chatName + "." + r.Name
			fmt.Println(chatName)
		}
	}
	return chatName
}
