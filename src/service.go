package main

/// A basic implementation of the service listener ecosystem

import (
	"fmt"
	//	"io/ioutil"
	"context"
	"net/http"
)


type Service struct {
	Name  string
	Stype string        `bson:"restaurant_id,omitempty"`
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
	
	coll := db.Client.Database("goStatus").Collection("services")
	result, err := coll.InsertOne(context.TODO(), newService)
	if err != nil {
		panic(err)
	}
	
	fmt.Printf("New service insertion result '%s'\n", result)
}

