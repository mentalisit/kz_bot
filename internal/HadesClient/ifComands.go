package HadesClient

import (
	"fmt"
	"kz_bot/internal/models"
	"regexp"
	"strconv"
	"strings"
)

func (h *Hades) ifComands(m models.MessageHades) (command bool) {
	str, ok := strings.CutPrefix(m.Text, ". ")
	if ok {
		str = strings.ToLower(str)

		if h.AddFriendToList(m) {
			return true
		}

		arr := strings.Split(str, " ")
		arrlen := len(arr)
		if arrlen == 1 {

		} else if arrlen == 2 {
			if h.lastWs(arr, m) {
				return true
			}
			if h.replayId(arr, m) {
				return true
			}
			if h.historyWs(arr, m) {
				return true
			}
			if h.letInId(arr, m) {
				return true
			}
		}
	}
	return false
}

func (h *Hades) lastWs(arg []string, m models.MessageHades) bool {
	if arg[0] == "повтор" && arg[1] == "бз" {
		corporation := h.getConfig(m.Corporation)
		message := models.MessageHadesClient{
			Text:        "",
			Sender:      m.Sender,
			Avatar:      "",
			ChannelType: 0,
			Corporation: corporation.Corp,
			Command:     "повтор бз",
			Messager:    m.Messager,
		}
		fmt.Printf("lastWs %+v\n", mes)
		h.toGame <- message
		h.delSendMessageIfTip("отправка повтора последней бз", m, corporation)

	}
	return false
}
func (h *Hades) replayId(arg []string, m models.MessageHades) bool {
	if arg[0] == "повтор" {
		match, _ := regexp.MatchString("^[0-9]+$", arg[1])
		if match {
			corporation := h.getConfig(m.Corporation)
			message := models.MessageHadesClient{
				Text:        arg[1],
				Sender:      m.Sender,
				Avatar:      "",
				ChannelType: 0,
				Corporation: corporation.Corp,
				Command:     "повтор",
				Messager:    m.Messager,
			}
			fmt.Printf("replayId %+v\n", mes)
			h.toGame <- message
			h.delSendMessageIfTip("отправка повтора "+arg[1], m, corporation)
			return true
		}
	}
	return false
}
func (h *Hades) historyWs(arg []string, m models.MessageHades) bool {
	if arg[0] == "история" && arg[1] == "бз" {
		corporation := h.getConfig(m.Corporation)
		message := models.MessageHadesClient{
			Text:        "",
			Sender:      m.Sender,
			Avatar:      "",
			ChannelType: 0,
			Corporation: corporation.Corp,
			Command:     "история бз",
			Messager:    m.Messager,
		}
		fmt.Printf("historyWs %+v\n", mes)
		h.toGame <- message
		h.delSendMessageIfTip("готовлю список  бз", m, corporation)
		return true
	}
	return false
}

func (h *Hades) letInId(arg []string, m models.MessageHades) bool {
	if arg[0] == "впустить" {
		match, _ := regexp.MatchString("^[0-9]+$", arg[1])
		if match {
			corporation := h.getConfig(m.Corporation)
			message := models.MessageHadesClient{
				Text:        arg[1],
				Sender:      m.Sender,
				Avatar:      "",
				ChannelType: 0,
				Corporation: corporation.Corp,
				Command:     "впустить",
				Messager:    m.Messager,
			}
			fmt.Printf("letInId %+v\n", mes)
			h.toGame <- message
			h.delSendMessageIfTip("впустить отправленно  "+arg[1], m, corporation)
			return true
		}
	}
	return false
}

func (h *Hades) AddFriendToList(m models.MessageHades) bool {
	re := regexp.MustCompile(`^\. Добавить ([0-2]) (.+)`)
	matches := re.FindStringSubmatch(m.Text)
	if len(matches) > 0 {
		config := h.getConfig(m.Corporation)
		if config.Corp != "" {
			rang, _ := strconv.Atoi(matches[1])
			name := matches[2]
			h.storage.HadesClient.InsertMember(config.Corp, name, rang)
			h.member[matches[2]] = models.AllianceMember{
				CorpName: config.Corp,
				UserName: name,
				Rang:     rang,
			}
			t := fmt.Sprintf("Добавлен игрок %s в копрорацию %s", name, config.Corp)
			h.delSendMessageIfTip(t, m, config)
			h.log.Println(t)
			return true
		}
	}
	return false
}

func (h *Hades) delSendMessageIfTip(text string, m models.MessageHades, corporation models.CorporationHadesClient) {
	if m.Messager == "ds" {
		go h.cl.Ds.SendChannelDelSecond(corporation.DsChat, "```"+text+"```", 10)
		go h.cl.Ds.DeleteMesageSecond(corporation.DsChat, m.Ds.MessageId, 10)
	}
	if m.Messager == "tg" {
		go h.cl.Tg.SendChannelDelSecond(corporation.TgChat, text, 10)
		go h.cl.Tg.DelMessageSecond(corporation.TgChat, m.Tg.MessageId, 10)
	}
}
