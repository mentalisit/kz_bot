package dbasePostgres

import (
	"context"
	"kz_bot/internal/models"
)

type CorpConfig interface {
	ReadBotCorpConfig()                                 //Чтение из бд конфигураций корпораций при запуске бота
	DeleteTgChannel(chatid int64)                       //отключение бота от чата в телеграм
	DeleteDsChannel(chatid string)                      //отключение бота от чата в дискорд
	DeleteWaChannel(chatid string)                      //отключение бота от чата в Wa
	AddTgCorpConfig(chatName string, chatid int64)      //добавление чата телеграм в конфиг корпораций
	AddDsCorpConfig(chatName, chatid, guildid string)   //добавление чата дискорд в конфиг корпораций
	AddWaCorpConfig(chatName, chatid string)            //добавление чата WA в конфиг корпораций
	AutoHelpUpdateMesid(newMesidHelp, dschannel string) //обновления месидДсХелп
}

func (d *Db) ReadBotCorpConfig() {
	results, err := d.Db.Query(context.Background(), `SELECT * FROM kzbot.config`)
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
	_, err := d.Db.Exec(context.Background(), "delete from kzbot.config where tgchannel = $1 ", chatid)
	if err != nil {
		d.log.Println("Ошибка удаления с бд корп телеги", err)
	}
}
func (d *Db) DeleteDsChannel(chatid string) {
	_, err := d.Db.Exec(context.Background(), "delete from kzbot.config where dschannel = $1 ", chatid)
	if err != nil {
		d.log.Println("Ошибка удаления с бд корп дискорд", err)
	}
}
func (d *Db) DeleteWaChannel(chatid string) {
	_, err := d.Db.Exec(context.Background(), "delete from kzbot.config where wachannel = $1 ", chatid)
	if err != nil {
		d.log.Println("Ошибка удаления с бд корп wa", err)
	}
}
func (d *Db) AddTgCorpConfig(chatName string, chatid int64) {
	d.log.Println(chatName, "Добавлена в конфиг корпораций ")
	insertConfig := `INSERT INTO kzbot.config (corpname,dschannel,tgchannel,wachannel,mesiddshelp,mesidtghelp,delmescomplite,guildid) 
						VALUES ($1,$2,$3,$4,$5,$6,$7,$8)`
	_, err := d.Db.Exec(context.Background(), insertConfig, chatName, "", chatid, "", "", 0, 0, "")
	if err != nil {
		d.log.Println("Ошибка внесения конфигурации ", err)
	}
	//c := corpsConfig.CorpConfig{}
	d.CorpConfig.AddCorp(chatName, "", chatid, "", 1, "", 0, "")
}
func (d *Db) AddDsCorpConfig(chatName, chatid, guildid string) {
	insertConfig := `INSERT INTO kzbot.config (corpname,dschannel,tgchannel,wachannel,mesiddshelp,mesidtghelp,delmescomplite,guildid)
					VALUES ($1,$2,$3,$4,$5,$6,$7,$8)`
	_, err := d.Db.Exec(context.Background(), insertConfig, chatName, chatid, 0, "", "", 0, 0, guildid)
	if err != nil {
		d.log.Println("Ошибка внесения конфигурации ", err)
	}
	//c := corpsConfig.CorpConfig{}
	d.CorpConfig.AddCorp(chatName, chatid, 0, "", 1, "", 0, guildid)
}
func (d *Db) AddWaCorpConfig(chatName, chatid string) {
	insertConfig := `INSERT INTO kzbot.config (corpname,dschannel,tgchannel,wachannel,mesiddshelp,mesidtghelp,delmescomplite,guildid)
					VALUES ($1,$2,$3,$4,$5,$6,$7,$8)`
	_, err := d.Db.Exec(context.Background(), insertConfig, chatName, "", 0, chatid, "", 0, 0, "")
	if err != nil {
		d.log.Println("Ошибка внесения конфигурации ", err)
	}
	//c := corpsConfig.CorpConfig{}
	d.CorpConfig.AddCorp(chatName, "", 0, chatid, 1, "", 0, "")
}
func (d *Db) AutoHelpUpdateMesid(newMesidHelp, dschannel string) {
	updateString := `update kzbot.config set "mesiddshelp"=$1 where "dschannel"=$2`
	_, err := d.Db.Exec(context.Background(), updateString, newMesidHelp, dschannel)
	if err != nil {
		d.log.Println("ОШибка обновления месИд для автосправки ", err)
	}
}
