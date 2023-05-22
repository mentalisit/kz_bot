package Relay

import (
	"context"
	"kz_bot/internal/models"
)

func (s *RelayStorage) ReadAllRelay() []models.RelayConfig {
	read :=
		`SELECT * FROM kzbot.relay`
	results, err := s.client.Query(context.Background(), read)
	if err != nil {
		s.log.Println("Ошибка чтения крнфигурации relay ", err)
	}
	var corps []models.RelayConfig

	for results.Next() {
		var t models.RelayConfig
		err = results.Scan(&t.Id, &t.RelayName, &t.RelayAlias, &t.GuildName,
			&t.DsChannel, &t.TgChannel, &t.WaChannel,
			&t.GuildId, &t.Country, &t.Prefix)
		corps = append(corps, t)
	}
	return corps
}
