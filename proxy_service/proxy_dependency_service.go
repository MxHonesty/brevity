// This module implements a Proxy of the DependencyService defined and
// implemented inside the service package. The Proxy forwards the requests in the
// form of command.Command instances to the server.
package proxy_service

import (
	"brevity/client"
	"brevity/dependency"
)

// TODO: implement
// Every method creates a command.Command instance.
// Has a Sender member that is responsible for encoding
// and sending the command.Command to the server.
// TODO: Create CommandSender inside client package.
type ProxyDependencyService struct {
	client *client.Client  // Used for sending the commands to the server.
}

func (p ProxyDependencyService) AddDependency(dependentId uint64, dependentOnId ...uint64) error {
	panic("implement me")
}

func (p ProxyDependencyService) RemoveDependency(id uint64) bool {
	panic("implement me")
}

func (p ProxyDependencyService) GetAllDependencies() []dependency.Dependency {
	panic("implement me")
}
