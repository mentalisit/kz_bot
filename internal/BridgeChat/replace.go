package BridgeChat

//func replaceTextMap(text string, m map[string]string) string {
//	mentionPattern := `@(\w+)|<@(\w+)>`
//	mentionRegex := regexp.MustCompile(mentionPattern)
//	text = mentionRegex.ReplaceAllStringFunc(text, func(match string) string {
//		if value, ok := m[match]; ok {
//			// Если значение найдено, заменяем упоминание на значение из map
//			return value
//		}
//
//		// Если значение не найдено, оставляем упоминание без изменений
//		return match
//	})
//	//fmt.Println(modifiedText)
//	return text
//}
