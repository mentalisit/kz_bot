package dbasePostgres

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jackc/pgx/v4"
	"strconv"
	"time"

	"kz_bot/internal/models"
)

type DbInterface interface {
	NumberQueueLvl(lvlkz, CorpName string) int                                                                 //Номер катки по уровню
	ReadAll(lvlkz, CorpName string) (users models.Users)                                                       //чтение игроков в очереди
	InsertQueue(dsmesid, wamesid, CorpName, name, nameMention, tip, lvlkz, timekz string, tgmesid, numkzN int) //внесение данных сбора
	ElseTrue(name string) models.Sborkz                                                                        //для выхода из очереди при другом старте
	DeleteQueue(name, lvlkz, CorpName string)                                                                  //Если игрок покидает очередь
	UpdateMitutsQueue(name, CorpName string) models.Sborkz                                                     //проверка хочет ли игрок продолжить время в очереди
	TimerInsert(dsmesid, dschatid string, tgmesid int, tgchatid int64, timed int)                              //внесение ид сообщения в бд
	TimerDeleteMessage() []models.Timer                                                                        //удаление из таймера
	P30Pl(lvlkz, CorpName, name string) int                                                                    //+30 минут если в очереди
	UpdateTimedown(lvlkz, CorpName, name string)                                                               //при нажатии плюса при остатке меньше 3х минут
	ReadMesIdDS(mesid string) (string, error)
	Queue(corpname string) []string
	AutoHelp() []models.BotConfig
	AutoHelpUpdateMesid(newMesidHelp, dschannel string)
	MinusMin() []models.Sborkz
	OneMinutsTimer() []string
	MessageUpdateMin(corpname string) ([]string, []int, []string)
	MessageupdateDS(dsmesid string, config models.BotConfig) models.InMessage
	MessageupdateTG(tgmesid int, config models.BotConfig) models.InMessage
	ReadStatistic(Name string) string
	Shutdown()
}

func (d *Db) Shutdown() {
	d.Db.Close()
}

