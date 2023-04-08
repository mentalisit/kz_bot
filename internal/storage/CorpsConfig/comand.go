package CorpsConfig

import (
	"context"
	"fmt"
	"kz_bot/internal/storage/CorpsConfig/db"
)

func (c *Corps) AddTgCorpConfig(ctx context.Context, chatName string, chatid int64) error {
	c.log.Println(chatName, "Добавлена в конфиг корпораций ")
	err := c.db.AddTgCorp(ctx, chatName, chatid)
	if err != nil {
		return err
	}
	c.corp.AddCorp(chatName, "", chatid, "", 1, "", "", "")
	return nil
}
func (c *Corps) AddDsCorpConfig(ctx context.Context, chatName, chatid, guildid string) error {
	c.log.Println(chatName, "Добавлена в конфиг корпораций ")
	err := c.db.AddDsCorp(ctx, chatName, chatid, guildid)
	if err != nil {
		return err
	}
	c.corp.AddCorp(chatName, chatid, 0, "", 1, "", "", guildid)
	return nil
}
func (c *Corps) AddWaCorpConfig(ctx context.Context, chatName, chatid string) error {
	c.log.Println(chatName, "Добавлена в конфиг корпораций ")
	err := c.db.AddWaCorp(ctx, chatName, chatid)
	if err != nil {
		return err
	}
	c.corp.AddCorp(chatName, "", 0, chatid, 1, "", "", "")
	return nil
}

func (c *Corps) DeleteTg(ctx context.Context, chatid int64) error {
	err := c.db.DeleteTgChannel(ctx, chatid)
	if err != nil {
		return err
	}
	return nil
}
func (c *Corps) DeleteDs(ctx context.Context, chatid string) error {
	err := c.db.DeleteDsChannel(ctx, chatid)
	if err != nil {
		return err
	}
	return nil
}
func (c *Corps) DeleteWa(ctx context.Context, chatid string) error {
	err := c.db.DeleteWaChannel(ctx, chatid)
	if err != nil {
		return err
	}
	return nil
}

func (c *Corps) ReadCorps() {
	listCorp := c.db.ReadBotCorpConfig(context.Background())
	var corp []string
	for _, t := range listCorp {
		c.corp.AddCorp(t.CorpName, t.DsChannel, t.TgChannel, t.WaChannel,
			t.DelMesComplite, t.MesidDsHelp, t.Country, t.GuildId)

		corp = append(corp, fmt.Sprintf(" %s, ", t.CorpName))
	}
	fmt.Printf("Конфиг корпораций:%s\n", corp)
}

func (c *Corps) AutoHelpUpdateMesid(ctx context.Context, newMesidHelp, dschannel string) error {
	err := c.db.AutoHelpUpdateMesid(ctx, newMesidHelp, dschannel)
	if err != nil {
		return err
	}
	return nil
}
func (c *Corps) AutoHelp() []db.ConfigCorp {
	return c.db.AutoHelp()
}
