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
	Uri    string
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
	uri = os.Getenv("MONGODB_URI")
	ret.Uri = uri
	
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)
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
