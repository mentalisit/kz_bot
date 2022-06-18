package discordClient

import (
	"fmt"
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
	corpsConfig "kz_bot/internal/clients/corpConfig"
	"kz_bot/internal/dbase/dbaseMysql"
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
	corpsConfig.CorpConfig
	dbase *dbaseMysql.Db
}

func (d *Ds) EmbedDS(name1, name2, name3, name4, lvlkz string, numkz int) discordgo.MessageEmbed {
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
func (d *Ds) SendEmbedText(chatid, title, text string) *discordgo.Message {
	Emb := &discordgo.MessageEmbed{
		Author:      &discordgo.MessageEmbedAuthor{},
		Color:       16711680,
		Description: text,
		Title:       title,
	}
	m, _ := d.d.ChannelMessageSendEmbed(chatid, Emb)
	return m
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
	err := d.d.ChannelMessageDelete(chatid, mesid)
	if err != nil {
		fmt.Println("Ошибка удаления дискорд сообщения ", err)
	}
}
func (d *Ds) SendChannelDelSecond(chatid, text string, second int) {
	message, err := d.d.ChannelMessageSend(chatid, text)
	if err != nil {
		fmt.Println("ошибка отправки сообщения SendChannelDelSecond", err)
	}
	if second <= 60 {
		go func() {
			time.Sleep(time.Duration(second) * time.Second)
			err := d.d.ChannelMessageDelete(chatid, message.ID)
			if err != nil {
				fmt.Println("Ошибка удаления через секунды ", err)
			}
		}()
	} else {
		d.dbase.TimerInsert(message.ID, chatid, 0, 0, second)
	}

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
		role = d.CreateRole(rolPing, guildid)
		return role.Mention()
	} else {
		return role.Mention()
	}
}
func (d *Ds) CreateRole(rolPing, guildid string) *discordgo.Role {
	newRole, err := d.d.GuildRoleCreate(guildid)
	if err != nil {
		fmt.Println("ошибка создании новой роли ", err)
	}
	role, err := d.d.GuildRoleEdit(guildid, newRole.ID, rolPing, newRole.Color, newRole.Hoist, 37080064, true)
	if err != nil {
		fmt.Println("Ошибка изменения новой роли", err)
		err = d.d.GuildRoleDelete(guildid, newRole.ID)
		if err != nil {
			fmt.Println("ошибка удаления новой роли ", err)
		}
	}
	return role
}
func (d *Ds) DeleteMesageSecond(chatid, mesid string, second int) {
	if second > 60 {
		d.dbase.TimerInsert(mesid, chatid, 0, 0, second)
	} else {
		go func() {
			time.Sleep(time.Duration(second) * time.Second)
			d.d.ChannelMessageDelete(chatid, mesid)
		}()
	}

}
func (d *Ds) EditComplex(dsmesid, dschatid string, Embeds discordgo.MessageEmbed) {
	a := &discordgo.MessageEdit{
		Content: &mesContentNil,
		Embed:   &Embeds,
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
func (d *Ds) SendComplex(chatid string, embeds discordgo.MessageEmbed) string { //отправка текста комплексного сообщения
	mesCompl, err := d.d.ChannelMessageSendComplex(chatid, &discordgo.MessageSend{
		Content: mesContentNil,
		Embed:   &embeds,
	})
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
func (d *Ds) Subscribe(nameid, argRoles, guildid string) string {
	g, err := d.d.State.Guild(guildid)
	if err != nil {
		fmt.Println("Ошибка запроса стате.гуилд,читаю гуилд", err)
		g, err = d.d.Guild(guildid)
		if err != nil {
			log.Println("Ошибка чтения гуилд ... паниковать ", err)
		}
	}

	exist, role := d.roleExists(g, argRoles)

	if !exist { //если нет роли
		role = d.CreateRole(argRoles, guildid)
	}

	member, err := d.d.GuildMember(guildid, nameid)
	if err != nil {
		fmt.Println("Ошибка чтения участников гуилд", err)
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
		fmt.Println("Ошибка выдачи роли ", err)
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
func (d *Ds) Unsubscribe(nameid, argRoles, guildid string) string {
	var unsubscribe int = 0
	g, err := d.d.State.Guild(guildid)
	if err != nil {
		fmt.Println("Ошибка запроса стате.гуилд,читаю гуилд", err)
		g, err = d.d.Guild(guildid)
		if err != nil {
			log.Println("Ошибка чтения гуилд ... паниковать ", err)
		}
	}

	exist, role := d.roleExists(g, argRoles)
	if !exist { //если нет роли
		unsubscribe = 1
	}

	member, err := d.d.GuildMember(guildid, nameid)
	if err != nil {
		fmt.Println("Ошибка чтения участников гуилд", err)
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
			fmt.Println("Ошибка снятия роли ", err)
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
func (d *Ds) EditMessage(chatID, messageID, content string) {
	_, err := d.d.ChannelMessageEdit(chatID, messageID, content)
	if err != nil {
		fmt.Println("Ошибка изменения текса сообщения ", err)
	}
}
func (d *Ds) SendEmbedTime(chatid, text string) string { //отправка текста с двумя реакциями
	message, err := d.d.ChannelMessageSend(chatid, text)
	if err != nil {
		fmt.Println("ошибка отправки текста ", err)
	}
	err = d.d.MessageReactionAdd(chatid, message.ID, emPlus)
	if err != nil {
		log.Println("Ошибка добавления эмоджи ", emPlus, err)
	}
	err = d.d.MessageReactionAdd(chatid, message.ID, emMinus)
	if err != nil {
		log.Println("Ошибка добавления эмоджи ", emMinus, err)
	}
	return message.ID
}
