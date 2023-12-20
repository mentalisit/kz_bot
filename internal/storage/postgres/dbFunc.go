package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"kz_bot/internal/models"
	"kz_bot/pkg/utils"
	"strconv"
	"time"
)

func (d *Db) ReadAll(ctx context.Context, lvlkz, CorpName string) (users models.Users) {
	if d.debug {
		fmt.Println("ReadAll lvlkz, CorpName", lvlkz, CorpName)
	}
	u := models.Users{
		User1: models.Sborkz{},
		User2: models.Sborkz{},
		User3: models.Sborkz{},
		User4: models.Sborkz{},
	}
	user := 1
	sel := "SELECT * FROM kzbot.sborkz WHERE lvlkz = $1 AND corpname = $2 AND active = 0"
	results, err := d.db.Query(ctx, sel, lvlkz, CorpName)
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
	if d.debug {
		fmt.Println("ReadAll u", u.User1.Name, u.User2.Name, u.User3.Name, u.User4.Name)
	}
	return u
}
func (d *Db) InsertQueue(ctx context.Context, dsmesid, wamesid, CorpName, name, nameMention, tip, lvlkz, timekz string, tgmesid, numkzN int) {
	numevent := 0 // d.NumActiveEvent(CorpName)
	tm := time.Now()
	mdate := (tm.Format("2006-01-02"))
	mtime := (tm.Format("15:04"))
	if d.debug {
		fmt.Println("InsertQueue", CorpName, name, lvlkz, timekz)
	}
	timekzz, errs := strconv.Atoi(timekz)
	if timekzz == 0 {
		d.log.Println("Ошибка инсЕрта время кз не может быть нолем ", name, timekz, errs)
		timekzz = 1
	}
	ctx = context.Background()

	insertSborkztg1 := `INSERT INTO kzbot.sborkz(corpname,name,mention,tip,dsmesid,tgmesid,wamesid,time,date,lvlkz,
                   numkzn,numberkz,numberevent,eventpoints,active,timedown) 
				VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16)`
	_, err := d.db.Exec(ctx, insertSborkztg1, CorpName, name, nameMention, tip, dsmesid, tgmesid,
		wamesid, mtime, mdate, lvlkz, numkzN, 0, numevent, 0, 0, timekzz)
	if err != nil {
		d.log.Println("Ошибка записи старта очереди", err)
	}
}

func (d *Db) ElseTrue(ctx context.Context, name string) []models.Sborkz {
	if d.debug {
		fmt.Println("ElseTrue", name)
	}
	sel := "SELECT * FROM kzbot.sborkz WHERE name = $1 AND active = 0"
	results, err := d.db.Query(ctx, sel, name)
	if err != nil {
		d.log.Println("Ошибка извлечения игрока с других очередей ", err)
	}
	var tt []models.Sborkz
	var t models.Sborkz
	for results.Next() {
		err = results.Scan(&t.Id, &t.Corpname, &t.Name, &t.Mention, &t.Tip, &t.Dsmesid, &t.Tgmesid, &t.Wamesid, &t.Time, &t.Date, &t.Lvlkz, &t.Numkzn, &t.Numberkz, &t.Numberevent, &t.Eventpoints, &t.Active, &t.Timedown)
		tt = append(tt, t)
	}
	if d.debug {
		fmt.Println("ElseTrue", name, t)
	}
	return tt
}
func (d *Db) DeleteQueue(ctx context.Context, name, lvlkz, CorpName string) {
	if d.debug {
		fmt.Println("DeleteQueue", name, lvlkz, CorpName)
	}
	del := "delete from kzbot.sborkz where name = $1 AND lvlkz = $2 AND corpname = $3 AND active = 0"
	_, err := d.db.Exec(ctx, del, name, lvlkz, CorpName)
	if err != nil {
		d.log.Println("Ошибка удаления из очереди ", err)
	}
}

