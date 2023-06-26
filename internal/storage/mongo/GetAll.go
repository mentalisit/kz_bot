package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"kz_bot/internal/models"
)

func (d *DB) GetAllCorporationHades() []models.CorporationHadesClient {
	collection := d.s.Database("HadesClient").Collection("Corporation")
	var mm []models.CorporationHadesClient
	find, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		d.log.Println("GetAllCorporationHades() " + err.Error())
		return nil
	}
	err = find.All(context.Background(), &mm)
	if err != nil {
		d.log.Println("GetAllCorporationHades() " + err.Error())
		return nil
	}
	//fmt.Printf("%+v", mm)
	return mm
}
func (d *DB) GetAllGameMesId() []models.GameMessageId {
	collection := d.s.Database("HadesClient").Collection("AllianceMesId")
	var m []models.GameMessageId
	find, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		d.log.Println("GetAllGameMesId() " + err.Error())
		return nil
	}
	err = find.All(context.Background(), &m)
	if err != nil {
		d.log.Println("GetAllGameMesId() " + err.Error())
		return nil
	}

	return m
}
func (d *DB) GetAllGameWs1MesId() []models.GameMessageIdWs1 {
	collection := d.s.Database("HadesClient").Collection("Ws1MesId")
	var m []models.GameMessageIdWs1
	find, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		d.log.Println("GetAllGameWs1MesId() " + err.Error())
		return nil
	}
	err = find.All(context.Background(), &m)
	if err != nil {
		d.log.Println("GetAllGameWs1MesId() " + err.Error())
		return nil
	}

	return m
}
func (d *DB) GetAllMember() []models.AllianceMember {
	collection := d.s.Database("HadesClient").Collection("AllianceMember")
	var m []models.AllianceMember
	find, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		d.log.Println("GetAllMember() " + err.Error())
		return nil
	}
	err = find.All(context.Background(), &m)
	if err != nil {
		d.log.Println("GetAllMember() " + err.Error())
		return nil
	}

	return m
}
