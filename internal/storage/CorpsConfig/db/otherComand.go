package db

//func (d *Repository) ReadBotCorpConfig(ctx context.Context) []ConfigCorp {
//	read :=
//		`SELECT * FROM kzbot.config`
//	results, err := d.client.Query(ctx, read)
//	if err != nil {
//		d.log.Println("Ошибка чтения крнфигурации корпораций", err)
//	}
//	var corps []ConfigCorp
//
//	for results.Next() {
//		var t ConfigCorp
//		err = results.Scan(&t.Id, &t.CorpName, &t.DsChannel, &t.TgChannel, &t.WaChannel,
//			&t.MesidDsHelp, &t.MesidTgHelp, &t.DelMesComplite, &t.GuildId, &t.Country)
//		fmt.Printf("ReadBotCorpConfig: %s \n", t.CorpName)
//		corps = append(corps, t)
//	}
//
//	return corps
//}

//func (d *Repository) AutoHelpUpdateMesid(ctx context.Context, newMesidHelp, dschannel string) error {
//	updateString :=
//		`update kzbot.config
//		set "mesiddshelp"=$1
//		where "dschannel"=$2`
//	_, err := d.client.Exec(ctx, updateString, newMesidHelp, dschannel)
//	if err != nil {
//		d.log.Println("ОШибка обновления месИд для автосправки ", err)
//		return err
//	}
//	return nil
//}

//func (d *Repository) AutoHelp() []ConfigCorp {
//	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
//	defer cancel()
//	sel := "SELECT dschannel,mesiddshelp FROM kzbot.config"
//	results, err := d.client.Query(ctx, sel)
//	if err != nil {
//		d.log.Println("Ошибка получения автосправки с бд", err)
//	}
//	h := ConfigCorp{}
//	var a []ConfigCorp
//	for results.Next() {
//		err = results.Scan(&h.DsChannel, &h.MesidDsHelp)
//		a = append(a, h)
//	}
//	return a
//}
