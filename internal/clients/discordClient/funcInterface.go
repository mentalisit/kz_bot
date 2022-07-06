package discordClient

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
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

type Ds interface {
	Send(chatid, text string) string
	SendChannelDelSecond(chatid, text string, second int)
	SendComplexContent(chatid, text string) string
	SendEmbedText(chatid, title, text string) *discordgo.Message
	SendComplex(chatid string, embeds discordgo.MessageEmbed) string
	EditComplex(dsmesid, dschatid string, Embeds discordgo.MessageEmbed)
	DeleteMesageSecond(chatid, mesid string, second int)
	DeleteMessage(chatid, mesid string)
	RoleToIdPing(rolePing, guildid string) string
	Subscribe(nameid, argRoles, guildid string) string
	Unsubscribe(nameid, argRoles, guildid string) string
	AddEnojiRsQueue(chatid, mesid string)
	CheckAdmin(nameid string, chatid string) bool
	BotName() string
	EmbedDS(name1, name2, name3, name4, lvlkz string, numkz int) discordgo.MessageEmbed
	EditMessage(chatID, messageID, content string)
	SendEmbedTime(chatid, text string) string
	Help(Channel string)
	Autohelp()
}

func (d *Discord) EmbedDS(name1, name2, name3, name4, lvlkz string, numkz int) discordgo.MessageEmbed {
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
func (d *Discord) SendEmbedText(chatid, title, text string) *discordgo.Message {
	Emb := &discordgo.MessageEmbed{
		Author:      &discordgo.MessageEmbedAuthor{},
		Color:       16711680,
		Description: text,
		Title:       title,
	}
	m, err := d.d.ChannelMessageSendEmbed(chatid, Emb)
	if err != nil {
		d.log.Println("Ошибка отправки сообщения со вставкой ", err)
	}
	return m
}
func (d *Discord) CheckAdmin(nameid string, chatid string) bool {
	perms, err := d.d.UserChannelPermissions(nameid, chatid)
	if err != nil {
		d.log.Println("ошибка проверки админ ли ", err)
	}
	if perms&discordgo.PermissionAdministrator != 0 {
		return true
	} else {
		return false
	}
}
func (d *Discord) AddEnojiRsQueue(chatid, mesid string) {
	err := d.d.MessageReactionAdd(chatid, mesid, emOK)
	if err != nil {
		d.log.Println("Ошибка добавления реакции на сообщения "+emOK, err)
	}
	err = d.d.MessageReactionAdd(chatid, mesid, emCancel)
	if err != nil {
		d.log.Println("Ошибка добавления реакции на сообщения "+emCancel, err)
	}
	err = d.d.MessageReactionAdd(chatid, mesid, emRsStart)
	if err != nil {
		d.log.Println("Ошибка добавления реакции на сообщения "+emRsStart, err)
	}
	err = d.d.MessageReactionAdd(chatid, mesid, emPl30)
	if err != nil {
		d.log.Println("Ошибка добавления реакции на сообщения "+emPl30, err)
	}

}
func (d *Discord) DeleteMessage(chatid, mesid string) {
	_ = d.d.ChannelMessageDelete(chatid, mesid)
	//надо разобраться как игнорировать ошибку сообщение не найдено
	//if err != nil {d.log.Println("Ошибка удаления дискорд сообщения ", chatid, mesid, err)}
}
func (d *Discord) SendChannelDelSecond(chatid, text string, second int) {
	if text != "" {
		message, err := d.d.ChannelMessageSend(chatid, text)
		if err != nil {
			d.log.Println("ошибка отправки сообщения SendChannelDelSecond", err)
		}
		if second <= 60 {
			go func() {
				time.Sleep(time.Duration(second) * time.Second)
				_ = d.d.ChannelMessageDelete(chatid, message.ID)
				//if err != nil { d.log.Println("Ошибка удаления через секунды ", err) }
			}()
		} else {
			d.dbase.TimerInsert(message.ID, chatid, 0, 0, second)
		}
	}
}
func (d *Discord) RoleToIdPing(rolePing, guildid string) string {
	//создаю переменную
	rolPing := "кз" + rolePing // добавляю буквы
	if guildid == "" {
		d.log.Panic("почему то нет гуилд ид")
		panic("почему то нет гуилд ид")
	}
	g, err := d.d.Guild(guildid)
	if err != nil {
		d.log.Println("ошибка получении гильдии при получении роли", err)
	}
	exist, role := d.roleExists(g, rolPing)
	if !exist {
		//создаем роль и возврашаем пинг
		role = d.CreateRole(rolPing, guildid)
		return role.Mention()
	} else {
		return role.Mention()
	}
}
func (d *Discord) CreateRole(rolPing, guildid string) *discordgo.Role {
	newRole, err := d.d.GuildRoleCreate(guildid)
	if err != nil {
		d.log.Println("ошибка создании новой роли ", err)
	}
	role, err := d.d.GuildRoleEdit(guildid, newRole.ID, rolPing, newRole.Color, newRole.Hoist, 37080064, true)
	if err != nil {
		d.log.Println("Ошибка изменения новой роли", err)
		err = d.d.GuildRoleDelete(guildid, newRole.ID)
		if err != nil {
			d.log.Println("ошибка удаления новой роли ", err)
		}
	}
	return role
}
func (d *Discord) DeleteMesageSecond(chatid, mesid string, second int) {
	if second > 60 {
		d.dbase.TimerInsert(mesid, chatid, 0, 0, second)
	} else {
		go func() {
			time.Sleep(time.Duration(second) * time.Second)
			d.d.ChannelMessageDelete(chatid, mesid)
		}()
	}
}
func (d *Discord) EditComplex(dsmesid, dschatid string, Embeds discordgo.MessageEmbed) {
	_, _ = d.d.ChannelMessageEditComplex(&discordgo.MessageEdit{
		Content: &mesContentNil,
		Embed:   &Embeds,
		ID:      dsmesid,
		Channel: dschatid,
	})
	//if err != nil { d.log.Println("Ошибка редактирования комплексного сообщения ", err) }
}
func (d *Discord) BotName() string { //получаем имя бота
	u, err := d.d.User("@me")
	if err != nil {
		d.log.Println("Ошибка получения имени бота", err)
	}
	return u.Username
}
func (d *Discord) SendComplexContent(chatid, text string) string { //отправка текста комплексного сообщения
	mesCompl, err := d.d.ChannelMessageSendComplex(chatid, &discordgo.MessageSend{
		Content: text})
	if err != nil {
		d.log.Println("Ошибка отправки комплексного сообщения ", err)
	}
	return mesCompl.ID
}
func (d *Discord) SendComplex(chatid string, embeds discordgo.MessageEmbed) string { //отправка текста комплексного сообщения
	mesCompl, err := d.d.ChannelMessageSendComplex(chatid, &discordgo.MessageSend{
		Content: mesContentNil,
		Embed:   &embeds,
	})
	if err != nil {
		d.log.Println("Ошибка отправки комплексного сообщения ", err)
	}
	return mesCompl.ID
}
func (d *Discord) Send(chatid, text string) string { //отправка текста
	message, err := d.d.ChannelMessageSend(chatid, text)
	if err != nil {
		d.log.Println("ошибка отправки текста ", err)
	}
	return message.ID
}
func (d *Discord) Subscribe(nameid, argRoles, guildid string) string {
	g, err := d.d.State.Guild(guildid)
	if err != nil {
		d.log.Println("Ошибка запроса стате.гуилд,читаю гуилд", err)
		g, err = d.d.Guild(guildid)
		if err != nil {
			d.log.Println("Ошибка чтения гуилд ... паниковать ", err)
		}
	}

	exist, role := d.roleExists(g, argRoles)

	if !exist { //если нет роли
		role = d.CreateRole(argRoles, guildid)
	}

	member, err := d.d.GuildMember(guildid, nameid)
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

	err = d.d.GuildMemberRoleAdd(guildid, nameid, role.ID)
	if err != nil {
		d.log.Println("Ошибка выдачи роли ", err)
		subscribe = 2
	}
	var text string
	if subscribe == 0 {
		text = fmt.Sprintf("%s Теперь вы подписаны на %s", member.Mention(), role.Name)
	} else if subscribe == 1 {
		text = fmt.Sprintf("%s Вы уже подписаны на %s", member.Mention(), role.Name)
	} else if subscribe == 2 {
		text = "ошибка: недостаточно прав для выдачи роли " + role.Name
	}
	return text
}
func (d *Discord) Unsubscribe(nameid, argRoles, guildid string) string {
	var unsubscribe int = 0
	g, err := d.d.State.Guild(guildid)
	if err != nil {
		d.log.Println("Ошибка запроса стате.гуилд,читаю гуилд", err)
		g, err = d.d.Guild(guildid)
		if err != nil {
			d.log.Println("Ошибка чтения гуилд ... паниковать ", err)
		}
	}

	exist, role := d.roleExists(g, argRoles)
	if !exist { //если нет роли
		unsubscribe = 1
	}

	member, err := d.d.GuildMember(guildid, nameid)
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
		err = d.d.GuildMemberRoleRemove(guildid, nameid, role.ID)
		if err != nil {
			d.log.Println("Ошибка снятия роли ", err)
			unsubscribe = 3
		}
	}
	text := ""
	if unsubscribe == 0 {
		text = fmt.Sprintf("%s Вы не подписаны на роль %s", member.Mention(), role.Name)
	} else if unsubscribe == 1 {
		text = fmt.Sprintf("%s Роли %s нет на сервере  ", member.Mention(), argRoles)
	} else if unsubscribe == 2 {
		text = fmt.Sprintf("%s Вы отписались от роли %s", member.Mention(), argRoles)
	} else if unsubscribe == 3 {
		text = "ошибка: недостаточно прав для снятия роли  " + role.Name
	}
	return text
}
func (d *Discord) EditMessage(chatID, messageID, content string) {
	_, err := d.d.ChannelMessageEdit(chatID, messageID, content)
	if err != nil {
		d.log.Println("Ошибка изменения текса сообщения ", err)
	}
}
func (d *Discord) SendEmbedTime(chatid, text string) string { //отправка текста с двумя реакциями
	message, err := d.d.ChannelMessageSend(chatid, text)
	if err != nil {
		d.log.Println("ошибка отправки текста ", err)
	}
	err = d.d.MessageReactionAdd(chatid, message.ID, emPlus)
	if err != nil {
		d.log.Println("Ошибка добавления эмоджи ", emPlus, err)
	}
	err = d.d.MessageReactionAdd(chatid, message.ID, emMinus)
	if err != nil {
		d.log.Println("Ошибка добавления эмоджи ", emMinus, err)
	}
	return message.ID
}
