package Relay

import (
	"context"
	"errors"
	"kz_bot/internal/models"
)

func (s *RelayStorage) AddRelay(c models.RelayConfig) error {
	insertConfig := `INSERT INTO kzbot.relay 
	    ("RelayName","RelayAlias","GuildName","DsChannel","TgChannel","WaChannel","GuildId","Country","Prefix")
		VALUES 
		    ($1,$2,$3,$4,$5,$6,$7,$8,$9)`
	_, err := s.client.Exec(context.Background(), insertConfig,
		c.RelayName, c.RelayAlias, c.GuildName,
		c.DsChannel, c.TgChannel, c.WaChannel,
		c.GuildId, c.Country, c.Prefix)
	if err != nil {
		s.log.Println("Ошибка внесения конфигурации relay ", err)
		return errors.New("Ошибка внесения конфигурации relay " + err.Error())
	}
	return nil
}
