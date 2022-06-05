package dbaseMysql

import (
	"database/sql"
	"fmt"
	corpsConfig "kz_bot/internal/clients/corpConfig"
	"kz_bot/internal/models"
)

type Db struct {
	db sql.DB
}
type DbInterface interface {
	AddTgCorpConfig(chatName string, chatid int64)
	DeleteTgchannel(chatid int64)
	ReadBotCorpConfig()
}

func (d Db) ReadBotCorpConfig() {
	fmt.Println(d.db.Ping())
	c := corpsConfig.CorpConfig{}
	results, err := d.db.Query("SELECT * FROM config")
	if err != nil {
		fmt.Println("Ошибка чтения крнфигурации корпораций", err)
	}
	var t models.TableConfig
	for results.Next() {
		err = results.Scan(&t.Id, &t.Corpname, &t.Dschannel, &t.Tgchannel, &t.Wachannel, &t.Mesiddshelp, &t.Mesidtghelp, &t.Delmescomplite, &t.Guildid)
		c.AddCorp(t.Corpname, t.Dschannel, t.Tgchannel, t.Wachannel, t.Delmescomplite, t.Mesiddshelp, t.Mesidtghelp, t.Guildid)
	}
}
func (d Db) DeleteTgchannel(chatid int64) {
	_, err := d.db.Exec("delete from config where tgchannel = ? ", chatid)
	if err != nil {
		fmt.Println("Ошибка удаления с бд корп телеги", err)
	}
}
func (d Db) AddTgCorpConfig(chatName string, chatid int64) {
	insertConfig := `INSERT INTO config (corpname,dschannel,tgchannel,wachannel,mesiddshelp,mesidtghelp,delmescomplite) VALUES (?,?,?,?,?,?,?)`
	statement, err := d.db.Prepare(insertConfig)
	if err != nil {
		fmt.Println("Ошибка подготовки внесения в бд конфигурации ", err)
	}
	_, err = statement.Exec(chatName, "", chatid, "", "", 0, 0)
	if err != nil {
		fmt.Println("Ошибка внесения конфигурации ", err)
	}
	c := corpsConfig.CorpConfig{}
	c.AddCorp(chatName, "", chatid, "", 1, "", 0, "")
}
