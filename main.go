package main

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

func main () {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Panic(err)
	}

	client.StartSession()

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	if err = client.Connect(ctx); err !=nil {
		log.Panic(err)
	}

	ctx, _ = context.WithTimeout(context.Background(), 2*time.Second)
	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		log.Panic(err)
	}

	collection := client.Database("reezorcar").Collection("cars")
	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
	res, err := collection.InsertOne(ctx, bson.M{"name": "pi", "value": 3.14159})
	if err != nil {
		log.Panic(err)
	}
	id := res.InsertedID
	log.Printf("%s", id)


}


