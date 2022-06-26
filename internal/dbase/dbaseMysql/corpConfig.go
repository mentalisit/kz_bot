package dbaseMysql

import (
	"kz_bot/internal/models"
)

func (d *Db) ReadBotCorpConfig() {
	//c := corpsConfig.CorpConfig{}
	results, err := d.Db.Query("SELECT * FROM config")
	if err != nil {
		d.log.Println("Ошибка чтения крнфигурации корпораций", err)
	}
	var t models.TableConfig
	var corp []string
	for results.Next() {
		err = results.Scan(&t.Id, &t.Corpname, &t.Dschannel, &t.Tgchannel, &t.Wachannel, &t.Mesiddshelp, &t.Mesidtghelp, &t.Delmescomplite, &t.Guildid)
		d.CorpConfig.AddCorp(t.Corpname, t.Dschannel, t.Tgchannel, t.Wachannel, t.Delmescomplite, t.Mesiddshelp, t.Mesidtghelp, t.Guildid)
		corp = append(corp, t.Corpname)
	}
	d.log.Println("Конфиг корпораций", corp)
}
func (d *Db) DeleteTgChannel(chatid int64) {
	_, err := d.Db.Exec("delete from config where tgchannel = ? ", chatid)
	if err != nil {
		d.log.Println("Ошибка удаления с бд корп телеги", err)
	}
}
func (d *Db) DeleteDsChannel(chatid string) {
	_, err := d.Db.Exec("delete from config where dschannel = ? ", chatid)
	if err != nil {
		d.log.Println("Ошибка удаления с бд корп дискорд", err)
	}
}
func (d *Db) AddTgCorpConfig(chatName string, chatid int64) {
	d.log.Println(chatName, "Добавлена в конфиг корпораций ")
	insertConfig := `INSERT INTO config (corpname,dschannel,tgchannel,wachannel,mesiddshelp,mesidtghelp,delmescomplite,guildid) 
						VALUES (?,?,?,?,?,?,?,?)`
	statement, err := d.Db.Prepare(insertConfig)
	if err != nil {
		d.log.Println("Ошибка подготовки внесения в бд конфигурации ", err)
	}
	_, err = statement.Exec(chatName, "", chatid, "", "", 0, 0, "")
	if err != nil {
		d.log.Println("Ошибка внесения конфигурации ", err)
	}
	//c := corpsConfig.CorpConfig{}
	d.CorpConfig.AddCorp(chatName, "", chatid, "", 1, "", 0, "")
}
func (d *Db) AddDsCorpConfig(chatName, chatid, guildid string) {
	insertConfig := `INSERT INTO config (corpname,dschannel,tgchannel,wachannel,mesiddshelp,mesidtghelp,delmescomplite,guildid) VALUES (?,?,?,?,?,?,?,?)`
	statement, err := d.Db.Prepare(insertConfig)
	if err != nil {
		d.log.Println("Ошибка подготовки внесения в бд конфигурации ", err)
	}
	_, err = statement.Exec(chatName, chatid, 0, "", "", 0, 0, guildid)
	if err != nil {
		d.log.Println("Ошибка внесения конфигурации ", err)
	}
	//c := corpsConfig.CorpConfig{}
	d.CorpConfig.AddCorp(chatName, chatid, 0, "", 1, "", 0, guildid)
}
