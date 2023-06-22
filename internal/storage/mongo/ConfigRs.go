package mongo

import (
	"context"
	"fmt"
	"kz_bot/internal/models"
)

func (d *DB) InsertConfigRs(c models.CorporationConfig) {
	collection := d.s.Database("RsBot").Collection("RsConfig")
	ins, err := collection.InsertOne(context.Background(), c)
	if err != nil {
		d.log.Println("InsertConfigRs " + err.Error())
	}
	fmt.Println(ins.InsertedID)
}
