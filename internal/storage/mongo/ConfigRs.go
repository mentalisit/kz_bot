package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"kz_bot/internal/models"
)

func (d *DB) ReadConfigRs() []models.CorporationConfig {
	collection := d.s.Database("RsBot").Collection("RsConfig")
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		d.log.Println("ReadConfigRs " + err.Error())
		return nil
	}
	var m []models.CorporationConfig
	err = cursor.All(context.Background(), &m)
	if err != nil {
		d.log.Println("ReadConfigRs2 " + err.Error())
		return nil
	}
	for _, config := range m {
		fmt.Printf("ReadConfigRs %+v", config)
	}
	return m
}
func (d *DB) InsertConfigRs(c models.CorporationConfig) {
	collection := d.s.Database("RsBot").Collection("RsConfig")
	ins, err := collection.InsertOne(context.Background(), c)
	if err != nil {
		d.log.Println("InsertConfigRs " + err.Error())
	}
	fmt.Println(ins.InsertedID)
}
func (d *DB) DeleteConfigRs(c models.CorporationConfig) {
	collection := d.s.Database("RsBot").Collection("RsConfig")
	ins, err := collection.DeleteOne(context.Background(), c)
	if err != nil {
		d.log.Println("DeleteConfigRs " + err.Error())
	}
	fmt.Println(ins.DeletedCount)
}
