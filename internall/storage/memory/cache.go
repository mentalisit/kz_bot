package memory

var P = New()

func New() *Memory {
	var arr Memory
	return &arr
}

type Memory []CorpporationConfig

type CorpConfig struct{}

func (c CorpConfig) ReloadConfig() {
	*P = *New()
}
func (c CorpConfig) AddCorp(CorpName string, DsChannel string, TgChannel int64, WaChannel string, DelMesComplite int, mesiddshelp string, Country string, guildid string) {
	corpConfig := CorpporationConfig{
		Type:           0xff,
		CorpName:       CorpName,
		DsChannel:      DsChannel,
		TgChannel:      TgChannel,
		WaChannel:      WaChannel,
		Country:        Country,
		DelMesComplite: DelMesComplite,
		MesidDsHelp:    mesiddshelp,
		Primer:         "",
		Guildid:        guildid,
	}
	*P = append(*P, corpConfig)
}
func (c CorpConfig) CheckChannelConfigDS(chatid string) (channelGood bool, config CorpporationConfig) {
	if chatid != "" {
		for _, pp := range *P {
			if chatid == pp.DsChannel {
				channelGood = true
				config = pp
				break
			}
		}
	}
	return channelGood, config
}
func (c CorpConfig) CheckChannelConfigWA(chatid string) (channelGood bool, config CorpporationConfig) {
	if chatid != "" {
		for _, pp := range *P {
			if chatid == pp.WaChannel {
				channelGood = true
				config = pp
				break
			}
		}
	}
	return channelGood, config
}
func (c CorpConfig) CheckChannelConfigTG(chatid int64) (channelGood bool, config CorpporationConfig) {
	if chatid != 0 {
		for _, pp := range *P {
			if chatid == pp.TgChannel {
				channelGood = true
				config = pp
				break
			}
		}
	}
	return channelGood, config
}
func (c CorpConfig) CheckCorpNameConfig(corpname string) (channelGood bool, config CorpporationConfig) {
	if corpname != "" { // есть ли корпа
		for _, pp := range *P {
			if corpname == pp.CorpName {
				channelGood = true
				config = pp
				break
			}
		}
	}
	return channelGood, config
}
func (c CorpConfig) ReadAllChannel() (chatDS []string, chatTG []int64, chatWA []string) {
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
