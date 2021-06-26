package service

import (
	"brevity/repository"
	"brevity/service_abstract"
)

// Common interface for all Service factories.
type AbstractServiceFactory interface {
	ScheduableService() service_abstract.AbsScheduableService
	DependencyService() service_abstract.AbsDependencyService
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
func NewLocalServiceFactory(repoFactory repository.Factory) LocalServiceFactory {
	return LocalServiceFactory{scheduableRepo: repoFactory.CreateTaskRepository(),
		dependencyRepo: repoFactory.CreateDependencyRepository()}
}

// Return a new instance of a AbsScheduableService.
func (b LocalServiceFactory) ScheduableService() service_abstract.AbsScheduableService {
	return NewScheduableService(b.scheduableRepo)
}

// Return a new instance of AbsDependencyService.
func (b LocalServiceFactory) DependencyService() service_abstract.AbsDependencyService {
	return NewDependencyService(b.dependencyRepo, b.scheduableRepo)
}