func (d *Db) NumberQueueLvl(lvlkz, CorpName string) int {
	var number int
	sel := "SELECT  number FROM kzbot.numkz WHERE lvlkz = $1 AND corpname = $2"
	row := d.Db.QueryRow(context.Background(), sel, lvlkz, CorpName)
	err := row.Scan(&number)
	if err != nil {
		if err == pgx.ErrNoRows {
			number = 0
			insertSmt := "INSERT INTO kzbot.numkz(lvlkz, number,corpname) VALUES ($1,$2,$3)"
			_, err = d.Db.Exec(context.Background(), insertSmt, lvlkz, number, CorpName)
			if err != nil {
				d.log.Println("Ошибка внесения нумкз", err)
			}
			return number + 1
		} else {
			d.log.Println("Ошибка чтения нумкз", err)
		}
	}
	return number + 1
}
func (d *Db) ReadAll(lvlkz, CorpName string) (users models.Users) {
	u := models.Users{
		User1: models.Sborkz{},
		User2: models.Sborkz{},
		User3: models.Sborkz{},
		User4: models.Sborkz{},
	}
	user := 1
	sel := "SELECT * FROM kzbot.sborkz WHERE lvlkz = $1 AND corpname = $2 AND active = 0"
	results, err := d.Db.Query(context.Background(), sel, lvlkz, CorpName)
	if err != nil {
		d.log.Println("Ошибка чтения активной очереди readall", err)
	}
	for results.Next() {
		var t models.Sborkz
		err = results.Scan(&t.Id, &t.Corpname, &t.Name, &t.Mention, &t.Tip, &t.Dsmesid,
			&t.Tgmesid, &t.Wamesid, &t.Time, &t.Date, &t.Lvlkz, &t.Numkzn, &t.Numberkz,
			&t.Numberevent, &t.Eventpoints, &t.Active, &t.Timedown)
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
func (d *Db) InsertQueue(dsmesid, wamesid, CorpName, name, nameMention, tip, lvlkz, timekz string, tgmesid, numkzN int) {
	numevent := d.NumActiveEvent(CorpName)
	tm := time.Now()
	mdate := (tm.Format("2006-01-02"))
	mtime := (tm.Format("15:04"))
	timekzz, errs := strconv.Atoi(timekz)
	if timekzz == 0 {
		d.log.Panic(errs)
	}

	insertSborkztg1 := `INSERT INTO kzbot.sborkz(corpname,name,mention,tip,dsmesid,tgmesid,wamesid,time,date,lvlkz,
                   numkzn,numberkz,numberevent,eventpoints,active,timedown) 
				VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16)`
	_, err := d.Db.Exec(context.Background(), insertSborkztg1, CorpName, name, nameMention, tip, dsmesid, tgmesid,
		wamesid, mtime, mdate, lvlkz, numkzN, 0, numevent, 0, 0, timekzz)
	if err != nil {
		d.log.Println("Ошибка записи старта очереди", err)
	}
}

func (d *Db) ElseTrue(name string) models.Sborkz {
	sel := "SELECT * FROM kzbot.sborkz WHERE name = $1 AND active = 0"
	results, err := d.Db.Query(context.Background(), sel, name)
	if err != nil {
		d.log.Println("Ошибка извлечения игрока с других очередей ", err)
	}
	var t models.Sborkz
	for results.Next() {
		err = results.Scan(&t.Id, &t.Corpname, &t.Name, &t.Mention, &t.Tip, &t.Dsmesid, &t.Tgmesid, &t.Wamesid, &t.Time, &t.Date, &t.Lvlkz, &t.Numkzn, &t.Numberkz, &t.Numberevent, &t.Eventpoints, &t.Active, &t.Timedown)

	}
	return t
}
func (d *Db) DeleteQueue(name, lvlkz, CorpName string) {
	del := "delete from kzbot.sborkz where name = $1 AND lvlkz = $2 AND corpname = $3 AND active = 0"
	_, err := d.Db.Exec(context.Background(), del, name, lvlkz, CorpName)
	if err != nil {
		d.log.Println("Ошибка удаления из очереди ", err)
	}
}
func (d *Db) UpdateMitutsQueue(name, CorpName string) models.Sborkz {
	sel := "SELECT * FROM kzbot.sborkz WHERE name = $1 AND corpname = $2 AND active = 0"
	results, err := d.Db.Query(context.Background(), sel, name, CorpName)
	if err != nil {
		d.log.Println("Ошибка проверки игрока в очереди для функции (-+) ", err)
	}
	var t models.Sborkz
	for results.Next() {

		err = results.Scan(&t.Id, &t.Corpname, &t.Name, &t.Mention, &t.Tip, &t.Dsmesid, &t.Tgmesid, &t.Wamesid, &t.Time,
			&t.Date, &t.Lvlkz, &t.Numkzn, &t.Numberkz, &t.Numberevent, &t.Eventpoints, &t.Active, &t.Timedown)

		if t.Name == name && t.Timedown <= 3 {
			upd := "update kzbot.sborkz set timedown = timedown + 30 where active = 0 AND name = $1 AND corpname = $2"
			_, err = d.Db.Exec(context.Background(), upd, t.Name, t.Corpname)
			if err != nil {
				d.log.Println("Ошибка обновления времени игрока в очереди для функции (-+) ", err)
			}
			return t
		}
	}
	return t
}
func (d *Db) TimerInsert(dsmesid, dschatid string, tgmesid int, tgchatid int64, timed int) {
	insertTimer := `INSERT INTO kzbot.timer(dsmesid,dschatid,tgmesid,tgchatid,timed) VALUES ($1,$2,$3,$4,$5)`
	_, err := d.Db.Exec(context.Background(), insertTimer, dsmesid, dschatid, tgmesid, tgchatid, timed)
	if err != nil {
		d.log.Println("Ошибка внесения в бд для удаления ", err)
	}
}
func (d *Db) TimerDeleteMessage() []models.Timer {
	upd := `update kzbot.timer set timed = timed - 60`
	_, err := d.Db.Exec(context.Background(), upd)
	if err != nil {
		d.log.Println("Ошибка удаления 60секунд", err)
	}

	sel := "SELECT * FROM kzbot.timer WHERE timed < 60"
	results, err := d.Db.Query(context.Background(), sel)
	if err != nil {
		if err != sql.ErrNoRows {
			d.log.Println("Ошибка чтения ид где меньше 60 секунд", err)
		}

	}
	var timedown []models.Timer
	for results.Next() {
		var t models.Timer
		err = results.Scan(&t.Id, &t.Dsmesid, &t.Dschatid, &t.Tgmesid, &t.Tgchatid, &t.Timed)
		timedown = append(timedown, t)

		del := "delete from kzbot.timer where  id = $1 "
		_, err = d.Db.Exec(context.Background(), del, t.Id)
		if err != nil {
			d.log.Println("Ошибка удаления по ид с таблицы таймера", err)
		}
	}
	return timedown
}
func (d *Db) ReadMesIdDS(mesid string) (string, error) {
	sel := "SELECT lvlkz FROM kzbot.sborkz WHERE dsmesid = $1 AND active = 0"
	results, err := d.Db.Query(context.Background(), sel, mesid)
	if err != nil {
		d.log.Println("Ошибка получения уровня кз по меседж айди", err)
	}
	a := []string{}
	var dsmesid string
	for results.Next() {
		var t models.Sborkz
		err = results.Scan(&t.Lvlkz)
		a = append(a, t.Lvlkz)
	}
	a = d.removeDuplicateElementString(a)
	if len(a) > 0 {
		dsmesid = a[0]
		return dsmesid, err
	} else {
		return "", err
	}
}
func (d *Db) P30Pl(lvlkz, CorpName, name string) int {
	var timedown int
	sel := "SELECT timedown FROM kzbot.sborkz WHERE lvlkz = $1 AND corpname = $2 AND active = 0 AND name = $3"
	results, err := d.Db.Query(context.Background(), sel, lvlkz, CorpName, name)
	if err != nil {
		d.log.Println("Ошибка получения оставшегося времени ", err)
	}
	for results.Next() {
		err = results.Scan(&timedown)
	}
	return timedown
}
func (d *Db) UpdateTimedown(lvlkz, CorpName, name string) {
	upd := `update kzbot.sborkz set timedown = timedown+30 where lvlkz = $1 AND corpname = $2 AND name = $3`
	_, err := d.Db.Exec(context.Background(), upd, lvlkz, CorpName, name)
	if err != nil {
		d.log.Println("Ошибка обновления времени ", err)
	}
}
func (d *Db) Queue(corpname string) []string {
	sel := "SELECT lvlkz FROM kzbot.sborkz WHERE corpname = $1 AND active = 0"
	results, err := d.Db.Query(context.Background(), sel, corpname)
	if err != nil {
		d.log.Println("Ошибка чтения левелов для очереди", err)
	}
	var lvl []string
	for results.Next() {
		var t models.Sborkz
		err = results.Scan(&t.Lvlkz)

		lvl = append(lvl, t.Lvlkz)

	}

	return lvl
}
func (d *Db) AutoHelp() []models.BotConfig {
	sel := "SELECT dschannel,mesiddshelp FROM kzbot.config"
	results, err := d.Db.Query(context.Background(), sel)
	if err != nil {
		d.log.Println("Ошибка получения автосправки с бд", err)
	}
	h := models.BotConfig{}
	var a []models.BotConfig
	for results.Next() {
		err = results.Scan(&h.DsChannel, &h.Config.MesidDsHelp)
		a = append(a, h)
	}
	return a
}
func (d *Db) MinusMin() []models.Sborkz {
	upd := `update kzbot.sborkz set timedown = timedown - 1 where active = 0`
	_, err := d.Db.Exec(context.Background(), upd)
	if err != nil {
		d.log.Println("Ошибка удаления минуты ", err)
	}

	sel := "SELECT * FROM kzbot.sborkz WHERE active = 0"
	results, err := d.Db.Query(context.Background(), sel)
	if err != nil {
		d.log.Println("Ошибка чтения после удаления минуты", err)
	}
	var tt []models.Sborkz
	for results.Next() {
		var t models.Sborkz
		err = results.Scan(&t.Id, &t.Corpname, &t.Name, &t.Mention, &t.Tip, &t.Dsmesid, &t.Tgmesid, &t.Wamesid, &t.Time, &t.Date, &t.Lvlkz, &t.Numkzn, &t.Numberkz, &t.Numberevent, &t.Eventpoints, &t.Active, &t.Timedown)
		tt = append(tt, t)

	}
	return tt
}
func (d *Db) OneMinutsTimer() []string {
	var count int //количество активных игроков
	sel := "SELECT  COUNT(*) as count FROM kzbot.sborkz WHERE active = 0"
	row := d.Db.QueryRow(context.Background(), sel)
	err := row.Scan(&count)
	if err != nil {
		d.log.Println("Ошибка подсчета активных игроков в очередях", err)
	}
	var CorpActive0 []string
	if count > 0 {
		a := []string{}
		aa := []string{}
		selC := "SELECT corpname FROM kzbot.sborkz WHERE active = 0"
		results, err := d.Db.Query(context.Background(), selC)
		if err != nil {
			d.log.Println("Ошибка чтения корпораций где есть активные очереди ", err)
		}
		var corpname string // ищим корпорации
		for results.Next() {
			err = results.Scan(&corpname)
			a = append(a, corpname)
		}
		a = d.removeDuplicateElementString(a)

		for _, corp := range a {
			skip := false
			for _, u := range aa {
				if corp == u {
					skip = true
					break
				}
			}
			if !skip {
				CorpActive0 = append(CorpActive0, corp)
			}
		}
	}
	return CorpActive0
}
func (d *Db) MessageUpdateMin(corpname string) ([]string, []int, []string) {
	var countCorp int
	ds := []string{}
	tg := []int{}
	wa := []string{}
	sel := "SELECT  COUNT(*) as count FROM kzbot.sborkz WHERE corpname = $1 AND active = 0"
	row := d.Db.QueryRow(context.Background(), sel, corpname)
	err := row.Scan(&countCorp)
	if err != nil {
		d.log.Println("Ошибка получения активных очередей корпорации ", err)
	}
	if countCorp > 0 {
		selS := "SELECT * FROM kzbot.sborkz WHERE corpname = $1 AND active = 0"
		results, err1 := d.Db.Query(context.Background(), selS, corpname)
		if err1 != nil {
			d.log.Println("Ошибка получения активных очередей корпорации2 ", err1)
		}
		for results.Next() {
			var t models.Sborkz
			err = results.Scan(&t.Id, &t.Corpname, &t.Name, &t.Mention, &t.Tip, &t.Dsmesid, &t.Tgmesid, &t.Wamesid, &t.Time, &t.Date, &t.Lvlkz, &t.Numkzn, &t.Numberkz, &t.Numberevent, &t.Eventpoints, &t.Active, &t.Timedown)
			ds = append(ds, t.Dsmesid)
			tg = append(tg, t.Tgmesid)
			wa = append(wa, t.Wamesid)
		}
	}
	ds = d.removeDuplicateElementString(ds)
	tg = d.removeDuplicateElementInt(tg)
	wa = d.removeDuplicateElementString(wa)
	return ds, tg, wa
}
func (d *Db) MessageupdateDS(dsmesid string, config models.BotConfig) models.InMessage {
	var in models.InMessage
	sel := "SELECT * FROM kzbot.sborkz WHERE dsmesid = $1 AND active = 0"
	results, err := d.Db.Query(context.Background(), sel, dsmesid)
	if err != nil {
		d.log.Println(err)
	}
	var t models.Sborkz
	for results.Next() {
		err = results.Scan(&t.Id, &t.Corpname, &t.Name, &t.Mention, &t.Tip, &t.Dsmesid, &t.Tgmesid, &t.Wamesid, &t.Time, &t.Date, &t.Lvlkz, &t.Numkzn, &t.Numberkz, &t.Numberevent, &t.Eventpoints, &t.Active, &t.Timedown)
	}
	in = models.InMessage{
		Tip:         "ds",
		Name:        t.Name,
		NameMention: t.Mention,
		Lvlkz:       t.Lvlkz,
		Ds: struct {
			Mesid   string
			Nameid  string
			Guildid string
		}{
			Mesid:   t.Dsmesid,
			Nameid:  "",
			Guildid: config.Config.Guildid,
		},
		Config: config,
		Option: struct {
			Callback bool
			Edit     bool
			Update   bool
			Queue    bool
		}{
			Callback: true,
			Edit:     true,
			Update:   false,
		},
	}

	return in
}
func (d *Db) MessageupdateTG(tgmesid int, config models.BotConfig) models.InMessage {
	sel := "SELECT * FROM kzbot.sborkz WHERE tgmesid = $1 AND active = 0"
	results, err := d.Db.Query(context.Background(), sel, tgmesid)
	if err != nil {
		d.log.Println(err)
	}
	var t models.Sborkz
	for results.Next() {
		err = results.Scan(&t.Id, &t.Corpname, &t.Name, &t.Mention, &t.Tip, &t.Dsmesid, &t.Tgmesid, &t.Wamesid, &t.Time, &t.Date, &t.Lvlkz, &t.Numkzn, &t.Numberkz, &t.Numberevent, &t.Eventpoints, &t.Active, &t.Timedown)
	}
	in := models.InMessage{
		Tip:         "tg",
		Name:        t.Name,
		NameMention: t.Mention,
		Lvlkz:       t.Lvlkz,
		Tg: struct {
			Mesid  int
			Nameid int64
		}{
			Mesid:  t.Tgmesid,
			Nameid: 0},
		Config: config,
		Option: struct {
			Callback bool
			Edit     bool
			Update   bool
			Queue    bool
		}{
			Callback: true,
			Edit:     true,
			Update:   false,
		},
	}
	return in
}
func (d *Db) ReadStatistic(Name string) string {
	num := 1
	str := "√ уровень кз время дата канал\n"
	tmp := ""
	sel := "SELECT * FROM kzbot.sborkz WHERE name = $1 AND active = 1"
	results, err := d.Db.Query(context.Background(), sel, Name)
	if err != nil {
		d.log.Println("Ошибка чтения statistic", err)
		if err == sql.ErrNoRows {
			return "Информация не найдена "
		}
	}
	for results.Next() {
		var t models.Sborkz
		err = results.Scan(&t.Id, &t.Corpname, &t.Name, &t.Mention, &t.Tip, &t.Dsmesid, &t.Tgmesid, &t.Wamesid, &t.Time,
			&t.Date, &t.Lvlkz, &t.Numkzn, &t.Numberkz, &t.Numberevent, &t.Eventpoints, &t.Active, &t.Timedown)
		tmp = fmt.Sprintf("%d. %s (%s %s) %s\n", num, t.Lvlkz, t.Time, t.Date, t.Corpname)
		num++
		str = str + tmp
		if num == 40 {
			break
		}
	}
	return str
}
