package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"kz_bot/internal/models"
)

func (d *DB) DBReadBridgeConfig() []models.BridgeConfig {
	var data []models.BridgeConfig
	collection := d.s.Database("BridgeChat").Collection("Bridge")
	find, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		d.log.ErrorErr(err)
		return nil
	}
	err = find.All(context.Background(), &data)
	if err != nil {
		d.log.ErrorErr(err)
		return nil
	}
	return data
}
func (d *DB) UpdateBridgeChat(br models.BridgeConfig) {
	collection := d.s.Database("BridgeChat").Collection("Bridge")
	filter := bson.M{"namerelay": br.NameRelay}
	collection.FindOneAndDelete(context.Background(), filter)
	d.InsertBridgeChat(br)
}
func (d *DB) InsertBridgeChat(br models.BridgeConfig) {
	collection := d.s.Database("BridgeChat").Collection("Bridge")
	bsonData, _ := bson.Marshal(br)
	_, err := collection.InsertOne(context.Background(), bsonData)
	if err != nil {
		d.log.ErrorErr(err)
		//return
	}
}
