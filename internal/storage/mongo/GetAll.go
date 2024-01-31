package mongo

//func (d *DB) GetAllCorporationHades() []models.CorporationHadesClient {
//	collection := d.s.Database("HadesClient").Collection("Corporation")
//	var mm []models.CorporationHadesClient
//	find, err := collection.Find(context.Background(), bson.M{})
//	if err != nil {
//		d.log.ErrorErr(err)
//		return nil
//	}
//	err = find.All(context.Background(), &mm)
//	if err != nil {
//		d.log.ErrorErr(err)
//		return nil
//	}
//	//fmt.Printf("%+v", mm)
//	return mm
//}
//func (d *DB) GetAllGameMesId() []models.GameMessageId {
//	collection := d.s.Database("HadesClient").Collection("AllianceMesId")
//	var m []models.GameMessageId
//	find, err := collection.Find(context.Background(), bson.M{})
//	if err != nil {
//		d.log.ErrorErr(err)
//		return nil
//	}
//	err = find.All(context.Background(), &m)
//	if err != nil {
//		d.log.ErrorErr(err)
//		return nil
//	}
//
//	return m
//}
//func (d *DB) GetAllGameWs1MesId() []models.GameMessageIdWs1 {
//	collection := d.s.Database("HadesClient").Collection("Ws1MesId")
//	var m []models.GameMessageIdWs1
//	find, err := collection.Find(context.Background(), bson.M{})
//	if err != nil {
//		d.log.ErrorErr(err)
//		return nil
//	}
//	err = find.All(context.Background(), &m)
//	if err != nil {
//		d.log.ErrorErr(err)
//		return nil
//	}
//
//	return m
//}
//func (d *DB) GetAllMember() []models.AllianceMember {
//	collection := d.s.Database("HadesClient").Collection("AllianceMember")
//	var m []models.AllianceMember
//	find, err := collection.Find(context.Background(), bson.M{})
//	if err != nil {
//		d.log.ErrorErr(err)
//		return nil
//	}
//	err = find.All(context.Background(), &m)
//	if err != nil {
//		d.log.ErrorErr(err)
//		return nil
//	}
//
//	return m
//}
