package CorpsConfig

//func (c *Corps) AddTgCorpConfig(ctx context.Context, chatName string, chatid int64) error {
//	c.log.Println(chatName, "Добавлена в конфиг корпораций ")
//	err := c.db.AddTgCorp(ctx, chatName, chatid)
//	if err != nil {
//		return err
//	}
//	//c.corp.AddCorp(chatName, "", chatid, "", 1, "", "", "")
//	return nil
//}

//	func (c *Corps) AddDsCorpConfig(ctx context.Context, chatName, chatid, guildid string) error {
//		c.log.Println(chatName, "Добавлена в конфиг корпораций ")
//		err := c.db.AddDsCorp(ctx, chatName, chatid, guildid)
//		if err != nil {
//			return err
//		}
//		//c.corp.AddCorp(chatName, chatid, 0, "", 1, "", "", guildid)
//		return nil
//	}
//func (c *Corps) AddWaCorpConfig(ctx context.Context, chatName, chatid string) error {
//	c.log.Println(chatName, "Добавлена в конфиг корпораций ")
//	err := c.db.AddWaCorp(ctx, chatName, chatid)
//	if err != nil {
//		return err
//	}
//	//c.corp.AddCorp(chatName, "", 0, chatid, 1, "", "", "")
//	return nil
//}

//func (c *Corps) DeleteTg(ctx context.Context, chatid int64) error {
//	err := c.db.DeleteTgChannel(ctx, chatid)
//	if err != nil {
//		return err
//	}
//	return nil
//}
//func (c *Corps) DeleteDs(ctx context.Context, chatid string) error {
//	err := c.db.DeleteDsChannel(ctx, chatid)
//	if err != nil {
//		return err
//	}
//	return nil
//}
//func (c *Corps) DeleteWa(ctx context.Context, chatid string) error {
//	err := c.db.DeleteWaChannel(ctx, chatid)
//	if err != nil {
//		return err
//	}
//	return nil
//}

//func (c *Corps) ReadCorps() []models.CorporationConfig {
//	listCorp := c.db.ReadBotCorpConfig(context.Background())
//	var cc []models.CorporationConfig
//	for _, corp := range listCorp {
//		c := models.CorporationConfig{
//			CorpName:       corp.CorpName,
//			DsChannel:      corp.DsChannel,
//			TgChannel:      corp.TgChannel,
//			WaChannel:      corp.WaChannel,
//			Country:        corp.Country,
//			DelMesComplite: corp.DelMesComplite,
//			MesidDsHelp:    corp.MesidDsHelp,
//			Guildid:        corp.GuildId,
//		}
//
//		cc = append(cc, c)
//	}
//	return cc
//}

//func (c *Corps) AutoHelpUpdateMesid(ctx context.Context, newMesidHelp, dschannel string) error {
//	err := c.db.AutoHelpUpdateMesid(ctx, newMesidHelp, dschannel)
//	if err != nil {
//		return err
//	}
//	return nil
//}
//func (c *Corps) AutoHelp() []db.ConfigCorp {
//	return c.db.AutoHelp()
//}
