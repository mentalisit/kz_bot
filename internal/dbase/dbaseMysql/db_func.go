package dbaseMysql

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"

	"kz_bot/internal/models"
)

type DbInterface interface {
	СountName(name, lvlkz, corpName string) int //проверка состоит ли игрок уже в очереди
	CountQueue(lvlkz, CorpName string) int      //проверка сколько игроков в очереди
	CountNameQueueCorp(name, corp string) (countNames int)
	CountNumberNameActive1(lvlkz, CorpName, name string) int                                                                    //проверка количество выполненых игр
	NumberQueueLvl(lvlkz, CorpName string) int                                                                                  //Номер катки по уровню
	ReadAll(lvlkz, CorpName string) (users models.Users)                                                                        //чтение игроков в очереди
	InsertQueue(dsmesid, wamesid, CorpName, name, nameMention, tip, lvlkz, timekz string, tgmesid, numkzN int)                  //внесение данных сбора
	MesidTgUpdate(mesidtg int, lvlkz string, corpname string)                                                                   //изменение ид сообщения в бд
	MesidDsUpdate(mesidds, lvlkz, corpname string)                                                                              //изменение ид сообщения в бд
	UpdateCompliteRS(lvlkz string, dsmesid string, tgmesid int, wamesid string, numberkz int, numberevent int, corpname string) //закрытие очереди кз
	CountNameQueue(name string) (countNames int)                                                                                //проверка игрока на наличие в очереди
	ElseTrue(name string) models.Sborkz                                                                                         //для выхода из очереди при другом старте
	DeleteQueue(name, lvlkz, CorpName string)                                                                                   //Если игрок покидает очередь
	UpdateMitutsQueue(name, CorpName string) models.Sborkz                                                                      //проверка хочет ли игрок продолжить время в очереди
	TimerInsert(dsmesid, dschatid string, tgmesid int, tgchatid int64, timed int)                                               //внесение ид сообщения в бд
	TimerDeleteMessage() []models.Timer                                                                                         //удаление из таймера
	P30Pl(lvlkz, CorpName, name string) int                                                                                     //+30 минут если в очереди
	UpdateTimedown(lvlkz, CorpName, name string)                                                                                //при нажатии плюса при остатке меньше 3х минут
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
	Shutdown() error
}

func (d *Db) Shutdown() error {
	err := d.Db.Close()
	return err
}

