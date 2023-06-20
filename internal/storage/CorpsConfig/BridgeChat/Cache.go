package BridgeChat

import "kz_bot/internal/models"

func (d *DB) CacheNameBridge(nameRelay string) (bool, models.BridgeConfig) {
	if len(d.BridgeConfig) != 0 {
		for _, config := range d.BridgeConfig {
			if config.NameRelay == nameRelay {
				return true, config
			}
		}
	}
	return false, models.BridgeConfig{}
}
func (d *DB) AddNewBridgeConfig(br models.BridgeConfig) {
	d.BridgeConfig = append(d.BridgeConfig, br)
	d.updateBridgeChat(d.BridgeConfig)
}
func (d *DB) AddBridgeConfigAppend(br models.BridgeConfig) {
	var Bridge []models.BridgeConfig
	for _, config := range d.BridgeConfig {
		if br.HostRelay != config.HostRelay {
			Bridge = append(Bridge, config)
		}
		if br.HostRelay == config.HostRelay {
			newConfig := config
			if len(br.ChannelTg) > 0 {
				newConfig.ChannelTg = append(newConfig.ChannelTg, br.ChannelTg[0])
			}
			if len(br.ChannelDs) > 0 {
				newConfig.ChannelDs = append(newConfig.ChannelDs, br.ChannelDs[0])
			}
			Bridge = append(Bridge, newConfig)
		}
	}
	d.BridgeConfig = Bridge
	d.updateBridgeChat(d.BridgeConfig)
}

func (d *DB) CacheCheckChannelConfigDS(dschat string) (bool, models.BridgeConfig) {
	for _, config := range d.BridgeConfig {
		for _, ds := range config.ChannelDs {
			if ds.ChannelId == dschat {
				return true, config
			}
		}
	}
	return false, models.BridgeConfig{}
}
func (d *DB) CacheCheckChannelConfigTg(tgchat int64) (bool, models.BridgeConfig) {
	for _, config := range d.BridgeConfig {
		for _, tg := range config.ChannelTg {
			if tg.ChannelId == tgchat {
				return true, config
			}
		}
	}
	return false, models.BridgeConfig{}
}
