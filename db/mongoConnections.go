package DB

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBinstance() *mongo.Client {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	MongoDb := os.Getenv("MONGODB_URL")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(MongoDb))
	if err != nil {
		log.Fatal(err)
	}
	// Optionally ping to verify connection
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
		return nil
	}
	fmt.Println("Connected to MongoDB successfully")
	return client
}

var Client *mongo.Client = DBinstance()

func OpennCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("BunLab").Collection(collectionName)
	return collection

}
