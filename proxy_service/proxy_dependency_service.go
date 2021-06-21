// This module implements a Proxy of the DependencyService defined and
// implemented inside the service package. The Proxy forwards the requests in the
// form of command.Command instances to the server.
package proxy_service

import (
	"brevity/client"
	"brevity/command"
	"brevity/dependency"
)

// Every method creates a command.Command instance.
// Has a Sender member that is responsible for encoding
// and sending the command.Command to the server.
type ProxyDependencyService struct {
	client *client.Client  // Used for sending the commands to the server.
	commandFactory command.AbstractCommandFactory
}

// Create a new instance of ProxyDependencyService. Gets pointer to client.Client
// as argument.
func NewProxyDependencyService(client *client.Client) *ProxyDependencyService {
	return &ProxyDependencyService{client: client,
		commandFactory: command.Factory{}}
}

func (p ProxyDependencyService) AddDependency(dependentId uint64, dependentOnId ...uint64) error {
	com := p.commandFactory.AddDependency(dependentId, dependentOnId)
	resp, _ := p.client.SendCommand(com)
	return resp.Data.(error)  // We know that the return type is an error.
}

func (p ProxyDependencyService) RemoveDependency(id uint64) bool {
	com := p.commandFactory.RemoveDependency(id)
	resp, _ := p.client.SendCommand(com)
	return resp.Data.(bool)  // The return type of is bool.
}

func (p ProxyDependencyService) GetAllDependencies() []dependency.Dependency {
	com := p.commandFactory.GetAllDependency()
	resp, _ := p.client.SendCommand(com)
	return resp.Data.([]dependency.Dependency)  // Return type is Dependency slice
}
