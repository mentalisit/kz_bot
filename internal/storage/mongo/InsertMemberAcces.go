package mongo

import (
	"context"
	"fmt"
	"kz_bot/internal/models"
)

func (d *DB) InsertMember(CorpName, UserName string, Rang int) {
	collection := d.s.Database("HadesClient").Collection("AllianceMember")

	m := &models.AllianceMember{
		CorpName: CorpName,
		UserName: UserName,
		Rang:     Rang,
	}
	ins, err := collection.InsertOne(context.Background(), m)
	if err != nil {
		d.log.Error(err.Error())
	}
	fmt.Println(ins.InsertedID)
}
