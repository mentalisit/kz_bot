package DiscordClient

import (
	"context"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"time"
)

//lang ok

func (d *Discord) AddEnojiRsQueue(chatid, mesid string) {
	err := d.s.MessageReactionAdd(chatid, mesid, emOK)
	if err != nil {
		d.log.Println("Ошибка добавления реакции на сообщения "+emOK, err)
	}
	err = d.s.MessageReactionAdd(chatid, mesid, emCancel)
	if err != nil {
		d.log.Println("Ошибка добавления реакции на сообщения "+emCancel, err)
	}
	err = d.s.MessageReactionAdd(chatid, mesid, emRsStart)
	if err != nil {
		d.log.Println("Ошибка добавления реакции на сообщения "+emRsStart, err)
	}
	err = d.s.MessageReactionAdd(chatid, mesid, emPl30)
	if err != nil {
		d.log.Println("Ошибка добавления реакции на сообщения "+emPl30, err)
	}

}
func (d *Discord) DeleteMessage(chatid, mesid string) {
	_ = d.s.ChannelMessageDelete(chatid, mesid)
}
func (d *Discord) DeleteMesageSecond(chatid, mesid string, second int) {
	if second > 60 {
		ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
		defer cancel()
		d.storage.Timers.TimerInsert(ctx, mesid, chatid, 0, 0, second)
	} else {
		go func() {
			time.Sleep(time.Duration(second) * time.Second)
			err := d.s.ChannelMessageDelete(chatid, mesid)
			if err != nil {
				fmt.Println("Ошибка удаления сообщения дискорда ", chatid, mesid, second)
			}
		}()
	}
}
func (d *Discord) EditComplex(dsmesid, dschatid string, Embeds discordgo.MessageEmbed) error {
	_, err := d.s.ChannelMessageEditComplex(&discordgo.MessageEdit{
		Content: &mesContentNil,
		Embed:   &Embeds,
		ID:      dsmesid,
		Channel: dschatid,
	})
	if err != nil {
		return err
	}
	return nil
}
func (d *Discord) Subscribe(nameid, argRoles, guildid string) int {
	g, err := d.s.State.Guild(guildid)
	if err != nil {
		d.log.Println("Ошибка запроса стате.гуилд,читаю гуилд", err)
		g, err = d.s.Guild(guildid)
		if err != nil {
			d.log.Println("Ошибка чтения гуилд ... паниковать ", err)
		}
	}

	exist, role := d.roleExists(g, argRoles)

	if !exist { //если нет роли
		role = d.createRole(argRoles, guildid)
	}

	member, err := d.s.GuildMember(guildid, nameid)
	if err != nil {
		d.log.Println("Ошибка чтения участников гуилд", err)
	}
	var subscribe int = 0
	if exist {
		for _, _role := range member.Roles {
			if _role == role.ID {
				subscribe = 1
			}
		}
	}

	err = d.s.GuildMemberRoleAdd(guildid, nameid, role.ID)
	if err != nil {
		d.log.Println("Ошибка выдачи роли ", err)
		subscribe = 2
	}

	return subscribe
}
func (d *Discord) Unsubscribe(nameid, argRoles, guildid string) int {
	var unsubscribe int = 0
	g, err := d.s.State.Guild(guildid)
	if err != nil {
		d.log.Println("Ошибка запроса стате.гуилд,читаю гуилд", err)
		g, err = d.s.Guild(guildid)
		if err != nil {
			d.log.Println("Ошибка чтения гуилд ... паниковать ", err)
		}
	}

	exist, role := d.roleExists(g, argRoles)
	if !exist { //если нет роли
		unsubscribe = 1
	}

	member, err := d.s.GuildMember(guildid, nameid)
	if err != nil {
		d.log.Println("Ошибка чтения участников гуилд", err)
	}
	if exist {
		for _, _role := range member.Roles {
			if _role == role.ID {
				unsubscribe = 2
			}
		}
	}
	if unsubscribe == 2 {
		err = d.s.GuildMemberRoleRemove(guildid, nameid, role.ID)
		if err != nil {
			d.log.Println("Ошибка снятия роли ", err)
			unsubscribe = 3
		}
	}

	return unsubscribe
}
func (d *Discord) EditMessage(chatID, messageID, content string) {
	_, err := d.s.ChannelMessageEdit(chatID, messageID, content)
	if err != nil {
		d.log.Println("Ошибка изменения текса сообщения ", err)
	}
}
func (d *Discord) EmbedDS(mapa map[string]string, numkz int) discordgo.MessageEmbed {
	return discordgo.MessageEmbed{
		Author: &discordgo.MessageEmbedAuthor{},
		Color:  16711680,
		Description: fmt.Sprintf("👇 %s <:rs:918545444425072671> %s (%d) ",
			d.storage.Words.GetWords(mapa["lang"], "jelaushieNa"), mapa["lvlkz"], numkz) +
			fmt.Sprintf("\n1️⃣ %s \n2️⃣ %s \n3️⃣ %s \n4️⃣ %s \n",
				mapa["name1"], mapa["name2"], mapa["name3"], mapa["name4"]),

		Fields: []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{
				Name: fmt.Sprintf(" %s %s\n%s %s\n%s %s",
					emOK, d.storage.Words.GetWords(mapa["lang"], "DlyaDobavleniya"),
					emCancel, d.storage.Words.GetWords(mapa["lang"], "DlyaVihodaIz"),
					emRsStart, d.storage.Words.GetWords(mapa["lang"], "prinuditelniStart")),
				Value:  d.storage.Words.GetWords(mapa["lang"], "DannieObnovleni") + ": ",
				Inline: true,
			}},
		Timestamp: time.Now().Format(time.RFC3339), // ТЕКУЩЕЕ ВРЕМЯ ДИСКОРДА
		Title:     d.storage.Words.GetWords(mapa["lang"], "ocheredKz"),
	}
}
