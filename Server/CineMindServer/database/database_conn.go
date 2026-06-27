package database

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func DBInstance() *mongo.Client {
	err := godotenv.Load(".env")

	if err != nil {
		log.Println("Warning: unable to find .env")
	}

	mongoURI := os.Getenv("MONGODB_URI")

	if mongoURI == "" {
		log.Fatal("MONGODB_URI not set")
	}

	clientOptions := options.Client().ApplyURI(mongoURI)

	client, err := mongo.Connect(clientOptions)

	if err != nil {
		log.Fatal("failed to create MongoDB client: ", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal("failed to connect to MongoDB: ", err)
	}

	return client
}

var Client *mongo.Client = DBInstance()

func OpenCollection(collectionName string) *mongo.Collection {
	databaseName := os.Getenv("DATABASE_NAME")

	if databaseName == "" {
		log.Fatal("DATABASE_NAME not set")
	}

	return Client.Database(databaseName).Collection(collectionName)
}
