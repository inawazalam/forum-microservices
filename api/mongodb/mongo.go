package mongodb

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//
func InitializeMongo() *mongo.Client {
	// Set client options
	var err error

	//DBURL := fmt.Sprintf("mongodb://%s:%s", DbHost, DbPort)
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	//clientOptions := options.Client.ApplyURI(DBURL)

	//fmt.Println(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName)
	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	return client
}
