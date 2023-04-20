package ReservCopyPaste

import (
	"fmt"
	"kz_bot/internal/config"
	"kz_bot/internal/hades/ReservCopyPaste/ReservCopy"
	"kz_bot/pkg/utils"
	"time"
)

func RunReserv() {
	ticker := time.NewTicker(time.Hour)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			{
				go write()
			}
		}
	}
}
func write() {
	r := ReservCopy.NewReservDB()
	wr := r.NewMessageWriteToPostgres()
	p := NewReservPostgres(config.Instance)
	p.WriteToCloud(wr)
}
func LoadBackup() {
	p := NewReservPostgres(config.Instance)
	jsonData := p.ReadJson()
	r := ReservCopy.NewReservDB()
	r.WriteToSQLite(jsonData)
	fmt.Println("LoadBackup ok")
	utils.UpdateRun()
}
