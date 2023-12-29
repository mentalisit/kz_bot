package mongo

//func (d *DB) GetCorporation(Corp string) models.CorporationHadesClient {
//	collection := d.s.Database("HadesClient").Collection("Corporation")
//	var mm models.CorporationHadesClient
//	err := collection.FindOne(context.Background(), bson.M{"corp": Corp}).Decode(&mm)
//	if err != nil {
//		d.log.Println(err)
//		return models.CorporationHadesClient{}
//	}
//	//fmt.Printf("%+v", mm)
//	return mm
//}
//func (d *DB) InsertCorporation(c models.CorporationHadesClient) {
//	collection := d.s.Database("HadesClient").Collection("Corporation")
//
//	ins, err := collection.InsertOne(context.Background(), c)
//	if err != nil {
//		d.log.Println(err)
//	}
//	fmt.Printf("%+v %+v\n", c, ins.InsertedID)
//}

//
//func (d *DB) GetCorporationDS(ChatId string) models.CorporationHadesClient {
//	collection := d.s.Database("HadesClient").Collection("Corporation")
//	var mm models.CorporationHadesClient
//	err := collection.FindOne(context.Background(), bson.M{"dschat": ChatId}).Decode(&mm)
//	if err != nil {
//		d.log.Println(err)
//		return models.CorporationHadesClient{}
//	}
//	//fmt.Printf("%+v", mm)
//	return mm
//}
//func (d *DB) GetCorporationDsWs1(ChatId string) models.CorporationHadesClient {
//	collection := d.s.Database("HadesClient").Collection("Corporation")
//	var mm models.CorporationHadesClient
//	err := collection.FindOne(context.Background(), bson.M{"dschatws1": ChatId}).Decode(&mm)
//	if err != nil {
//		d.log.Println(err)
//		return models.CorporationHadesClient{}
//	}
//	//fmt.Printf("%+v", mm)
//	return mm
//}
//
//func (d *DB) GetCorporationTG(ChatId int64) models.CorporationHadesClient {
//	collection := d.s.Database("HadesClient").Collection("Corporation")
//	var mm models.CorporationHadesClient
//	err := collection.FindOne(context.Background(), bson.M{"tgchat": ChatId}).Decode(&mm)
//	if err != nil {
//		d.log.Println(err)
//		return models.CorporationHadesClient{}
//	}
//	//fmt.Printf("%+v", mm)
//	return mm
//}
//func (d *DB) GetCorporationTgWs1(ChatId int64) models.CorporationHadesClient {
//	collection := d.s.Database("HadesClient").Collection("Corporation")
//	var mm models.CorporationHadesClient
//	err := collection.FindOne(context.Background(), bson.M{"tgchatws1": ChatId}).Decode(&mm)
//	if err != nil {
//		d.log.Println(err)
//		return models.CorporationHadesClient{}
//	}
//	//fmt.Printf("%+v", mm)
//	return mm
//}
