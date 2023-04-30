package db

import (
	"context"
	"github.com/pkg/errors"
)

func (d *Repository) AddTgCorp(ctx context.Context, chatName string, chatid int64) error {
	insertConfig := `INSERT INTO kzbot.config
   		(corpname,dschannel,tgchannel,wachannel,mesiddshelp,mesidtghelp,delmescomplite,guildid)
		VALUES
		    ($1,$2,$3,$4,$5,$6,$7,$8)`
	_, err := d.client.Exec(ctx, insertConfig, chatName, "", chatid, "", "", 0, 0, "")
	if err != nil {
		d.log.Println("Ошибка внесения конфигурации ", err)
		return errors.New("Ошибка внесения конфигурации ")
	}
	return nil
}
func (d *Repository) AddDsCorp(ctx context.Context, chatName, chatid, guildid string) error {
	insertConfig := `INSERT INTO kzbot.config 
    	(corpname,dschannel,tgchannel,wachannel,mesiddshelp,mesidtghelp,delmescomplite,guildid)
		VALUES 
		    ($1,$2,$3,$4,$5,$6,$7,$8)`
	_, err := d.client.Exec(ctx, insertConfig, chatName, chatid, 0, "", "", 0, 0, guildid)
	if err != nil {
		d.log.Println("Ошибка внесения конфигурации ", err)
		return errors.New("Ошибка внесения конфигурации ")
	}
	return nil
}
func (d *Repository) AddWaCorp(ctx context.Context, chatName, chatid string) error {
	insertConfig := `INSERT INTO kzbot.config 
    	(corpname,dschannel,tgchannel,wachannel,mesiddshelp,mesidtghelp,delmescomplite,guildid)
			VALUES 
			    ($1,$2,$3,$4,$5,$6,$7,$8)`
	_, err := d.client.Exec(ctx, insertConfig, chatName, "", 0, chatid, "", 0, 0, "")
	if err != nil {
		d.log.Println("Ошибка внесения конфигурации ", err)
		return errors.New("Ошибка внесения конфигурации ")
	}
	return nil
}

func (d *Repository) AddGlobalDsCorp(ctx context.Context, chatName, chatid, guildid string) error {
	insertConfig := `INSERT INTO kzbot.global 
	    (CorpName,CorpNameMin,DsChannel,TgChannel,WaChannel,GuildId,Country)
		VALUES 
		    ($1,$2,$3,$4,$5,$6,$7)`
	_, err := d.client.Exec(ctx, insertConfig, chatName, "", chatid, 0, "", guildid, "")
	if err != nil {
		d.log.Println("Ошибка внесения конфигурации global ", err)
		return errors.New("Ошибка внесения конфигурации global ")
	}
	return nil
}
