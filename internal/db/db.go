package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

var (
	Client *mongo.Client
)

func InitDb() {
	log.Println("Initialising Mongo Database")
	client, err := mongo.NewClient(options.Client().ApplyURI(
		fmt.Sprintf("mongodb+srv://%s:%s@all.iysnb.mongodb.net/%s?retryWrites=true&w=majority",
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_NAME")),
	))

	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	log.Printf("Connecting to: %s", os.Getenv("DB_NAME"))
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Database connection established!")
	Client = client
}

func Disconnect() {
	_ = Client.Disconnect(context.TODO())
}
