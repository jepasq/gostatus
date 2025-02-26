package main

/// A basic implementation of the service listener ecosystem

import (
	"fmt"
	"log"
	//	"io/ioutil"
	"context"
	"net/http"
	"encoding/json"
	
	"go.mongodb.org/mongo-driver/bson"
)


type Service struct {
	Name  string
	Stype string        `bson:"service_id,omitempty"`
}

/// A listener abstract type used to monitor a service
type ServiceListener interface {
	/// The type of microservice (suystemctl, other...)
	typeName() string
	
	/// Return the current status of this service
	/// For instance as a string but surely later as a plain structured type
	status() string
}

/// Abstract type of a listener saved to the database
type SavedListener interface {
	/// Mandatory string component
	name() string
}

/// A systemctl-based service
type SystemCtlListener struct {
	/// The service you can call with service status {serviceName}
	serviceName string
}

func (scl SystemCtlListener) typeName() string {
	return "SystemCtl"   // But could be type instead
}

func (l SystemCtlListener) status() string {
    return "result of service status " + l.serviceName
}

func ServiceListenerAdd(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// db is the objet returned by Connect(). We have to cast it
	db := ctx.Value("db").(Database) 
	/*	if db == nil {
		fmt.Printf("ERR: db from Context is nil. Can't continue.")
		return
	}
	*/
	id := r.URL.Query().Get("name")
	typ:= r.URL.Query().Get("type")
	
	fmt.Printf("%s: got /admin request\nbody:\n  name:%s\n  type:%s\n",
		ctx.Value(keyServerAddr),
		id, typ)

	
	
	// Insert value
	newService := Service{Name: id, Stype: typ}
	
	coll := db.Client.Database("gostatus").Collection("services")
	result, err := coll.InsertOne(context.TODO(), newService)
	if err != nil {
		panic(err)
	}
	
	fmt.Printf("New service insertion result '%s'\n", result)
}

func GetServicesJson(w http.ResponseWriter, r *http.Request) {
	db := r.Context().Value("db").(Database) 
	coll := db.Client.Database("gostatus").Collection("services")
	// ctx := r.Context()

	//find records
	//pass these options to the Find method
	//	findOptions := options.Find()
	//Set the limit of the number of record to find
	//	findOptions.SetLimit(5)
	//Define an array in which you can store the decoded documents
	var results []Service

	//Passing the bson.D{{}} as the filter matches docs in the collection
	cur, err := coll.Find(context.TODO(), bson.D{{}}/*, findOptions*/)
	if err !=nil {
		log.Fatal(err)
	}

	// results to JSON
	for cur.Next(context.TODO()) {
		//Create a value into which the single document can be decoded
		var elem Service
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		results =append(results, elem)
	}
	
	b, err := json.Marshal(results)
	if err != nil {
		fmt.Println(err)
		return
	}
	
	//	writeTemplate(w, "%s", "{}")
	fmt.Fprintf(w, "%s", string(b))
}
