package hlam

//
//func checkIdentity(code string) (string, error) {
//	apiURL := "https://bot.hs-compendium.com/compendium/applink/identities?ver=2&code=1"
//
//	// Подготовка запроса
//	req, err := http.NewRequest("GET", apiURL, nil)
//	if err != nil {
//		return "", err
//	}
//
//	// Установка параметров запроса
//	req.Header.Set("Cache-Control", "no-cache")
//	req.Header.Set("Authorization", code)
//
//	// Отправка запроса
//	client := &http.Client{}
//	resp, err := client.Do(req)
//	if err != nil {
//		return "", err
//	}
//	defer resp.Body.Close()
//
//	// Проверка успешного ответа
//	if resp.StatusCode < 200 || resp.StatusCode >= 500 {
//		return "", fmt.Errorf("Server Error")
//	}
//
//	// Чтение тела ответа
//	body, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		return "", err
//	}
//
//	// Декодирование JSON-строки в структуру Identity
//	var identity Identity
//	err = json.Unmarshal([]byte(body), &identity)
//	if err != nil {
//		fmt.Println("Error decoding JSON:", err)
//		return "", err
//	}
//
//	// Вывод данных
//	fmt.Println("User ID:", identity.User.ID)
//	fmt.Println("Username:", identity.User.Username)
//	fmt.Println("Discriminator:", identity.User.Discriminator)
//	fmt.Println("Avatar:", identity.User.Avatar)
//	fmt.Println("AvatarURL:", identity.User.AvatarURL)
//
//	fmt.Println("Token:", identity.Token)
//	return identity.Token, nil
//}
//
//func refreshConnection(token string) (*Identity, error) {
//	baseURL := "https://bot.hs-compendium.com/compendium"
//	url := fmt.Sprintf("%s/applink/refresh", baseURL)
//
//	req, err := http.NewRequest("GET", url, nil)
//	if err != nil {
//		return nil, fmt.Errorf("error creating request: %v", err)
//	}
//
//	// Установка заголовков
//	req.Header.Set("Cache-Control", "no-cache")
//	req.Header.Set("Content-Type", "application/json")
//	req.Header.Set("Authorization", token)
//
//	client := &http.Client{}
//	resp, err := client.Do(req)
//	if err != nil {
//		return nil, fmt.Errorf("error sending request: %v", err)
//	}
//	defer resp.Body.Close()
//
//	if resp.StatusCode < 200 || resp.StatusCode >= 500 {
//		return nil, fmt.Errorf("server error: %v", resp.Status)
//	}
//	// Чтение тела ответа
//	body, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		return nil, err
//	}
//	fmt.Printf("body %+v\n", string(body))
//	var identity Identity
//	decoder := json.NewDecoder(resp.Body)
//	if err := decoder.Decode(&identity); err != nil {
//		return nil, fmt.Errorf("error decoding JSON: %v", err)
//	}
//
//	// Обработка URL для изображений
//	//identity.Guilds.URL = fmt.Sprintf("https://cdn.discordapp.com/avatars/%s/%s.png", identity.User.ID, identity.User.Avatar)
//
//	return &identity, nil
//}
