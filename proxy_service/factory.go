package proxy_service

import (
	"brevity/client"
	"brevity/service_abstract"
)

// Factory that creates Services that communicate with the backend.
type ProxyServiceFactory struct {
	client *client.Client
}

// Crates a new instance of ProxyServiceFactory.
func NewProxyServiceFactory(client *client.Client) ProxyServiceFactory {
	return ProxyServiceFactory{client: client}
}

// Return a new instance of a AbsScheduableService.
func (b ProxyServiceFactory) ScheduableService() service_abstract.AbsScheduableService {
	return NewProxyScheduableService(b.client)
}

// Return a new instance of AbsDependencyService.
func (b ProxyServiceFactory) DependencyService() service_abstract.AbsDependencyService {
	return NewProxyDependencyService(b.client)
}
