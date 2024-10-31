package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DBinstance establishes a connection to the MongoDB database.
//
// Example:
//    client := DBinstance()
//    fmt.Println(client) // Expected output: &{... mongodb connection details ...}
//
// Returns:
//    *mongo.Client - A pointer to the mongo.Client that represents the database connection.
//
// Notes:
//    - This method will log.Fatal and exit if it cannot connect to the database, ensuring
//      that a valid connection is returned or the application is terminated.
//    - It is important to defer a call to client.Disconnect in the caller to clean up the connection
//      when it is no longer needed.
//    - The function assumes that the MongoDB server is running on localhost at port 27017.
//    - The context used for connection has a timeout of 10 seconds.
func DBinstance() *mongo.Client {
	MongoDb := "mongodb://localhost:27017"
	fmt.Print(MongoDb)

	client, err := mongo.NewClient(options.Client().ApplyURI(MongoDb))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	err = client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connected to mongodb")
	return client
}

var Client *mongo.Client = DBinstance()

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("restaurant").Collection(collectionName)

	return collection
}
