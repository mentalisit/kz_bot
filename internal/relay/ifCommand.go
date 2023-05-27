package relay

import (
	"fmt"
	"kz_bot/internal/models"
	"strings"
)

func (r *Relay) ifCommand() bool {
	fmt.Println("ifCommand() " + r.in.Text)
	after, _ := strings.CutPrefix(r.in.Text, ".")
	arg := strings.Split(after, " ")
	lenarg := len(arg)
	fmt.Printf("ifCommand() %d", lenarg)
	if lenarg == 1 {
		if arg[0] == "help" {
			//help
			return true
		}
	} else if lenarg == 2 {
		if arg[0] == "список" && arg[1] == "каналов" {
			channelGood := false
			configRelay := models.RelayConfig{}
			if r.in.Tip == "ds" {
				channelGood, configRelay = r.storage.CorpsConfig.RelayCache.CheckChannelConfigDS(r.in.Ds.ChatId)
			} else if r.in.Tip == "tg" {
				channelGood, _ = r.storage.CorpsConfig.RelayCache.CheckChannelConfigTG(r.in.Tg.ChatId)
			}
			if channelGood {
				list := r.storage.CorpsConfig.RelayCache.ListNameRelay(configRelay.RelayName)
				var text string
				for _, config := range list {
					if config.DsChannel != "" {
						text = text + "[DS]" + config.GuildName + "\n"
					}
					if config.TgChannel != 0 {
						text = text + "[TG]" + config.GuildName + "\n"
					}
				}
				go r.ifTipDelSend(text)
				return true
			}
		}

	} else if lenarg == 3 {
		if arg[0] == "создать" && arg[1] == "реле" {
			good, _ := r.storage.CorpsConfig.RelayCache.CheckChannelConfigRelayName(arg[2])
			if !good {
				relay := models.RelayConfig{
					RelayName:  arg[2],
					RelayAlias: arg[2],
					GuildName:  r.GuildName(),
					Country:    "ru",
					Prefix:     ".",
				}
				r.ifChannelTip(&relay)
				r.storage.CorpsConfig.RelayCache.AddCorp(relay)
				err := r.storage.CorpsConfig.RelayDB.AddRelay(relay)
				if err != nil {
					r.log.Println(err)
					return false
				}
				text := fmt.Sprintf("%s создано, \nиспользуй команду в другом канале для подключения `.подключить реле %s`", arg[2], arg[2])
				r.ifTipDelSend(text)
			} else {
				r.ifTipDelSend(arg[2] + " уже существует")
			}

		}
		if arg[0] == "подключить" && arg[1] == "реле" {
			good, _ := r.storage.CorpsConfig.RelayCache.CheckChannelConfigRelayName(arg[2])
			channelGood := false
			if r.in.Tip == "ds" {
				channelGood, _ = r.storage.CorpsConfig.RelayCache.CheckChannelConfigDS(r.in.Ds.ChatId)
			} else if r.in.Tip == "tg" {
				channelGood, _ = r.storage.CorpsConfig.RelayCache.CheckChannelConfigTG(r.in.Tg.ChatId)
			}

			if good && !channelGood {
				relay := models.RelayConfig{
					RelayName:  arg[2],
					RelayAlias: arg[2],
					GuildName:  r.GuildName(),
					Country:    "ru",
					Prefix:     ".",
				}
				r.ifChannelTip(&relay)
				r.storage.CorpsConfig.RelayCache.AddCorp(relay)
				err := r.storage.CorpsConfig.RelayDB.AddRelay(relay)
				if err != nil {
					r.log.Println(err)
					return false
				}
				text := fmt.Sprintf("Реле %s: добавлен текущий канал\nСписок подключеных канлов к реле %s доступен по команде `.список каналов`", arg[2], arg[2])
				r.ifTipDelSend(text)
			} else {
				r.ifTipDelSend(arg[2] + " уже существует")
			}

		}
	}
	return false
}
func (r *Relay) GuildName() string {
	if r.in.Tip == "ds" {
		return r.client.Ds.GuildChatName(r.in.Ds.ChatId, r.in.Ds.GuildId)
	}
	if r.in.Tip == "tg" {
		return r.in.Tg.GroupName
	}
	return ""
}
func (r *Relay) ifChannelTip(relay *models.RelayConfig) {
	if r.in.Tip == "ds" {
		relay.DsChannel = r.in.Ds.ChatId
		relay.GuildId = r.in.Ds.GuildId
	}
	if r.in.Tip == "tg" {
		relay.TgChannel = r.in.Tg.ChatId
	}
}
