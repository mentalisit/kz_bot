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
func (c CorpConfigGl) CheckChannelConfigTG(chatid int64) (channelGood bool, config ConfigGlobal) {
	if chatid != 0 {
		for _, pp := range *G {
			if chatid == pp.TgChannel {
				return true, pp
			}
		}
	}
	return false, ConfigGlobal{}
}
func (c CorpConfigGl) CheckChannelConfigWA(chatid string) (channelGood bool, config ConfigGlobal) {
	if chatid != "" {
		for _, pp := range *G {
			if chatid == pp.WaChannel {
				return true, pp
			}
		}
	}
	return false, ConfigGlobal{}
}

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
