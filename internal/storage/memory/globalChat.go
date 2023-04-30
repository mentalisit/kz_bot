package memory

var G = NewG()

func NewG() *MemoryGlobal {
	var arr MemoryGlobal
	return &arr
}

var BlackListNamesId []string

type MemoryGlobal []ConfigGlobal

type CorpConfigGl struct{}

func (c CorpConfigGl) ReloadConfig() {
	*P = *New()
}
func (c CorpConfigGl) AddCorp(CorpName string, DsChannel string, TgChannel int64, WaChannel string, Country string, guildid string) {
	corps := ConfigGlobal{
		CorpName:  CorpName,
		DsChannel: DsChannel,
		TgChannel: TgChannel,
		WaChannel: WaChannel,
		GuildId:   guildid,
		Country:   Country,
	}
	*G = append(*G, corps)
}
func (c CorpConfigGl) CheckChannelConfigDS(chatid string) (channelGood bool, config ConfigGlobal) {
	if chatid != "" {
		for _, pp := range *G {
			if chatid == pp.DsChannel {
				return true, pp
			}
		}
	}
	return false, ConfigGlobal{}
}

//	func (c CorpConfigGlobal) CheckChannelConfigWA(chatid string) (channelGood bool, config CorpporationConfig) {
//		if chatid != "" {
//			for _, pp := range *P {
//				if chatid == pp.WaChannel {
//					channelGood = true
//					config = pp
//					break
//				}
//			}
//		}
//		return channelGood, config
//	}
//
//	func (c CorpConfigGlobal) CheckChannelConfigTG(chatid int64) (channelGood bool, config CorpporationConfig) {
//		if chatid != 0 {
//			for _, pp := range *P {
//				if chatid == pp.TgChannel {
//					channelGood = true
//					config = pp
//					break
//				}
//			}
//		}
//		return channelGood, config
//	}
func (c CorpConfigGl) ReadAllChannel() (chatDS []string, chatTG []int64, chatWA []string) {
	for _, pp := range *P {
		if pp.DsChannel != "" {
			chatDS = append(chatDS, pp.DsChannel)
		}
	}
	for _, pp := range *P {
		if pp.TgChannel != 0 {
			chatTG = append(chatTG, pp.TgChannel)
		}
	}
	for _, pp := range *P {
		if pp.WaChannel != "" {
			chatWA = append(chatWA, pp.WaChannel)
		}
	}
	return chatDS, chatTG, chatWA
}
