package app

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoDBInit() *mongo.Database {
	ctx, cancel := NewMongoContext()
	defer cancel()

	URI := os.Getenv("MONGO_URI")
	option := options.Client().ApplyURI(URI)

	client, err := mongo.NewClient(option)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	DBName := os.Getenv("MONGO_DATABASE")
	database := client.Database(DBName)
	return database
}

func NewMongoContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 10*time.Second)
}

// func MongoDBInit() *mongo.Client {
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}
// 	URI := os.Getenv("MONGO_URI")

// 	client, err := mongo.NewClient(options.Client().ApplyURI(URI))
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
// 	err = client.Connect(ctx)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// ping the database
// 	err = client.Ping(ctx, nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println("Connected to mongoDB")
// 	return client
// }

// // Client instance
// var DB *mongo.Client = MongoDBInit()

// // Getting database collection
// func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
// 	collection := client.Database("appMarket").Collection(collectionName)
// 	return collection
// }
