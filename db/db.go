package db

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var Client *mongo.Client
var Database *mongo.Database

func ConnectDB() {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err := mongo.Connect(clientOptions)

	if err != nil {
		fmt.Println("FAILED TO CONNECT TO MONGO DB 📛 %w", err)
	}

	err = client.Ping(ctx, nil)

	if err != nil {
		fmt.Println("FAILED TO PING 📛 %w", err)
	}

	Client = client
	Database = client.Database("room-reserve")

	fmt.Println("SUCCESSFULLY CONNECTED TO MONGODB ✅")
}

func GetCollection(collectionName string) *mongo.Collection {
	return Database.Collection(collectionName)
}
