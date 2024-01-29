package hlam

//// GetCorpData retrieves data about the corporation using the Compendium API.
//func (c *CompendiumApiClient) GetCorpData(token string, roleId string) (CorpData, error) {
//	url := fmt.Sprintf("%s/cmd/corpdata?roleId=%s", c.URL, roleId)
//	req, err := http.NewRequest("GET", url, nil)
//	if err != nil {
//		fmt.Printf("53err  %+v\n", err)
//		return CorpData{}, err
//	}
//	fmt.Println("ok56")
//	req.Header.Set("Content-Type", "application/json")
//	req.Header.Set("Authorization", token)
//
//	client := &http.Client{}
//	resp, err := client.Do(req)
//	if err != nil {
//		fmt.Printf("63err  %+v\n", err)
//		return CorpData{}, err
//	}
//	fmt.Println("ok66")
//
//	fmt.Println(resp.Status)
//
//	defer resp.Body.Close()
//
//	if resp.StatusCode < 200 || resp.StatusCode >= 500 {
//		return CorpData{}, fmt.Errorf("Server Error")
//	}
//	fmt.Println("ok75")
//	body, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		fmt.Printf("75err  %+v\n", err)
//		return CorpData{}, err
//	}
//	fmt.Println("ok81")
//	fmt.Println(string(body))
//	var corpData CorpData
//	err = json.Unmarshal(body, &corpData)
//	if err != nil {
//		fmt.Printf("82err  %+v\n", err)
//		return CorpData{}, err
//	}
//
//	return corpData, nil
//}
