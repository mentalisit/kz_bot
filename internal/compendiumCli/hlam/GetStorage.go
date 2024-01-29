package hlam

//func GetStorage() (bool, Identity) {
//	// Загрузка существующей идентификации
//	identPath := filepath.Join(".", StorageKey)
//	identBytes, err := ioutil.ReadFile(identPath)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	var ident Identity
//	if err := json.Unmarshal(identBytes, &ident); err != nil {
//		log.Fatal(err)
//	}
//	if ident.Token == "" {
//		fmt.Println("err")
//	}
//	fmt.Printf("%+v\n", ident)
//	client := bot_api.NewCompendiumApiClient("https://bot.hs-compendium.com/compendium")
//
//	rv, err := client.Sync(ident.Token, "get", nil)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	rvJSON, err := json.MarshalIndent(rv, "", "   ")
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	fmt.Println(string(rvJSON))
//	fmt.Println("DONE")
//	return true, ident
//}
