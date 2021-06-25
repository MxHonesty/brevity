package service

import (
	"brevity/client"
	"brevity/proxy_service"
	"brevity/repository"
)

// Common interface for all Service factories.
type AbstractServiceFactory interface {
	ScheduableService() AbsScheduableService
	DependencyService() AbsDependencyService
}

// Factory that returns local versions of the Services. The factory methods will
// return new services that operate on the same repositories.
type LocalServiceFactory struct {
	scheduableRepo repository.TaskRepository
	dependencyRepo repository.DependencyRepository
}

// Create a new instance of a LocalServiceFactory. Takes as argument a
// repository.Factory. The repositories that the Services will use will be
// initialized using this factory.
func NewBasicServiceFactory(repoFactory repository.Factory) LocalServiceFactory {
	return LocalServiceFactory{scheduableRepo: repoFactory.CreateTaskRepository(),
		dependencyRepo: repoFactory.CreateDependencyRepository()}
}

// Return a new instance of a AbsScheduableService.
func (b LocalServiceFactory) ScheduableService() AbsScheduableService {
	return NewScheduableService(b.scheduableRepo)
}

// Return a new instance of AbsDependencyService.
func (b LocalServiceFactory) DependencyService() AbsDependencyService {
	return NewDependencyService(b.dependencyRepo, b.scheduableRepo)
}

// Factory that creates Services that communicate with the backend.
type ProxyServiceFactory struct {
	client *client.Client
}

// Crates a new instance of ProxyServiceFactory.
func NewProxyServiceFactory(client *client.Client) ProxyServiceFactory {
	return ProxyServiceFactory{client: client}
}

// Return a new instance of a AbsScheduableService.
func (b ProxyServiceFactory) ScheduableService() AbsScheduableService {
	return proxy_service.NewProxyScheduableService(b.client)
}

// Return a new instance of AbsDependencyService.
func (b ProxyServiceFactory) DependencyService() AbsDependencyService {
	return proxy_service.NewProxyDependencyService(b.client)
}
