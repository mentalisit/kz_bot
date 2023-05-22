package Relay

import "kz_bot/internal/models"

var R = NewR()

func NewR() *rl {
	var arr rl
	return &arr
}

//var BlackListNamesId []string

type rl []models.RelayConfig

type RelayC struct{}

func NewRelayC() *RelayC {
	return &RelayC{}
}

func (r *RelayC) ReloadConfig() {
	*R = *NewR()
}
func (r *RelayC) AddCorp(RelayName, RelayAlias, GuildName string, DsChannel string, TgChannel int64, WaChannel string, guildid string, Country string, prefix string) {
	relay := models.RelayConfig{
		RelayName:  RelayName,
		RelayAlias: RelayAlias,
		GuildName:  GuildName,
		DsChannel:  DsChannel,
		TgChannel:  TgChannel,
		WaChannel:  WaChannel,
		GuildId:    guildid,
		Country:    Country,
		Prefix:     prefix,
	}
	*R = append(*R, relay)
}
func (r *RelayC) CheckChannelConfigDS(chatid string) (channelGood bool, config models.RelayConfig) {
	if chatid != "" {
		for _, pp := range *R {
			if chatid == pp.DsChannel {
				return true, pp
			}
		}
	}
	return false, models.RelayConfig{}
}
func (r *RelayC) CheckChannelConfigTG(chatid int64) (channelGood bool, config models.RelayConfig) {
	if chatid != 0 {
		for _, pp := range *R {
			if chatid == pp.TgChannel {
				return true, pp
			}
		}
	}
	return false, models.RelayConfig{}
}
func (r *RelayC) CheckChannelConfigWA(chatid string) (channelGood bool, config models.RelayConfig) {
	if chatid != "" {
		for _, pp := range *R {
			if chatid == pp.WaChannel {
				return true, pp
			}
		}
	}
	return false, models.RelayConfig{}
}

func (r *RelayC) ReadAllChannel() (chatDS []string, chatTG []int64, chatWA []string) {
	for _, pp := range *R {
		if pp.DsChannel != "" {
			chatDS = append(chatDS, pp.DsChannel)
		}
	}
	for _, pp := range *R {
		if pp.TgChannel != 0 {
			chatTG = append(chatTG, pp.TgChannel)
		}
	}
	for _, pp := range *R {
		if pp.WaChannel != "" {
			chatWA = append(chatWA, pp.WaChannel)
		}
	}
	return chatDS, chatTG, chatWA
}