func (d *Db) ReadMesIdDS(ctx context.Context, mesid string) (string, error) {
	if d.debug {
		fmt.Println("ReadMesIdDS", mesid)
	}
	sel := "SELECT lvlkz FROM kzbot.sborkz WHERE dsmesid = $1 AND active = 0"
	results, err := d.db.Query(ctx, sel, mesid)
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
	a = utils.RemoveDuplicateElementString(a)
	if d.debug {
		fmt.Println("ReadMesIdDS", a)
	}
	if len(a) > 0 {
		dsmesid = a[0]
		return dsmesid, nil
	} else {
		return "", err
	}
}

func (d *Db) P30Pl(ctx context.Context, lvlkz, CorpName, name string) int {
	if d.debug {
		fmt.Println("P30Pl", lvlkz, CorpName, name)
	}
	var timedown int
	sel := "SELECT timedown FROM kzbot.sborkz WHERE lvlkz = $1 AND corpname = $2 AND active = 0 AND name = $3"
	results, err := d.db.Query(ctx, sel, lvlkz, CorpName, name)
	if err != nil {
		d.log.Println("Ошибка получения оставшегося времени ", err)
	}
	for results.Next() {
		err = results.Scan(&timedown)
	}
	if d.debug {
		fmt.Println("P30Pl", timedown)
	}
	return timedown
}
func (d *Db) UpdateTimedown(ctx context.Context, lvlkz, CorpName, name string) {
	if d.debug {
		fmt.Println("UpdateTimedown", lvlkz, CorpName, name)
	}
	upd := `update kzbot.sborkz set timedown = timedown+30 where lvlkz = $1 AND corpname = $2 AND name = $3`
	_, err := d.db.Exec(ctx, upd, lvlkz, CorpName, name)
	if err != nil {
		d.log.Println("Ошибка обновления времени ", err)
	}
}
func (d *Db) Queue(ctx context.Context, corpname string) []string {
	if d.debug {
		fmt.Println("Queue corpname", corpname)
	}
	sel := "SELECT lvlkz FROM kzbot.sborkz WHERE corpname = $1 AND active = 0"
	results, err := d.db.Query(ctx, sel, corpname)
	if err != nil {
		d.log.Println("Ошибка чтения левелов для очереди", err)
	}
	var lvl []string
	for results.Next() {
		var t models.Sborkz
		err = results.Scan(&t.Lvlkz)

		lvl = append(lvl, t.Lvlkz)

	}
	if d.debug {
		fmt.Println("Queue lvl", lvl)
	}

	return lvl
}

