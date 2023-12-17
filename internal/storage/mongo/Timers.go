package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"kz_bot/internal/models"
)

func (d *DB) TimerInsert(c models.Timer) {
	collection := d.s.Database("RsBot").Collection("Timers")
	_, err := collection.InsertOne(context.Background(), c)
	if err != nil {
		d.log.Println("TimerInsert " + err.Error())
	}
}

func (d *DB) TimerDeleteMessage() []models.Timer {
	collection := d.s.Database("RsBot").Collection("Timers")
	filter := bson.D{}
	update := bson.D{{"$inc", bson.D{{"timed", -60}}}}
	t, err := collection.UpdateMany(context.Background(), filter, update)
	if err != nil {
		//d.log.Println("TimerDeleteMessage24 " + err.Error())
		return nil
	}
	if t.MatchedCount > 0 {
		filter = bson.D{{"timed", bson.D{{"$lt", 60}}}}
		find, err := collection.Find(context.Background(), filter)
		if err != nil {
			fmt.Println("TimerDeleteMessage32 " + err.Error())
			return nil
		}
		var modAll []models.Timer
		err = find.All(context.Background(), &modAll)
		if err != nil {
			fmt.Println("TimerDeleteMessage38 " + err.Error())
			return nil
		}
		_, err = collection.DeleteMany(context.Background(), filter)
		if err != nil {
			fmt.Println("TimerDeleteMessage38 " + err.Error())
			return nil
		}
		return modAll
	}
	return nil
}
