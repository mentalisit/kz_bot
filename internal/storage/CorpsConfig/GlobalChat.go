package CorpsConfig

import (
	"context"
	"encoding/json"
	"fmt"
	"kz_bot/internal/storage/memory"
)

func (c *Corps) AddGlobalDsCorp(ctx context.Context, chatName, chatid, guildid string) error {
	c.log.Println(chatName, "Добавлена в конфиг корпораций ")
	err := c.db.AddGlobalDsCorp(ctx, chatName, chatid, guildid)
	if err != nil {
		return err
	}
	c.global.AddCorp(chatName, chatid, 0, "", "", guildid)
	return nil
}
func (c *Corps) DeleteGlobalDs(ctx context.Context, chatid string) error {
	err := c.db.DeleteGlobalDsChannel(ctx, chatid)
	if err != nil {
		return err
	}
	c.global.ReloadConfig()
	return nil
}
func (c *Corps) ReadGlobalCorps() {
	globalCorp := c.db.ReadGlobalCorpConfig(context.Background())
	var corp []string
	for _, t := range globalCorp {
		c.global.AddCorp(t.CorpName, t.DsChannel, t.TgChannel, t.WaChannel, t.Country, t.GuildId)

		corp = append(corp, fmt.Sprintf(" %s, ", t.CorpName))
	}
	fmt.Printf("Конфиг GlobalChat:%s\n", corp)
}
func (c *Corps) UpdateJsonBlackList(data []byte) {
	update := "UPDATE kzbot.blacklist SET names = $1"
	exec, err := c.client.Exec(context.Background(), update, data)
	if err != nil {
		return
	}
	if exec.RowsAffected() == 0 {
		insert := "INSERT INTO kzbot.blacklist (names) VALUES ($1)"
		_, err := c.client.Exec(context.Background(), insert, data)
		if err != nil {
			fmt.Println(err)
		}
	}
}
func (c *Corps) ReadJsonBlackList() {
	selectdata := "SELECT names FROM kzbot.blacklist"
	var data []byte
	err := c.client.QueryRow(context.Background(), selectdata).Scan(&data)
	if err != nil {
		fmt.Println(err)
	}
	var m []string
	err = json.Unmarshal(data, &m)
	if err != nil {
		fmt.Println(err)
		//return
	}
	memory.BlackListNamesId = m
}
