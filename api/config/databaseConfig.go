package config

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/**
 * Creating Global var to be accessed from outside the package.
 */
var (
	Collection *mongo.Collection
	Ctx        = context.TODO()
)

/**
 * Setup for Database and Initializing Collection for it to be used.
 * @function setup
 */
func Setup() {
	host := "127.0.0.1"
	port := "27017"

	connectionURI := "mongodb://" + host + ":" + port + "/"
	clientOptions := options.Client().ApplyURI(connectionURI)
	client, err := mongo.Connect(Ctx, clientOptions)
	hasErr(err)

	// Ping to check if the connect is live or not.
	err = client.Ping(Ctx, nil)
	hasErr(err)

	db := client.Database("todo_db")
	log.Printf("Database Connected...")

	Collection = db.Collection("todo")
}

/**
 * To show error if anything goes wrong or shows any error.
 * @function hasError
 * @param {error} e
 */
func hasErr(e error) {
	if e != nil {
		log.Println(e)
	}
}
