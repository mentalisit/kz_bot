package discordClient

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"time"
)

const (
	emOK      = "✅"
	emCancel  = "❎"
	emRsStart = "🚀"
	emPl30    = "⌛"
	emPlus    = "➕"
	emMinus   = "➖"
)

var mesContentNil string

type Ds struct {
	d discordgo.Session
}

type DiscordInterface interface {
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

func EmbedDS(name1, name2, name3, name4, lvlkz string, numkz int) discordgo.MessageEmbed {
	Embeds := &discordgo.MessageEmbed{
		Author: &discordgo.MessageEmbedAuthor{},
		Color:  16711680,
		Description: fmt.Sprintf("Желающие:👇 |  <:rs:918545444425072671> на %s (%d) ", lvlkz, numkz) +
			fmt.Sprintf(
				"\n1️⃣ %s "+
					"\n2️⃣ %s "+
					"\n3️⃣ %s "+
					"\n4️⃣ %s "+
					"\n", name1, name2, name3, name4),

		Fields: []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{
				Name:   fmt.Sprintf(" %s для добавления в очередь\n%s для выхода из очереди\n%s принудительный старт", emOK, emCancel, emRsStart),
				Value:  "Данные обновлены: ",
				Inline: true,
			}},
		Timestamp: time.Now().Format(time.RFC3339), // ТЕКУЩЕЕ ВРЕМЯ ДИСКОРДА
		Title:     "ОЧЕРЕДЬ КЗ  ",
	}
	return *Embeds
}

func (d *Ds) CheckAdmin(nameid string, chatid string) bool {
	perms, err := d.d.UserChannelPermissions(nameid, chatid)
	if err != nil {
		fmt.Println("ошибка проверки админ ли ", err)
	}
	if perms&discordgo.PermissionAdministrator != 0 {
		//logrus.Println("админ")
		return true
	} else {
		//logrus.Println("не админ")
		return false
	}
}
func (d *Ds) AddEnojiRsQueue(chatid, mesid string) {
	d.d.MessageReactionAdd(chatid, mesid, emOK)
	d.d.MessageReactionAdd(chatid, mesid, emCancel)
	d.d.MessageReactionAdd(chatid, mesid, emRsStart)
	d.d.MessageReactionAdd(chatid, mesid, emPl30)

}
func (d *Ds) DeleteMessage(chatid, mesid string) {
	d.d.ChannelMessageDelete(chatid, mesid)
}
func (d *Ds) SendChannelDelSecond(chatid, text string, second int) {
	message, err := d.d.ChannelMessageSend(chatid, text)
	if err != nil {
		fmt.Println("ошибка отправки сообщения SendChannelDelSecond", err)
	}
	go func() {
		time.Sleep(time.Duration(second) * time.Second)
		d.d.ChannelMessageDelete(chatid, message.ID)
	}()

}
func (d *Ds) RoleToIdPing(rolePing, guildid string) string {
	//создаю переменную
	rolPing := "кз" + rolePing // добавляю буквы
	g, err := d.d.Guild(guildid)
	if err != nil {
		fmt.Println("ошибка получении гильдии при получении роли", err)
	}
	exist, role := d.roleExists(g, rolPing)
	if !exist {
		//создаем роль и возврашаем пинг
		newRole, err := d.d.GuildRoleCreate(guildid)
		if err != nil {
			fmt.Println("ошибка создании новой роли ", err)
		}
		role, err = d.d.GuildRoleEdit(guildid, newRole.ID, rolPing, newRole.Color, newRole.Hoist, 37080064, true)
		if err != nil {
			fmt.Println("Ошибка изменения новой роли", err)
			err = d.d.GuildRoleDelete(guildid, newRole.ID)
			if err != nil {
				fmt.Println("ошибка удаления новой роли ", err)
			}
		}
		return role.Mention()
	} else {
		return role.Mention()
	}

	r, err := d.d.GuildRoles(guildid)
	if err != nil {
		fmt.Println("Ошибка чтения ролей ", err)
	}
	l := len(r) // количество ролей на сервере
	i := 0
	for i < l { //ищу роли в цикле
		if r[i].Name == rolPing {
			//pingId = r[i].ID
			return r[i].Mention()
			//return "<@&" + pingId + ">" // возвращаю пинг роли
		} else {
			i = i + 1 // продолжаю перебор
		}
	}
	return "(роль не найдена)" // если не нашол нужной роли
}
func (d *Ds) DeleteMesageSecond(chatid, mesid string, second int) {
	if second > 60 {
		//timerInsert(mesid, chatid, 0, 0, second)
	} else {
		go func() {
			time.Sleep(time.Duration(second) * time.Second)
			d.d.ChannelMessageDelete(chatid, mesid)
		}()
	}

}
func (d *Ds) EditComplex(dsmesid, dschatid string, Embeds *discordgo.MessageEmbed) {
	a := &discordgo.MessageEdit{
		Content: &mesContentNil,
		Embed:   Embeds,
		ID:      dsmesid,
		Channel: dschatid,
	}
	_, err := d.d.ChannelMessageEditComplex(a)
	if err != nil {
		fmt.Println("Ошибка редактирования комплексного сообщения ", err)
	}
}
func (d *Ds) BotName() string { //получаем имя бота
	u, _ := d.d.User("@me")
	return u.Username
}
func (d *Ds) SendComplexContent(chatid, text string) string { //отправка текста комплексного сообщения
	mesCompl, err := d.d.ChannelMessageSendComplex(chatid, &discordgo.MessageSend{
		Content: text})
	if err != nil {
		fmt.Println("Ошибка отправки комплексного сообщения ", err)
	}
	return mesCompl.ID
}
func (d *Ds) Send(chatid, text string) string { //отправка текста
	message, err := d.d.ChannelMessageSend(chatid, text)
	if err != nil {
		fmt.Println("ошибка отправки текста ", err)
	}
	return message.ID
}
