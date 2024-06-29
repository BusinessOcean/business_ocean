package beserver

import "google.golang.org/grpc"

type RegisterFunc func(s grpc.ServiceRegistrar)

// RegisterSingleService registers an individual service
func RegisterSingleService(server grpc.ServiceRegistrar, registerFunc RegisterFunc) {
	registerFunc(server)
}

// RegisterMultipleServices registers multiple services in a single call
func RegisterMultipleServices(server grpc.ServiceRegistrar, registrations ...RegisterFunc) {
	for _, register := range registrations {
		register(server)
	}
}