func (d *Db) OneMinutsTimer(ctx context.Context) []string {
	var count int //количество активных игроков
	sel := "SELECT  COUNT(*) as count FROM kzbot.sborkz WHERE active = 0"
	row := d.db.QueryRow(ctx, sel)
	err := row.Scan(&count)
	if err != nil {
		d.log.Println("Ошибка подсчета активных игроков в очередях", err)
	}
	var CorpActive0 []string
	if count > 0 {
		a := []string{}
		aa := []string{}
		selC := "SELECT corpname FROM kzbot.sborkz WHERE active = 0"
		results, err1 := d.db.Query(ctx, selC)
		if err1 != nil {
			d.log.Println("Ошибка чтения корпораций где есть активные очереди ", err)
		}
		var corpname string // ищим корпорации
		for results.Next() {
			err = results.Scan(&corpname)
			a = append(a, corpname)
		}
		a = utils.RemoveDuplicateElementString(a)

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
	if d.debug {
		if len(CorpActive0) > 0 {
			fmt.Println("OneMinutsTimer", CorpActive0)
		}
	}
	return CorpActive0
}
func (d *Db) MessageUpdateMin(ctx context.Context, corpname string) ([]string, []int, []string) {
	if d.debug {
		fmt.Println("MessageUpdateMin", corpname)
	}
	var countCorp int
	ds := []string{}
	tg := []int{}
	wa := []string{}
	sel := "SELECT  COUNT(*) as count FROM kzbot.sborkz WHERE corpname = $1 AND active = 0"
	row := d.db.QueryRow(ctx, sel, corpname)
	err := row.Scan(&countCorp)
	if err != nil {
		d.log.Println("Ошибка получения активных очередей корпорации ", err)
	}
	if countCorp > 0 {
		selS := "SELECT * FROM kzbot.sborkz WHERE corpname = $1 AND active = 0"
		results, err1 := d.db.Query(ctx, selS, corpname)
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
	ds = utils.RemoveDuplicateElementString(ds)
	tg = utils.RemoveDuplicateElementInt(tg)
	wa = utils.RemoveDuplicateElementString(wa)
	if d.debug {
		fmt.Println("MessageUpdateMin", "ds", ds, "tg", tg, "wa", wa)
	}
	return ds, tg, wa
}
func (d *Db) MessageupdateDS(ctx context.Context, dsmesid string, config models.CorporationConfig) models.InMessage {
	if d.debug {
		fmt.Println("MessageupdateDS", dsmesid, config.CorpName)
	}
	sel := "SELECT * FROM kzbot.sborkz WHERE dsmesid = $1 AND active = 0"
	results, err := d.db.Query(ctx, sel, dsmesid)
	if err != nil {
		d.log.Println(err)
	}
	var t models.Sborkz
	for results.Next() {
		err = results.Scan(&t.Id, &t.Corpname, &t.Name, &t.Mention, &t.Tip, &t.Dsmesid, &t.Tgmesid, &t.Wamesid, &t.Time, &t.Date, &t.Lvlkz, &t.Numkzn, &t.Numberkz, &t.Numberevent, &t.Eventpoints, &t.Active, &t.Timedown)
	}
	in := models.InMessage{
		Tip:         "ds",
		Name:        t.Name,
		NameMention: t.Mention,
		Lvlkz:       t.Lvlkz,
		Timekz:      string(t.Timedown),
		Ds: struct {
			Mesid   string
			Nameid  string
			Guildid string
			Avatar  string
		}{
			Mesid:   t.Dsmesid,
			Nameid:  "",
			Guildid: config.Guildid,
		},
		Config: config,
		Option: models.Option{
			Edit:   true,
			Update: true},
	}
	return in

}
func (d *Db) MessageupdateTG(ctx context.Context, tgmesid int, config models.CorporationConfig) models.InMessage {
	if d.debug {
		fmt.Println("MessageupdateTG", tgmesid, config.CorpName)
	}
	sel := "SELECT * FROM kzbot.sborkz WHERE tgmesid = $1 AND active = 0"
	results, err := d.db.Query(ctx, sel, tgmesid)
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
		Timekz:      string(t.Timedown),
		Tg: struct {
			Mesid int
			//Nameid int64
		}{
			Mesid: t.Tgmesid,
			//Nameid: 0
		},
		Config: config,
		Option: models.Option{
			Edit:   true,
			Update: true},
	}
	return in
}
func (d *Db) NumberQueueLvl(ctx context.Context, lvlkz, CorpName string) (int, error) {
	//lvlkz, errc := strconv.Atoi(lvlkzs)
	//if errc != nil {
	//d.log.Println("ошибка преобразования в инт lkz lks", lvlkz, lvlkzs)
	//return 0, errc
	//}
	if d.debug {
		fmt.Println("NumberQueueLvl", lvlkz, CorpName)
	}
	var number int
	sel := "SELECT  number FROM kzbot.numkz WHERE lvlkz = $1 AND corpname = $2"
	row := d.db.QueryRow(ctx, sel, lvlkz, CorpName)
	err := row.Scan(&number)
	if err != nil {
		if err == pgx.ErrNoRows {
			number = 0
			insertSmt := "INSERT INTO kzbot.numkz(lvlkz, number,corpname) VALUES ($1,$2,$3)"
			_, err = d.db.Exec(ctx, insertSmt, lvlkz, number, CorpName)
			if err != nil {
				d.log.Println("Ошибка внесения нумкз", err)
			}
			return number + 1, nil
		} else {
			d.log.Println("Ошибка чтения нумкз", err)
			return 0, err
		}
	}
	if d.debug {
		fmt.Println("NumberQueueLvl", number)
	}
	return number + 1, nil
}
