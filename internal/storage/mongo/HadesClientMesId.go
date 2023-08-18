package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
)

func (d *DB) InsertCorpMesId(CorpName string, mId int64) {
	collection := d.s.Database("HadesClient").Collection("AllianceMesId")
	type m struct {
		MessageId int64  `bson:"MessageId"`
		CorpName  string `bson:"CorpName"`
	}
	mes := m{
		MessageId: mId,
		CorpName:  CorpName,
	}
	ins, err := collection.InsertOne(context.Background(), mes)
	if err != nil {
		d.log.Println(err)
	}
	fmt.Printf("%+v %+v\n", mes, ins.InsertedID)
}
func (d *DB) GetCorpMesId(CorpName string) int64 {
	collection := d.s.Database("HadesClient").Collection("AllianceMesId")

	var mm struct {
		MessageId int64  `bson:"MessageId"`
		CorpName  string `bson:"CorpName"`
	}

	err := collection.FindOne(context.Background(), bson.M{"CorpName": CorpName}).Decode(&mm)
	if err != nil {
		d.log.Println(err)
		return 0
	}

	return mm.MessageId
}
func (d *DB) UpdateCorpMesId(CorpName string, mID int64) error {
	collection := d.s.Database("HadesClient").Collection("AllianceMesId")
	filter := bson.M{"CorpName": CorpName}
	update := bson.M{"CorpName": CorpName, "MessageId": mID}
	one, err := collection.ReplaceOne(context.Background(), filter, update)
	if err != nil {
		d.log.Println(err)
		return err
	}
	//fmt.Printf("%+v\n", one)
	if one.MatchedCount == 0 {
		d.InsertCorpMesId(CorpName, mID)
	}
	return nil
}

func (d *DB) InsertWs1MesId(CorpName string, mId int64, StarId int64) {
	collection := d.s.Database("HadesClient").Collection("Ws1MesId")
	type m struct {
		MessageId int64  `bson:"MessageId"`
		CorpName  string `bson:"CorpName"`
		StarId    int64  `bson:"StarId"`
	}
	mes := m{
		MessageId: mId,
		CorpName:  CorpName,
		StarId:    StarId,
	}
	ins, err := collection.InsertOne(context.Background(), mes)
	if err != nil {
		d.log.Println(err)
	}
	fmt.Printf("%+v %+v\n", mes, ins.InsertedID)
}
func (d *DB) GetWs1MesId(CorpName string, StarId int64) int64 {
	collection := d.s.Database("HadesClient").Collection("Ws1MesId")

	var mm struct {
		MessageId int64  `bson:"MessageId"`
		CorpName  string `bson:"CorpName"`
		StarId    int64  `bson:"StarId"`
	}

	err := collection.FindOne(context.Background(), bson.M{"CorpName": CorpName, "StarId": StarId}).Decode(&mm)
	if err != nil {
		//d.log.Println(err)
		return 0
	}

	return mm.MessageId
}
func (d *DB) UpdateWs1MesId(CorpName string, mID, StarId int64) {
	collection := d.s.Database("HadesClient").Collection("Ws1MesId")
	filter := bson.M{"CorpName": CorpName, "StarId": StarId}
	update := bson.M{"CorpName": CorpName, "MessageId": mID, "StarId": StarId}
	one, err := collection.ReplaceOne(context.Background(), filter, update)
	if err != nil {
		fmt.Println(err)
		return
	}
	//fmt.Printf("%+v\n", one)
	if one.MatchedCount == 0 {
		d.InsertWs1MesId(CorpName, mID, StarId)
	}
}
