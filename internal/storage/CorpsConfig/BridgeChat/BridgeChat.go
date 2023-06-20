package BridgeChat

import (
	"context"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"kz_bot/internal/models"
	"kz_bot/pkg/clientDB/postgresqlS"
)

type DB struct {
	client       postgresqlS.Client
	log          *logrus.Logger
	BridgeConfig []models.BridgeConfig
}

func NewDB(client postgresqlS.Client, log *logrus.Logger) *DB {
	d := &DB{
		client: client,
		log:    log,
	}
	d.BridgeConfig = d.DBReadBridgeConfig()
	return d
}
func (d *DB) DBReadBridgeConfig() []models.BridgeConfig {
	selectdata := "SELECT json FROM kzbot.bridgechat where typejson = $1"
	var data []byte
	err := d.client.QueryRow(context.Background(), selectdata, "bridgechat").Scan(&data)
	if err != nil {
		d.log.Println(err)
		return nil
	}
	var br []models.BridgeConfig
	err = json.Unmarshal(data, &br)
	if err != nil {
		d.log.Println(err)
		return nil
	}
	return br
}
func (d *DB) updateBridgeChat(br []models.BridgeConfig) {
	data, err := json.Marshal(br)
	if err != nil {
		d.log.Println(err)
		return
	}
	update := "UPDATE kzbot.bridgechat SET json = $1 where typejson = $2"
	exec, err := d.client.Exec(context.Background(), update, data, "bridgechat")
	if err != nil {
		return
	}
	if exec.RowsAffected() == 0 {
		insert := "INSERT INTO kzbot.bridgechat (typejson, json) VALUES ($1,$2)"
		_, err = d.client.Exec(context.Background(), insert, "bridgechat", data)
		if err != nil {
			d.log.Println(err)
		}
	}
}
