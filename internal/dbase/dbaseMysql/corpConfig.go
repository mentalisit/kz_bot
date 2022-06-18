package dbaseMysql

import (
	"fmt"

	corpsConfig "kz_bot/internal/clients/corpConfig"
	"kz_bot/internal/models"
)

func (d *Db) ReadBotCorpConfig() {
	c := corpsConfig.CorpConfig{}
	results, err := d.Db.Query("SELECT * FROM config")
	if err != nil {
		fmt.Println("Ошибка чтения крнфигурации корпораций", err)
	}
	var t models.TableConfig
	for results.Next() {
		err = results.Scan(&t.Id, &t.Corpname, &t.Dschannel, &t.Tgchannel, &t.Wachannel, &t.Mesiddshelp, &t.Mesidtghelp, &t.Delmescomplite, &t.Guildid)
		c.AddCorp(t.Corpname, t.Dschannel, t.Tgchannel, t.Wachannel, t.Delmescomplite, t.Mesiddshelp, t.Mesidtghelp, t.Guildid)
		fmt.Println("Конфиг корпораци", t)
	}
}
func (d *Db) DeleteTgChannel(chatid int64) {
	_, err := d.Db.Exec("delete from config where tgchannel = ? ", chatid)
	if err != nil {
		fmt.Println("Ошибка удаления с бд корп телеги", err)
	}
}
func (d *Db) DeleteDsChannel(chatid string) {
	_, err := d.Db.Exec("delete from config where dschannel = ? ", chatid)
	if err != nil {
		fmt.Println("Ошибка удаления с бд корп дискорд", err)
	}
}
func (d *Db) AddTgCorpConfig(chatName string, chatid int64) {
	fmt.Println(chatName, chatid, &d.Db)
	insertConfig := `INSERT INTO config (corpname,dschannel,tgchannel,wachannel,mesiddshelp,mesidtghelp,delmescomplite,guildid) 
						VALUES (?,?,?,?,?,?,?,?)`
	r, e := d.Db.Exec(insertConfig, chatName, "", chatid, "", "", 0, 0, "")
	fmt.Println(r, e)
	//statement, err := d.Db.Prepare(insertConfig)
	//if err != nil {
	//	fmt.Println("Ошибка подготовки внесения в бд конфигурации ", err)
	//}
	//_, err = statement.Exec(chatName, "", chatid, "", "", 0, 0,"")
	_, err := d.Db.Exec(insertConfig, chatName, "", chatid, "", "", 0, 0, "")
	if err != nil {
		fmt.Println("Ошибка внесения конфигурации ", err)
	}
	c := corpsConfig.CorpConfig{}
	c.AddCorp(chatName, "", chatid, "", 1, "", 0, "")
}
func (d *Db) AddDsCorpConfig(chatName, chatid, guildid string) {
	insertConfig := `INSERT INTO config (corpname,dschannel,tgchannel,wachannel,mesiddshelp,mesidtghelp,delmescomplite,guildid) VALUES (?,?,?,?,?,?,?,?)`
	//statement, err := d.Db.Prepare(insertConfig)
	//if err != nil {
	//	fmt.Println("Ошибка подготовки внесения в бд конфигурации ", err)
	//}
	_, err := d.Db.Exec(insertConfig, chatName, chatid, 0, "", "", 0, 0, guildid)
	//_, err = statement.Exec(chatName, chatid, 0, "", "", 0, 0, guildid)
	if err != nil {
		fmt.Println("Ошибка внесения конфигурации ", err)
	}
	c := corpsConfig.CorpConfig{}
	c.AddCorp(chatName, chatid, 0, "", 1, "", 0, guildid)
}
