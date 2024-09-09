package database

import (
    "context"
    "log"
    "time"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)



var DB *mongo.Database

func InitDB(){
	clientOptions := options.Client().ApplyURI("mongodb+srv://chitluridevicharan:charan@cluster0.qgrpql7.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0")
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    err = client.Connect(ctx)
    if err != nil {
        log.Fatal(err)
    }

    DB = client.Database("lms")

}