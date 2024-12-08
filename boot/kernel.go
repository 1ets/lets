package boot

import (
	"github.com/1ets/lets/drivers"
	"github.com/1ets/lets/frameworks"
	"github.com/1ets/lets/loader"
)

// List of initializer
var Initializer = []func(){
	loader.Environment,
	loader.Logger,
	loader.Network,
	loader.Replica,
}

// List of framework that start on lets
var Servers = []func(){
	frameworks.Http,
	frameworks.Grpc,
	frameworks.RabbitMQ,
	frameworks.Tcp,
	frameworks.WebSocket,
	drivers.MySQL,
	drivers.Redis,
	drivers.MongoDB,
}

// Add initialization function and run before application starting
func AddInitializer(init func()) {
	Initializer = append(Initializer, init)
}

// Add initialization function and run before application starting
func AddServers(server func()) {
	Servers = append(Servers, server)
}

func AddStopper(stopper func()) {
	loader.Stopper = append(loader.Stopper, stopper)
}

// Bootstrap vars and configuration
func OnInit() {
	Initializer = append(Initializer, loader.Launching)
	for _, initializer := range Initializer {
		initializer()
	}
}

// Bootstrap frameworks
func OnMain() {
	for _, runner := range Servers {
		go runner()
	}

	loader.RunningForever()
}
