package main

/// A basic implementation of the service listener ecosystem

import (
	"fmt"
	//	"io/ioutil"
	"net/http"
)


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

	db := ctx.Value("db");
	if db == nil {
		fmt.Printf("ERR: db from Context is nil. Can't continue.")
		return
	}
	
	id := r.URL.Query().Get("name")
	typ:= r.URL.Query().Get("type")
	
	fmt.Printf("%s: got /admin request\nbody:\n  name:%s\n  type:%s\n",
		ctx.Value(keyServerAddr),
		id, typ)

}

