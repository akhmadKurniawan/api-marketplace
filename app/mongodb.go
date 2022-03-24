package app

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoDBInit() *mongo.Database {
	ctx, cancel := NewMongoContext()
	defer cancel()

	mongoPoolMin, err := strconv.Atoi(os.Getenv("MONGO_POOL_MIN"))
	if err != nil {
		log.Fatal(err)
	}

	mongoPoolMax, err := strconv.Atoi(os.Getenv("MONGO_POOL_MAX"))
	if err != nil {
		log.Fatal(err)
	}

	mongoMaxIdleTime, err := strconv.Atoi(os.Getenv("MONGO_MAX_IDLE_TIME_SECOND"))
	if err != nil {
		log.Fatal(err)
	}

	URI := os.Getenv("MONGO_URI")
	option := options.Client().
		ApplyURI(URI).
		SetMinPoolSize(uint64(mongoPoolMin)).
		SetMaxPoolSize(uint64(mongoPoolMax)).
		SetMaxConnIdleTime(time.Duration(mongoMaxIdleTime) * time.Second)

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
	fmt.Println("Connected to mongoDB", database)
	return database
}

func NewMongoContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 10*time.Second)
}
