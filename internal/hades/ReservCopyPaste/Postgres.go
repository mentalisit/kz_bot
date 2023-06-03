package ReservCopyPaste

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"kz_bot/internal/config"
	"kz_bot/pkg/clientDB/postgresqlS"
	"log"
)

type ReservPostgres struct {
	client postgresqlS.Client
}

func NewReservPostgres(cfg *config.ConfigBot) *ReservPostgres {
	p := &ReservPostgres{}
	client, err := postgresqlS.NewClient(context.Background(), logrus.New(), 5, cfg)
	if err != nil {
		log.Println("Ошибка подключения к облачной ДБ " + err.Error())
		return p
	}
	p.client = client
	return p
}
func (p *ReservPostgres) WriteToCloud(data []byte) {
	update := "UPDATE kzbot.reserv SET data = $1"
	exec, err := p.client.Exec(context.Background(), update, data)
	if err != nil {
		return
	}
	if exec.RowsAffected() == 0 {
		insert := "INSERT INTO kzbot.reserv (data) VALUES ($1)"
		_, err := p.client.Exec(context.Background(), insert, data)
		if err != nil {
			fmt.Println(err)
		}
	}
}
func (p *ReservPostgres) ReadJson() []byte {
	selectdata := "SELECT data FROM kzbot.reserv"
	var data []byte
	err := p.client.QueryRow(context.Background(), selectdata).Scan(&data)
	if err != nil {
		fmt.Println(err)
	}
	return data
}