func (d *Db) СountName(name, lvlkz, corpName string) int {
	var countNames int
	row := d.Db.QueryRow("SELECT  COUNT(*) as count FROM sborkz WHERE name = ? AND lvlkz = ? AND corpname = ? AND active = 0",
		name, lvlkz, corpName)
	err := row.Scan(&countNames)
	if err != nil {
		d.log.Println("Ошибка проверки в очереди ли игрок  ", err)
		return d.СountName(name, lvlkz, corpName)
	}
	return countNames
}
func (d *Db) CountQueue(lvlkz, CorpName string) int { //проверка сколько игровок в очереди
	var count int
	row := d.Db.QueryRow("SELECT  COUNT(*) as count FROM sborkz WHERE lvlkz = ? AND corpname = ? AND active = 0",
		lvlkz, CorpName)
	err := row.Scan(&count)
	if err != nil {
		d.log.Println("Ошибка проверки количества игроков в очереди", err)
	}
	return count
}
func (d *Db) CountNumberNameActive1(lvlkz, CorpName, name string) int { // выковыриваем из базы значение количества походов на кз
	var countNumberNameActive1 int
	row := d.Db.QueryRow(
		"SELECT  COUNT(*) as count FROM sborkz WHERE lvlkz = ? AND corpname = ? AND name = ? AND active = 1",
		lvlkz, CorpName, name)
	err := row.Scan(&countNumberNameActive1)
	if err != nil {
		d.log.Println("Ошибка чтения количества игр", err)
	}
	return countNumberNameActive1
}
func (d *Db) NumberQueueLvl(lvlkz, CorpName string) int {
	var number int
	row := d.Db.QueryRow("SELECT  number FROM numkz WHERE lvlkz = ? AND corpname = ?",
		lvlkz, CorpName)
	err := row.Scan(&number)
	if err != nil {
		if err == sql.ErrNoRows {
			number = 0
			insertSmt := "INSERT INTO numkz(lvlkz, number,corpname) VALUES (?,?,?)"
			statement, err := d.Db.Prepare(insertSmt)
			if err != nil {
				d.log.Println("Ошибка подготовки внесения нумкз", err)
			}
			_, err = statement.Exec(lvlkz, number, CorpName)
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
	results, err := d.Db.Query("SELECT * FROM sborkz WHERE lvlkz = ? AND corpname = ? AND active = 0",
		lvlkz, CorpName)
	if err != nil {
		d.log.Println("Ошибка чтения активной очереди readall", err)
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
func (d *Db) InsertQueue(dsmesid, wamesid, CorpName, name, nameMention, tip, lvlkz, timekz string, tgmesid, numkzN int) {
	numevent := d.NumActiveEvent(CorpName)
	tm := time.Now()
	mdate := (tm.Format("2006-01-02"))
	mtime := (tm.Format("15:04"))
	timekzz, errs := strconv.Atoi(timekz)
	if timekzz == 0 {
		d.log.Panic(errs)
	}

	insertSborkztg1 := `INSERT INTO sborkz(corpname,name,mention,tip,dsmesid,tgmesid,wamesid,time,date,lvlkz,
                   numkzn,numberkz,numberevent,eventpoints,active,timedown) 
				VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
	_, err := d.Db.Exec(insertSborkztg1, CorpName, name, nameMention, tip, dsmesid, tgmesid,
		wamesid, mtime, mdate, lvlkz, numkzN, 0, numevent, 0, 0, timekzz)
	if err != nil {
		d.log.Println("Ошибка записи старта очереди", err)
	}
}
func (d *Db) MesidTgUpdate(mesidtg int, lvlkz string, corpname string) {
	_, err := d.Db.Exec(
		`update sborkz set tgmesid = ? where lvlkz = ? AND corpname = ? `,
		mesidtg, lvlkz, corpname)
	if err != nil {
		d.log.Println("Ошибка измениния месайди телеги", err)
	}
}
func (d *Db) MesidDsUpdate(mesidds, lvlkz, corpname string) {
	_, err := d.Db.Exec(
		`update sborkz set dsmesid = ? where lvlkz = ? AND corpname = ? `,
		mesidds, lvlkz, corpname)
	if err != nil {
		d.log.Println("Ошибка измениния месайди дискорда ", err)
	}
}
func (d *Db) UpdateCompliteRS(lvlkz string, dsmesid string, tgmesid int, wamesid string, numberkz int, numberevent int, corpname string) {
	_, err := d.Db.Exec(
		`update sborkz set active = 1,dsmesid = ?,tgmesid = ?,wamesid = ?,numberkz = ?,numberevent = ? 
				where lvlkz = ? AND corpname = ? AND active = 0`,
		dsmesid, tgmesid, wamesid, numberkz, numberevent, lvlkz, corpname)
	if err != nil {
		d.log.Println("Ошибка сохранения закрытия очереди", err)
	}
	_, err = d.Db.Exec(`update numkz set number=number+1 where lvlkz = ? AND corpname = ?`, lvlkz, corpname)
	if err != nil {
		d.log.Println("Ошибка обновления нумкзз", err)
	}
	if numberevent > 0 {
		_, err := d.Db.Exec(
			`update rsevent set number = number+1  where corpname = ? AND activeevent = 1`, corpname)
		if err != nil {
			d.log.Println("Ошибка обновления номера катки ивента ", err)
		}
	}
}
func (d *Db) NumberQueueEvents(CorpName string) int {
	var number int
	row := d.Db.QueryRow("SELECT  number FROM rsevent WHERE activeevent = 1 AND corpname = ? ", CorpName)
	err := row.Scan(&number)
	if err != nil {
		d.log.Println("Ошибка получения номера очереди с таблицы rsevent", err)
	}
	return number
}

func (d *Db) CountNameQueue(name string) (countNames int) { //проверяем есть ли игрок в других очередях
	row := d.Db.QueryRow("SELECT  COUNT(*) as count FROM sborkz WHERE name = ? AND active = 0", name)
	err := row.Scan(&countNames)
	if err != nil {
		d.log.Println("Ошибка проверки игрока в других очередях ", err)
	}
	return countNames
}
func (d *Db) CountNameQueueCorp(name, corp string) (countNames int) { //проверяем есть ли игрок в других очередях
	row := d.Db.QueryRow("SELECT  COUNT(*) as count FROM sborkz WHERE name = ? AND corpname = ? AND active = 0", name, corp)
	err := row.Scan(&countNames)
	if err != nil {
		d.log.Println("Ошибка проверки игрока в других очередях этой корпы ", err)
	}
	return countNames
}
func (d *Db) ElseTrue(name string) models.Sborkz {
	results, err := d.Db.Query("SELECT * FROM sborkz WHERE name = ? AND active = 0", name)
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
	_, err := d.Db.Exec("delete from sborkz where name = ? AND lvlkz = ? AND corpname = ? AND active = 0",
		name, lvlkz, CorpName)
	if err != nil {
		d.log.Println("Ошибка удаления из очереди ", err)
	}
}
func (d *Db) UpdateMitutsQueue(name, CorpName string) models.Sborkz {
	results, err := d.Db.Query("SELECT * FROM sborkz WHERE name = ? AND corpname = ? AND active = 0",
		name, CorpName)
	if err != nil {
		d.log.Println("Ошибка проверки игрока в очереди для функции (-+) ", err)
	}
	var t models.Sborkz
	for results.Next() {

		err = results.Scan(&t.Id, &t.Corpname, &t.Name, &t.Mention, &t.Tip, &t.Dsmesid, &t.Tgmesid, &t.Wamesid, &t.Time,
			&t.Date, &t.Lvlkz, &t.Numkzn, &t.Numberkz, &t.Numberevent, &t.Eventpoints, &t.Active, &t.Timedown)

		if t.Name == name && t.Timedown <= 3 {
			_, err := d.Db.Exec("update sborkz set timedown = timedown + 30 where active = 0 AND name = ? AND corpname = ?",
				t.Name, t.Corpname)
			if err != nil {
				d.log.Println("Ошибка обновления времени игрока в очереди для функции (-+) ", err)
			}
			return t
		}
	}
	return t
}
func (d *Db) TimerInsert(dsmesid, dschatid string, tgmesid int, tgchatid int64, timed int) {
	insertTimer := `INSERT INTO timer(dsmesid,dschatid,tgmesid,tgchatid,timed) VALUES (?,?,?,?,?)`
	_, err := d.Db.Exec(insertTimer, dsmesid, dschatid, tgmesid, tgchatid, timed)
	if err != nil {
		d.log.Println("Ошибка внесения в бд для удаления ", err)
	}
}
func (d *Db) TimerDeleteMessage() []models.Timer {
	_, err := d.Db.Exec(`update timer set timed = timed - 60`)
	if err != nil {
		d.log.Println("Ошибка удаления 60секунд", err)
	}

	results, err := d.Db.Query("SELECT * FROM timer WHERE timed < 60")
	if err != nil {
		d.log.Println("Ошибка чтения ид где меньше 60 секунд", err)
	}
	var timedown []models.Timer
	for results.Next() {
		var t models.Timer
		err = results.Scan(&t.Id, &t.Dsmesid, &t.Dschatid, &t.Tgmesid, &t.Tgchatid, &t.Timed)
		timedown = append(timedown, t)

		_, err = d.Db.Exec("delete from timer where  id = ? ", t.Id)
		if err != nil {
			d.log.Println("Ошибка удаления по ид с таблицы таймера", err)
		}
	}
	return timedown
}
func (d *Db) ReadMesIdDS(mesid string) (string, error) {
	results, err := d.Db.Query("SELECT lvlkz FROM sborkz WHERE dsmesid = ? AND active = 0", mesid)
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
	results, err := d.Db.Query("SELECT timedown FROM sborkz WHERE lvlkz = ? AND corpname = ? AND active = 0 AND name = ?",
		lvlkz, CorpName, name)
	if err != nil {
		d.log.Println("Ошибка получения оставшегося времени ", err)
	}
	for results.Next() {
		err = results.Scan(&timedown)
	}
	return timedown
}
func (d *Db) UpdateTimedown(lvlkz, CorpName, name string) {
	_, err := d.Db.Exec(`update sborkz set timedown = timedown+30 where lvlkz = ? AND corpname = ? AND name = ?`,
		lvlkz, CorpName, name)
	if err != nil {
		d.log.Println("Ошибка обновления времени ", err)
	}
}
func (d *Db) Queue(corpname string) []string {
	results, err := d.Db.Query("SELECT lvlkz FROM sborkz WHERE corpname = ? AND active = 0", corpname)
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
	results, err := d.Db.Query("SELECT dschannel,mesiddshelp FROM config")
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
	_, err := d.Db.Exec(`update sborkz set timedown = timedown - 1 where active = 0`)
	if err != nil {
		d.log.Println("Ошибка удаления минуты ", err)
	}

	results, err := d.Db.Query("SELECT * FROM sborkz WHERE active = 0")
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
	row := d.Db.QueryRow("SELECT  COUNT(*) as count FROM sborkz WHERE active = 0")
	err := row.Scan(&count)
	if err != nil {
		d.log.Println("Ошибка подсчета активных игроков в очередях", err)
	}
	var CorpActive0 []string
	if count > 0 {
		a := []string{}
		aa := []string{}
		results, err := d.Db.Query("SELECT corpname FROM sborkz WHERE active = 0")
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
	row := d.Db.QueryRow("SELECT  COUNT(*) as count FROM sborkz WHERE corpname = ? AND active = 0", corpname)
	err := row.Scan(&countCorp)
	if err != nil {
		d.log.Println("Ошибка получения активных очередей корпорации ", err)
	}
	if countCorp > 0 {
		results, err := d.Db.Query("SELECT * FROM sborkz WHERE corpname = ? AND active = 0", corpname)
		if err != nil {
			d.log.Println("Ошибка получения активных очередей корпорации2 ", err)
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
	results, err := d.Db.Query("SELECT * FROM sborkz WHERE dsmesid = ? AND active = 0", dsmesid)
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
			Avatar  string
		}{
			Mesid:   t.Dsmesid,
			Nameid:  "",
			Guildid: config.Config.Guildid,
		},
		Config: config,
		Option: models.Option{
			Edit:   true,
			Update: true},
	}

	return in
}
func (d *Db) MessageupdateTG(tgmesid int, config models.BotConfig) models.InMessage {
	results, err := d.Db.Query("SELECT * FROM sborkz WHERE tgmesid = ? AND active = 0", tgmesid)
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
		Option: models.Option{
			Edit:   true,
			Update: true},
	}
	return in
}
func (d *Db) ReadStatistic(Name string) string {
	num := 1
	str := "√ уровень кз время дата канал\n"
	tmp := ""
	results, err := d.Db.Query("SELECT * FROM sborkz WHERE name = ? AND active = 1", Name)
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
