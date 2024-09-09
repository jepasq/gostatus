package main

// Database connection and related

import (
	"os"
	"log"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Package constants
const (
	// the local mongo URI given as example
	local_URI = "mongodb://localhost:27017/<database>"
	doc_URI = "www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable"
)


type Database struct {
	Client *mongo.Client
}

func CheckConnString() {
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("Set your 'MONGODB_URI' environment variable. \n" +
			"  Should be something like '" + local_URI + "'\n" +
			"  See: " + doc_URI)
	}
}

func Connect() (Database) {	
	ret := Database{}

	var uri string
	if uri = os.Getenv("MONGODB_URI"); uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environment variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
	
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	ret.Client = client
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	return ret
}
