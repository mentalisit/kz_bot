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
func (b *Bridge) CacheCheckChannelConfigTg(chatIdTg int64) (bool, models.BridgeConfig) {
	for _, config := range b.configs {
		for _, tg := range config.ChannelTg {
			if tg.ChannelId == chatIdTg {
				return true, config
			}
		}
	}
	return false, models.BridgeConfig{}
}
