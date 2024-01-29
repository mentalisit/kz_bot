package hlam

//func main2() {
//
//	code := "PzDr-ub4U-iRG5"
//
//	client := bot_api.NewCompendiumApiClient()
//
//	// Подтверждение идентификации с использованием предоставленного кода
//	preConnectIdent, err := client.CheckIdentity(code)
//	if err != nil {
//		log.Println(err)
//	}
//
//	fmt.Printf("Successfully submitted code and retrieved identity:User: %s Guild: %s\n", preConnectIdent.User.Username, preConnectIdent.Guild.Name)
//
//	// Подтверждение идентификации и подключение
//	ident, err := client.Connect(preConnectIdent)
//	if err != nil {
//		log.Println(err)
//	}
//
//	fmt.Printf("Successfully connected:User: %s Guild: %s\n", ident.User.Username, ident.Guild.Name)
//
//	// Периодическое обновление подключения
//	_, err = client.RefreshConnection(ident.Token)
//	if err != nil {
//		log.Println(err)
//	}
//	fmt.Println("Successfully refreshed connection")
//
//	// Сохранение информации о подключении
//	identPath := filepath.Join(".", StorageKey)
//	identJSON, err := json.Marshal(ident)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	err = os.WriteFile(identPath, identJSON, 0644)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// Получение данных синхронизации
//	data, err := client.Sync(ident.Token, "get", nil)
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Printf("%+v\n", data)
//	dataJSON, err := json.MarshalIndent(data, "", "   ")
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	fmt.Println(string(dataJSON))
//	fmt.Println("DONE")
//}
