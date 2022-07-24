package dal

import (
	"model"
	"os"

	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func connectToDatabase() *mongo.Client {
	dbUri := os.Getenv("DB_URI")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dbUri))
	if err != nil {
		panic(err)
	}

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	return client
}

func GetAllTalents() []model.Talent {
	client := connectToDatabase()
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	talentsCollection := client.Database("crew-api").Collection("talents")
	var talents []model.Talent
	cursor, dbFindError := talentsCollection.Find(context.TODO(), bson.D{})

	if dbFindError != nil {
		panic(dbFindError)
	}

	for cursor.Next(context.TODO()) {
		var talent model.Talent
		err := cursor.Decode(&talent)
		if err != nil {
			panic(err)
		}
		talents = append(talents, talent)
	}

	return talents
}

func AddTalent(talent model.Talent) error {
	client := connectToDatabase()
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	talentsCollection := client.Database("crew-api").Collection("talents")
	_, dbAddError := talentsCollection.InsertOne(context.TODO(), talent)
	return dbAddError
}
