package main

/// A basic implementation of a service listener

/// A listener abstract type used to monitor a service
type ServiceListener interface {
	/// Return the current status of this service
	/// For instance as a string but surely later as a plain structured type
	status() string
}


/// A systemctl-based service
type SystemCtlListener struct {
	/// The service you can call with service status {serviceName}
	serviceName string
}

func (l SystemCtlListener) status() string {
    return "result of service status " + l.serviceName
}

