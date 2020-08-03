package controllers

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//
func (server *Server) InitilizeMongo(DbDriver, DbUser, DbPassword, DbPort, DbHost string) {

	if DbDriver == "mongodb" {
		// Set client options
		var err error

		DBURL := fmt.Sprintf("mongodb://%s:%s", DbHost, DbPort)
		clientOptions := options.Client().ApplyURI(DBURL)

		server.Client, err = mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			fmt.Println(err)
		}

		// Check the connection
		err = server.Client.Ping(context.TODO(), nil)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("Connected to MongoDB!")

	}

}
