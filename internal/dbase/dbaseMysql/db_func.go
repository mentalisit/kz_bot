package dbaseMysql

import (
	"database/sql"
	"fmt"
	corpsConfig "kz_bot/internal/clients/corpConfig"
	"kz_bot/internal/models"
	"log"
	"time"
)

type Db struct {
	db sql.DB
}
type DbInterface interface {
	AddTgCorpConfig(chatName string, chatid int64)
	DeleteTgchannel(chatid int64)
	ReadBotCorpConfig()
}

func (d *Db) ReadBotCorpConfig() {
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
func (d *Db) DeleteTgChannel(chatid int64) {
	_, err := d.db.Exec("delete from config where tgchannel = ? ", chatid)
	if err != nil {
		fmt.Println("Ошибка удаления с бд корп телеги", err)
	}
}
func (d *Db) DeleteDsChannel(chatid string) {
	_, err := d.db.Exec("delete from config where dschannel = ? ", chatid)
	if err != nil {
		fmt.Println("Ошибка удаления с бд корп дискорд", err)
	}
}
func (d *Db) AddTgCorpConfig(chatName string, chatid int64) {
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
func (d *Db) AddDsCorpConfig(chatName, chatid, guildid string) {
	insertConfig := `INSERT INTO config (corpname,dschannel,tgchannel,wachannel,mesiddshelp,mesidtghelp,delmescomplite) VALUES (?,?,?,?,?,?,?)`
	statement, err := d.db.Prepare(insertConfig)
	if err != nil {
		fmt.Println("Ошибка подготовки внесения в бд конфигурации ", err)
	}
	_, err = statement.Exec(chatName, chatid, 0, "", "", 0, 0, guildid)
	if err != nil {
		fmt.Println("Ошибка внесения конфигурации ", err)
	}
	c := corpsConfig.CorpConfig{}
	c.AddCorp(chatName, chatid, 0, "", 1, "", 0, guildid)
}
func (d *Db) СountName(name, lvlkz, corpName string) int {
	var countNames int
	row := d.db.QueryRow("SELECT  COUNT(*) as count FROM sborkz WHERE name = ? AND lvlkz = ? AND corpname = ? AND active = 0",
		name, lvlkz, corpName)
	err := row.Scan(&countNames)
	if err != nil {
		fmt.Println("Ошибка проверки в очереди ли игрок  ", err)
	}
	return countNames
}
func (d *Db) CountQueue(lvlkz, CorpName string) int { //проверка сколько игровок в очереди
	var count int
	row := d.db.QueryRow("SELECT  COUNT(*) as count FROM sborkz WHERE lvlkz = ? AND corpname = ? AND active = 0",
		lvlkz, CorpName)
	err := row.Scan(&count)
	if err != nil {
		fmt.Println("Ошибка проверки количества игроков в очереди", err)
	}
	return count
}
func (d *Db) CountNumberNameActive1(lvlkz, CorpName, name string) int { // выковыриваем из базы значение количества походов на кз
	var countNumberNameActive1 int
	row := d.db.QueryRow("SELECT  COUNT(*) as count FROM sborkz WHERE lvlkz = ? AND corpname = ? AND name = ? AND active = 1",
		lvlkz, CorpName, name)
	err := row.Scan(&countNumberNameActive1)
	if err != nil {
		fmt.Println("Ошибка чтения количества игр", err)
	}
	return countNumberNameActive1
}
func (d *Db) NumberQueueLvl(lvlkz, CorpName string) int {
	var number int
	row := d.db.QueryRow("SELECT  number FROM numkz WHERE lvlkz = ? AND corpname = ?",
		lvlkz, CorpName)
	err := row.Scan(&number)
	if err != nil {
		fmt.Println("Ошибка чтения нумкз", err)
	}
	if number == 0 {
		insertSmt := "INSERT INTO numkz(lvlkz, number,corpname) VALUES (?,?,?)"
		statement, err := d.db.Prepare(insertSmt)
		if err != nil {
			fmt.Println("Ошибка подготовки внесения нумкз", err)
		}
		_, err = statement.Exec(lvlkz, number, CorpName)
		if err != nil {
			fmt.Println("Ошибка внесения нумкз", err)
		}
	} else {
		return number
	}
	return number
}
func (d *Db) ReadAll(lvlkz, CorpName string) (users models.Users) {
	u := models.Users{
		User1: models.Sborkz{},
		User2: models.Sborkz{},
		User3: models.Sborkz{},
		User4: models.Sborkz{},
	}
	user := 1
	results, err := d.db.Query("SELECT * FROM sborkz WHERE lvlkz = ? AND corpname = ? AND active = 0",
		lvlkz, CorpName)
	if err != nil {
		fmt.Println("Ошибка чтения активной очереди readall", err)
	}
	for results.Next() {
		var t models.Sborkz
		err = results.Scan(&t.Id, &t.Corpname, &t.Name, &t.Mention, &t.Tip, &t.Dsmesid, &t.Tgmesid, &t.Wamesid, &t.Time,
			&t.Date, &t.Lvlkz, &t.Numkzn, &t.Numberkz, &t.Numberevent, &t.Eventpoints, &t.Active, &t.Timedown)
		if user == 1 {
			u.User1 = t
		} else if user == 2 {
			u.User2 = t
		} else if user == 3 {
			u.User3 = t
		} else if user == 4 {
			u.User4 = t
		}
		user = user + 1
	}
	return u
}
func (d *Db) SubscPing(nameMention, lvlkz, CorpName string, tipPing int, TgChannel int64) string {
	var name1, names, men string
	var u models.Users
	if tipPing == 3 {
		u = d.ReadAll(lvlkz, CorpName)
	}

	if rows, err := d.db.Query("SELECT nameid FROM subscribe WHERE lvlkz = ? AND chatid = ? AND tip = ?",
		lvlkz, TgChannel, tipPing); err == nil {
		for rows.Next() {
			rows.Scan(&name1)
			if nameMention == name1 || u.User1.Mention == name1 || u.User2.Mention == name1 || u.User3.Mention == name1 {
				continue
			}
			names = name1 + " "
			men = names + men
		}
		rows.Close()
	}
	return men
}
func (d *Db) InsertQueue(dsmesid, wamesid, CorpName, name, nameMention, tip, lvlkz, timekz string, tgmesid, numkzN int) {
	numevent := 0 //qweryNumevent1(in)
	tm := time.Now()
	mdate := (tm.Format("2006-01-02"))
	mtime := (tm.Format("15:04"))

	insertSborkztg1 := `INSERT INTO sborkz(corpname,name,mention,tip,dsmesid,tgmesid,wamesid,time,date,lvlkz,
                   numkzn,numberkz,numberevent,eventpoints,active,timedown) 
				VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
	_, err := d.db.Exec(insertSborkztg1, CorpName, name, nameMention, tip, dsmesid, tgmesid,
		wamesid, mtime, mdate, lvlkz, numkzN, 0, numevent, 0, 0, timekz)
	if err != nil {
		fmt.Println("Ошибка записи старта очереди", err)
	}
}
func (d *Db) EmReadUsers(name, tip string) models.EmodjiUser {
	results, err := d.db.Query("SELECT * FROM users WHERE name = ? AND tip = ?", name, tip)
	if err != nil {
		fmt.Println("Ощибка чтения эмоджи с БД", err)
	}
	var t models.EmodjiUser
	for results.Next() {
		err = results.Scan(&t.Id, &t.Tip, &t.Name, &t.Em1, &t.Em2, &t.Em3, &t.Em4)
	}
	return t
}
func (d *Db) MesidTgUpdate(mesidtg int, lvlkz string, corpname string) {
	_, err := d.db.Exec(
		`update sborkz set tgmesid = ? where lvlkz = ? AND corpname = ? `,
		mesidtg, lvlkz, corpname)
	if err != nil {
		fmt.Println("Ошибка измениния месайди телеги", err)
	}
}
func (d *Db) MesidDsUpdate(mesidds, lvlkz, corpname string) {
	_, err := d.db.Exec(
		`update sborkz set dsmesid = ? where lvlkz = ? AND corpname = ? `,
		mesidds, lvlkz, corpname)
	if err != nil {
		log.Println("Ошибка измениния месайди дискорда ", err)
	}
}
func (d *Db) UpdateCompliteRS(lvlkz string, dsmesid string, tgmesid int, wamesid string, numberkz int, numberevent int, corpname string) {
	_, err := d.db.Exec(
		`update sborkz set active = 1,dsmesid = ?,tgmesid = ?,wamesid = ?,numberkz = ?,numberevent = ? 
				where lvlkz = ? AND corpname = ? AND active = 0`,
		dsmesid, tgmesid, wamesid, numberkz, numberevent, lvlkz, corpname)
	if err != nil {
		fmt.Println("Ошибка сохранения закрытия очереди", err)
	}
	if numberevent > 0 {
		_, err := d.db.Exec(
			`update rsevent set number = number+1  where corpname = ? AND activeevent = 1`, corpname)
		if err != nil {
			log.Println("Ошибка обновления номера катки ивента ", err)
		}
	}
}
func (d *Db) CountNameQueue(name string) (countNames int) { //проверяем есть ли игрок в других очередях
	row := d.db.QueryRow("SELECT  COUNT(*) as count FROM sborkz WHERE name = ? AND active = 0", name)
	err := row.Scan(&countNames)
	if err != nil {
		fmt.Println("Ошибка проверки игрока в других очередях ", err)
	}
	return countNames
}
func (d *Db) ElseTrue(name string) models.Sborkz {
	results, err := d.db.Query("SELECT * FROM sborkz WHERE name = ? AND active = 0", name)
	if err != nil {
		fmt.Println("Ошибка извлечения игрока с других очередей ", err)
	}
	var t models.Sborkz
	for results.Next() {
		err = results.Scan(&t.Id, &t.Corpname, &t.Name, &t.Mention, &t.Tip, &t.Dsmesid, &t.Tgmesid, &t.Wamesid, &t.Time, &t.Date, &t.Lvlkz, &t.Numkzn, &t.Numberkz, &t.Numberevent, &t.Eventpoints, &t.Active, &t.Timedown)

	}
	return t
}
func (d *Db) DeleteQueue(name, lvlkz, CorpName string) {
	_, err := d.db.Exec("delete from sborkz where name = ? AND lvlkz = ? AND corpname = ? AND active = 0",
		name, lvlkz, CorpName)
	if err != nil {
		fmt.Println("Ошибка удаления из очереди ", err)
	}
}
func (d *Db) UpdateMitutsQueue(name, CorpName string) models.Sborkz {
	results, err := d.db.Query("SELECT * FROM sborkz WHERE name = ? AND corpname = ? AND active = 0",
		name, CorpName)
	if err != nil {
		fmt.Println("Ошибка проверки игрока в очереди для функции (-+) ", err)
	}
	var t models.Sborkz
	for results.Next() {

		err = results.Scan(&t.Id, &t.Corpname, &t.Name, &t.Mention, &t.Tip, &t.Dsmesid, &t.Tgmesid, &t.Wamesid, &t.Time,
			&t.Date, &t.Lvlkz, &t.Numkzn, &t.Numberkz, &t.Numberevent, &t.Eventpoints, &t.Active, &t.Timedown)

		if t.Name == name && t.Timedown <= 3 {
			_, err := d.db.Exec("update sborkz set timedown = timedown + 30 where active = 0 AND name = ? AND corpname = ?",
				t.Name, t.Corpname)
			if err != nil {
				fmt.Println("Ошибка обновления времени игрока в очереди для функции (-+) ", err)
			}
			return t
		}
	}
	return t
}
func (d *Db) CheckSubscribe(name, lvlkz string, TgChannel int64, tipPing int) int {
	var counts int
	row := d.db.QueryRow("SELECT  COUNT(*) as count FROM subscribe WHERE name = ? AND lvlkz = ? AND chatid = ? AND tip = ?",
		name, lvlkz, TgChannel, tipPing)
	err := row.Scan(&counts)
	if err != nil {
		log.Println("Ошибка проврки активной подписки ", err)
	}
	return counts
}
func (d *Db) Subscribe(name, nameMention, lvlkz string, tipPing int, TgChannel int64) {
	insertSubscribe := `INSERT INTO subscribe (name, nameid, lvlkz, tip, chatid, timestart, timeend) VALUES (?,?,?,?,?,?,?)`
	statement, err := d.db.Prepare(insertSubscribe)
	_, err = statement.Exec(name, nameMention, lvlkz, tipPing, TgChannel, 0, 0)
	if err != nil {
		log.Println("Ошибка внесения в таблицу подписок ", err)
	}
}
func (d *Db) Unsubscribe(name, lvlkz string, TgChannel int64, tipPing int) {
	_, err := d.db.Exec("delete from subscribe where name = ? AND lvlkz = ? AND chatid = ? AND tip = ?",
		name, lvlkz, TgChannel, tipPing)
	if err != nil {
		fmt.Println("Ошибка удаления подписки с БД", err)
	}
}
