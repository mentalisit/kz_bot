package db

import "context"

func (d *Repository) DeleteTgChannel(ctx context.Context, chatid int64) error {
	_, err := d.client.Exec(ctx, "delete from kzbot.config where tgchannel = $1 ", chatid)
	if err != nil {
		d.log.Println("Ошибка удаления с бд корп телеги", err)
		return err
	}
	return nil
}
func (d *Repository) DeleteDsChannel(ctx context.Context, chatid string) error {
	_, err := d.client.Exec(ctx, "delete from kzbot.config where dschannel = $1 ", chatid)
	if err != nil {
		d.log.Println("Ошибка удаления с бд корп дискорд", err)
		return err
	}
	return nil
}
func (d *Repository) DeleteWaChannel(ctx context.Context, chatid string) error {
	_, err := d.client.Exec(ctx, "delete from kzbot.config where wachannel = $1 ", chatid)
	if err != nil {
		d.log.Println("Ошибка удаления с бд корп wa", err)
		return err
	}
	return nil
}

func (d *Repository) DeleteGlobalDsChannel(ctx context.Context, chatid string) error {
	_, err := d.client.Exec(ctx, "delete from kzbot.global where DsChannel = $1 ", chatid)
	if err != nil {
		d.log.Println("Ошибка удаления с бд корп global дискорд", err)
		return err
	}
	return nil
}
