package BridgeChat

import "kz_bot/internal/models"

func (b *Bridge) CacheNameBridge(nameRelay string) (bool, models.BridgeConfig) {
	if len(b.configs) != 0 {
		for _, config := range b.configs {
			if config.NameRelay == nameRelay {
				return true, config
			}
		}
	}
	return false, models.BridgeConfig{}
}
func (b *Bridge) AddNewBridgeConfig(br models.BridgeConfig) {
	b.configs[br.NameRelay] = br
	b.storage.BridgeConfig.InsertBridgeChat(br)
}
func (b *Bridge) AddBridgeConfig(br models.BridgeConfig) {
	a := b.configs[br.NameRelay]
	if len(br.ChannelDs) > 0 {
		a.ChannelDs = append(a.ChannelDs, br.ChannelDs...)
	}
	if len(br.ChannelTg) > 0 {
		a.ChannelTg = append(a.ChannelTg, br.ChannelTg...)
	}
	b.storage.BridgeConfig.UpdateBridgeChat(a)
	b.configs[br.NameRelay] = a
}
func (b *Bridge) CacheCheckChannelConfigDS(chatIdDs string) (bool, models.BridgeConfig) {
	for _, config := range b.configs {
		for _, ds := range config.ChannelDs {
			if ds.ChannelId == chatIdDs {
				return true, config
			}
		}
	}
	return false, models.BridgeConfig{}
}
func (b *Bridge) CacheCheckChannelConfigTg(chatIdTg string) (bool, models.BridgeConfig) {
	for _, config := range b.configs {
		for _, tg := range config.ChannelTg {
			if tg.ChannelId == chatIdTg {
				return true, config
			}
		}
	}
	return false, models.BridgeConfig{}
}
