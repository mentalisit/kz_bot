package BridgeChat

import (
	"fmt"
	"kz_bot/internal/models"
	"strings"
)

func (b *Bridge) Command() {
	after, _ := strings.CutPrefix(b.in.Text, ".")
	arg := strings.Split(after, " ")
	lenarg := len(arg)
	if lenarg == 1 {
		if arg[0] == "help" {
			//help
			return
		}
	} else if lenarg == 2 {
		if arg[0] == "список" && arg[1] == "каналов" {
			if b.in.Config.HostRelay == "" {
				return
			}
			text := fmt.Sprintf("Список каналов хоста %s\n", b.in.Config.HostRelay)
			if len(b.in.Config.ChannelDs) > 0 {
				for _, d := range b.in.Config.ChannelDs {
					text = text + "[DS]" + d.AliasName + " (" + d.CorpChannelName + ")\n"
				}
			}
			if len(b.in.Config.ChannelTg) > 0 {
				for _, d := range b.in.Config.ChannelTg {
					text = text + "[TG]" + d.AliasName + " (" + d.CorpChannelName + ")\n"
				}
			}
			go b.ifTipDelSend(text)
			return
		}
	} else if lenarg == 3 {
		if arg[0] == "создать" && arg[1] == "реле" {
			good, _ := b.CacheNameBridge(arg[2])
			if !good {
				bridge := models.BridgeConfig{
					NameRelay:         arg[2],
					HostRelay:         b.GuildName(),
					Role:              []string{},
					ForbiddenPrefixes: []string{},
				}
				b.ifChannelTip(&bridge)
				b.AddNewBridgeConfig(bridge)
				text := fmt.Sprintf("%s создано, \nиспользуй команду в другом канале для подключения .подключить реле %s", arg[2], arg[2])
				b.ifTipDelSend(text)
			} else {
				b.ifTipDelSend(arg[2] + " уже существует")
			}
			return
		}
		if arg[0] == "подключить" && arg[1] == "реле" {
			good, host := b.CacheNameBridge(arg[2])
			channelGood := false
			if b.in.Tip == "ds" {
				channelGood, _ = b.CacheCheckChannelConfigDS(b.in.Ds.ChatId)
			} else if b.in.Tip == "tg" {
				channelGood, _ = b.CacheCheckChannelConfigTg(b.in.Tg.ChatId)
			}
			if good && !channelGood {
				bridge := models.BridgeConfig{
					NameRelay: arg[2],
					HostRelay: host.HostRelay,
				}
				b.ifChannelTip(&bridge)
				b.AddNewBridgeConfig(bridge)
				text := fmt.Sprintf("Реле %s: добавлен текущий канал\nСписок подключеных канлов к реле %s доступен по команде `.список каналов`", arg[2], arg[2])
				b.ifTipDelSend(text)
				//отправить сообщение о подкоючении канала
			} else {
				b.ifTipDelSend(arg[2] + " уже существует")
			}
			return
		}
	}
}
