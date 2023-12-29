package HadesClient

//
//func (h *Hades) ifComands(m models.MessageHades) (command bool) {
//	str, ok := strings.CutPrefix(m.Text, ". ")
//	if ok {
//		str = strings.ToLower(str)
//
//		if h.AddFriendToList(m) {
//			return true
//		}
//
//		arr := strings.Split(str, " ")
//		arrlen := len(arr)
//		if arrlen == 1 {
//
//		} else if arrlen == 2 {
//			if h.lastWs(arr, m) {
//				return true
//			}
//			if h.replayId(arr, m) {
//				return true
//			}
//			if h.historyWs(arr, m) {
//				return true
//			}
//			if h.letInId(arr, m) {
//				return true
//			}
//			if h.listAccess(arr, m) {
//				return true
//			}
//		}
//	}
//	return false
//}
//
//func (h *Hades) lastWs(arg []string, m models.MessageHades) bool {
//	if arg[0] == "повтор" && arg[1] == "бз" {
//		corporation := h.getConfig(m.Corporation)
//		message := models.MessageHadesClient{
//			Text:        "",
//			Sender:      m.Sender,
//			Avatar:      "",
//			ChannelType: 0,
//			Corporation: corporation.Corp,
//			Command:     "повтор бз",
//			Messager:    m.Messager,
//		}
//		fmt.Printf("lastWs %+v\n", mes)
//		h.toGame <- message
//		h.delSendMessageIfTip("отправка повтора последней бз", m, corporation, 10)
//
//	}
//	return false
//}
//func (h *Hades) replayId(arg []string, m models.MessageHades) bool {
//	if arg[0] == "повтор" {
//		match, _ := regexp.MatchString("^[0-9]+$", arg[1])
//		if match {
//			corporation := h.getConfig(m.Corporation)
//			message := models.MessageHadesClient{
//				Text:        arg[1],
//				Sender:      m.Sender,
//				Avatar:      "",
//				ChannelType: 0,
//				Corporation: corporation.Corp,
//				Command:     "повтор",
//				Messager:    m.Messager,
//			}
//			fmt.Printf("replayId %+v\n", mes)
//			h.toGame <- message
//			h.delSendMessageIfTip("отправка повтора "+arg[1], m, corporation, 10)
//			return true
//		}
//	}
//	return false
//}
//func (h *Hades) historyWs(arg []string, m models.MessageHades) bool {
//	if arg[0] == "история" && arg[1] == "бз" {
//		corporation := h.getConfig(m.Corporation)
//		message := models.MessageHadesClient{
//			Text:        "",
//			Sender:      m.Sender,
//			Avatar:      "",
//			ChannelType: 0,
//			Corporation: corporation.Corp,
//			Command:     "история бз",
//			Messager:    m.Messager,
//		}
//		fmt.Printf("historyWs %+v\n", mes)
//		h.toGame <- message
//		h.delSendMessageIfTip("готовлю список  бз", m, corporation, 10)
//		return true
//	}
//	return false
//}
//
//func (h *Hades) letInId(arg []string, m models.MessageHades) bool {
//	if arg[0] == "впустить" {
//		match, _ := regexp.MatchString("^[0-9]+$", arg[1])
//		if match {
//			corporation := h.getConfig(m.Corporation)
//			message := models.MessageHadesClient{
//				Text:        arg[1],
//				Sender:      m.Sender,
//				Avatar:      "",
//				ChannelType: 0,
//				Corporation: corporation.Corp,
//				Command:     "впустить",
//				Messager:    m.Messager,
//			}
//			fmt.Printf("letInId %+v\n", mes)
//			h.toGame <- message
//			h.delSendMessageIfTip("впустить отправленно  "+arg[1], m, corporation, 10)
//			return true
//		}
//	}
//	return false
//}
//
//func (h *Hades) listAccess(arg []string, m models.MessageHades) bool {
//	if arg[0] == "список" && arg[1] == "имён" {
//		corporation := h.getConfig(m.Corporation)
//		var text = "  Список имён доверенный\n"
//		var text2 = "  Список имён\n"
//		var n1 = 1
//		var n2 = 1
//		for _, s := range h.member {
//			if s.CorpName == "1" {
//				text = text + fmt.Sprintf("%d %s(%d) \n", n1, s.UserName, s.Rang)
//				n1++
//			}
//			if m.Corporation == s.CorpName {
//				text2 = text2 + fmt.Sprintf("%d %s(%d)\n", n2, s.UserName, s.Rang)
//				n2++
//			}
//		}
//		text = text + text2
//
//		h.delSendMessageIfTip(text, m, corporation, 120)
//		return true
//	}
//	return false
//}
//func (h *Hades) AddFriendToList(m models.MessageHades) bool {
//	re := regexp.MustCompile(`^\. Добавить ([0-2]) (.+)`)
//	matches := re.FindStringSubmatch(m.Text)
//	if len(matches) > 0 {
//		config := h.getConfig(m.Corporation)
//		if config.Corp != "" {
//			rang, _ := strconv.Atoi(matches[1])
//			name := matches[2]
//			h.storage.HadesClient.InsertMember(config.Corp, name, rang)
//			h.member[matches[2]] = models.AllianceMember{
//				CorpName: config.Corp,
//				UserName: name,
//				Rang:     rang,
//			}
//			t := fmt.Sprintf("Добавлен игрок %s в копрорацию %s", name, config.Corp)
//			h.delSendMessageIfTip(t, m, config, 10)
//			h.log.Println(t)
//			return true
//		}
//	}
//	return false
//}
//
//func (h *Hades) delSendMessageIfTip(text string, m models.MessageHades, corporation models.CorporationHadesClient, second int) {
//	if m.Messager == "ds" {
//		go h.cl.Ds.SendChannelDelSecond(corporation.DsChat, "```"+text+"```", second)
//		go h.cl.Ds.DeleteMesageSecond(corporation.DsChat, m.MessageId, 10)
//	}
//	if m.Messager == "tg" {
//		go h.cl.Tg.SendChannelDelSecond(corporation.TgChat, text, second)
//		go h.cl.Tg.DelMessageSecond(corporation.TgChat, m.MessageId, 10)
//	}
//}
