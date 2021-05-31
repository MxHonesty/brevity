package service

import (
	"brevity/dependency"
	"brevity/repository"
	"brevity/task"
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

func (srv *DependencyService) AddDependency(dependentOn task.Scheduable, dependent task.Scheduable) {
	srv.depRepo.Add(dependency.NewDependency([]task.Scheduable{dependentOn},
		dependent, srv.currentId))

}

// Remove the Dependency with the given id.
// Returns true if the dependency was removed.
// Returns false if no dependency with that id was found.
func (srv *DependencyService) RemoveDependency(id uint64) bool {
	err := srv.depRepo.Remove(id)
	if err == nil {
		return true
	} else {
		return false
	}
}

// Returns a slice of copies of all the dependencies.
func (srv *DependencyService) GetAllDependencies() []dependency.Dependency {
	vec := srv.depRepo.GetAll()
	slc := vec.Values()  // Slice of values from the vector.
	temp := make([]dependency.Dependency, 0)  // For copying elements

	for _, el := range slc {
		newEl := el.(dependency.Dependency)
		copyEl := newEl.Copy()  // Make a deep copy.
		temp = append(temp, *copyEl)
	}

	return temp
}
