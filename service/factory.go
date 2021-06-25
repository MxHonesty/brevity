package service

import "brevity/repository"

// Common interface for all Service factories.
type AbstractServiceFactory interface {
	ScheduableService() AbsScheduableService
	DependencyService() AbsDependencyService
}

// Factory that returns local versions of the Services. The factory methods will
// return new services that operate on the same repositories.
type BasicServiceFactory struct {
	scheduableRepo repository.TaskRepository
	dependencyRepo repository.DependencyRepository
}

// Create a new instance of a BasicServiceFactory. Takes as argument a
// repository.Factory. The repositories that the Services will use will be
// initialized using this factory.
func NewBasicServiceFactory(repoFactory repository.Factory) BasicServiceFactory {
	return BasicServiceFactory{scheduableRepo: repoFactory.CreateTaskRepository(),
		dependencyRepo: repoFactory.CreateDependencyRepository()}
}

// Return a new instance of a AbsScheduableService.
func (b BasicServiceFactory) ScheduableService() AbsScheduableService {
	return NewScheduableService(b.scheduableRepo)
}

func (b BasicServiceFactory) DependencyService() AbsDependencyService {
	return NewDependencyService(b.dependencyRepo, b.scheduableRepo)
}

// Factory that creates Services that communicate with the backend.
type BackendService struct {

}

func (b BackendService) ScheduableService() AbsScheduableService {
	panic("implement me")
}

func (b BackendService) DependencyService() AbsDependencyService {
	panic("implement me")
}
