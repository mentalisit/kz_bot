package discordClient

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strings"
)

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
