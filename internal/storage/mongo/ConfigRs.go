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
		d.log.Error(err.Error())
		return nil
	}
	var m []models.CorporationConfig
	err = cursor.All(context.Background(), &m)
	if err != nil {
		d.log.Error(err.Error())
		return nil
	}
	return m
}
func (d *DB) InsertConfigRs(c models.CorporationConfig) {
	//d.s.Database("RsBot").CreateCollection(context.Background(), "RsConfig")
	collection := d.s.Database("RsBot").Collection("RsConfig")
	ins, err := collection.InsertOne(context.Background(), c)
	if err != nil {
		d.log.Error(err.Error())
	}
	fmt.Println(ins.InsertedID)
}
func (d *DB) DeleteConfigRs(c models.CorporationConfig) {
	collection := d.s.Database("RsBot").Collection("RsConfig")
	ins, err := collection.DeleteOne(context.Background(), c)
	if err != nil {
		d.log.Error(err.Error())
	}
	fmt.Println(ins.DeletedCount)
}
func (d *DB) AutoHelpUpdateMesid(c models.CorporationConfig) {
	collection := d.s.Database("RsBot").Collection("RsConfig")
	filter := bson.M{"dschannel": c.DsChannel}
	//update := bson.M{"dschannel": dschannel, "mesiddshelp": newMesidHelp}
	_, err := collection.ReplaceOne(context.Background(), filter, c)
	if err != nil {
		d.log.Error(err.Error())
	}
}
func (d *DB) AutoHelp() []models.CorporationConfig {
	corp := d.ReadConfigRs()
	var c []models.CorporationConfig
	for _, config := range corp {
		if config.DsChannel != "" {
			c = append(c, config)
		}
	}
	return c
}
