package db

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoClient() *mongo.Client {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Get mongodb uri from environment
	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		log.Fatal("'MONGO_URI' environment not set")
	}

	// Context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to mongodb
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	defer Disconnect(client)

	return client
}

func Disconnect(client *mongo.Client) {
	errMsg := recover()
	if errMsg != nil {
		log.Fatal(errMsg)
	}
}
