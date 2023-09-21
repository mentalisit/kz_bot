package HadesClient

import (
	"kz_bot/internal/models"
	"regexp"
	"time"
)

type save struct {
	kz       string
	mesid    string
	mesidtg  int
	corpname string
	time     int64
}

var SaveId []save

func (h *Hades) ifEditMessage(msg models.MessageHadesClient, corp models.CorporationHadesClient) bool {
	if msg.Command == "text" {
		sender := "(🎮)" + msg.Sender
		flag := false
		var S []save
		re := regexp.MustCompile(`^Инициировал поиск КРАСНОЙ ЗВЕЗДЫ ур\.(1[0-2]|[5-9])$`)
		ok := re.MatchString(msg.Text)
		if ok {
			match := re.FindStringSubmatch(msg.Text)
			if len(match) > 1 {
				var s save
				s.kz = match[1]
				if corp.DsChat != "" {
					msgd := ifRsSearch(msg)
					s.mesid = h.cl.Ds.SendWebhookForHades(msgd.Text, sender, corp.DsChat, corp.GuildId, msgd.Avatar)
				}
				if corp.TgChat != "" {
					text := sender + ": " + msg.Text
					s.mesidtg = h.cl.Tg.SendChannel(corp.TgChat, text)
					if corp.Corp == "Ещё один Миф" {
						go h.cl.Tg.SendChannel("-1001527134772/13261", text) /// HardCode
					}
				}
				s.corpname = msg.Corporation
				s.time = time.Now().Unix()
				SaveId = append(SaveId, s)
				return true
			}
		}
		reg := regexp.MustCompile(`^КЗ-([5-9]|1[0-2])(\s[2|3]\sиз\s4)$`)
		okg := reg.MatchString(msg.Text)
		if okg {
			match := reg.FindStringSubmatch(msg.Text)
			if len(match) > 1 {
				if len(SaveId) != 0 {
					for _, s := range SaveId {
						if time.Now().Unix() < s.time+60 {
							S = append(S, s)
							if s.kz == match[1] {
								if corp.DsChat != "" {
									h.cl.Ds.EditWebhookForHades(msg.Text, sender, corp.DsChat, corp.GuildId, msg.Avatar, s.mesid)
								}
								if corp.TgChat != "" {
									text := sender + ": " + msg.Text
									h.cl.Tg.EditText(corp.TgChat, s.mesidtg, text)
								}
								flag = true
							}
						}
					}
					SaveId = S
					return flag
				}
			}
		}
		ref := regexp.MustCompile(`^КЗ-([5-9]|1[0-2])(\sФулка)`)
		okf := ref.MatchString(msg.Text)
		if okf {
			match := ref.FindStringSubmatch(msg.Text)
			if len(match) > 1 {
				if len(SaveId) != 0 {
					for _, s := range SaveId {
						if time.Now().Unix() < s.time+60 {
							S = append(S, s)
							if s.kz == match[1] {
								if corp.DsChat != "" {
									h.cl.Ds.EditWebhookForHades(msg.Text, sender, corp.DsChat, corp.GuildId, msg.Avatar, s.mesid)
								}
								if corp.TgChat != "" {
									text := sender + ": " + msg.Text
									h.cl.Tg.EditText(corp.TgChat, s.mesidtg, text)
								}
								flag = true
							}
						}
					}
					SaveId = S
					return flag
				}
			}
		}
	}
	return false
}

func ifRsSearch(msg models.MessageHadesClient) models.MessageHadesClient {
	if msg.Command == "text" && msg.Corporation == "UKR Spase" {
		re := regexp.MustCompile(`КРАСНОЙ ЗВЕЗДЫ ур\.([5-9]|10)`)
		msg.Text = re.ReplaceAllStringFunc(msg.Text, textToRoleRsUkrSpace)
	}
	return msg
}
func (h *Hades) numToRole() {
	if h.in.Command == "text" && h.in.Corporation == "UKR Spase" {
		reRS := regexp.MustCompile(`^([5-9]|[10])(\?+)$`)
		arg := reRS.FindAllStringSubmatch(h.in.Text, -1)
		if len(arg) > 0 {
			h.in.Text = numtoroleUKRSpace(arg[0][1])
		}
	}
	return
}

// замена текста КРАСНОЙ ЗВЕЗДЫ ур на пинг дискорда
func textToRoleRsUkrSpace(s string) string {
	switch s {
	case "КРАСНОЙ ЗВЕЗДЫ ур.5":
		return "<@&763476853364228106>"
	case "КРАСНОЙ ЗВЕЗДЫ ур.6":
		return "<@&763476906850779170>"
	case "КРАСНОЙ ЗВЕЗДЫ ур.7":
		return "<@&763476952455446568>"
	case "КРАСНОЙ ЗВЕЗДЫ ур.8":
		return "<@&763477036831998002>"
	case "КРАСНОЙ ЗВЕЗДЫ ур.9":
		return "<@&788847032215142420>"
	case "КРАСНОЙ ЗВЕЗДЫ ур.10":
		return "<@&788846996836450385>"
	default:
		return s
	}
}

// замена цифры на пинг дискорда
func numtoroleUKRSpace(s string) string {
	switch s {
	case "5":
		return "<@&763476853364228106>"
	case "6":
		return "<@&763476906850779170>"
	case "7":
		return "<@&763476952455446568>"
	case "8":
		return "<@&763477036831998002>"
	case "9":
		return "<@&788847032215142420>"
	case "10":
		return "<@&788846996836450385>"
	default:
		return s
	}
}
