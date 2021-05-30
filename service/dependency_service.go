package service

import (
	"brevity/dependency"
	"brevity/repository"
)

// Struct responsible for serving dependency.Dependency
// related functionalities.
type DependencyService struct {
	taskRepo repository.TaskRepository
	depRepo repository.DependencyRepository
	currentId uint64  // Used for storing the id of the next dependency.Dependency
}

// Create a new DependencyService from a repository.Factory.
func NewDependencyService(factory repository.Factory) *DependencyService {
	taskRepo := factory.CreateTaskRepository()
	depRepo := factory.CreateDependencyRepository()
	return &DependencyService{taskRepo: taskRepo, depRepo: depRepo, currentId: 0}
}

func (srv *DependencyService) AddDependency() {
	panic("implement me")
}

// Remove the Dependency with the given id.
// Returns true if the dependency was removed.
// Returns false if no dependency with that id was found.
func (srv *DependencyService) RemoveDependency(id uint64) bool {
	panic("implement me")
}

// Returns a slice of copies of all the dependencies.
func (srv *DependencyService) GetAllDependencies() []dependency.Dependency {
	panic("implement me")
}
