package service

import (
	"brevity/dependency"
	"brevity/repository"
	"errors"
	"fmt"
)

// Struct responsible for serving dependency.Dependency
// related functionalities.
type DependencyService struct {
	taskRepo repository.TaskRepository
	depRepo repository.DependencyRepository
	currentId uint64  // Used for storing the id of the next dependency.Dependency
}

// Create a new DependencyService from a repository.Factory.
func NewDependencyService(depRepo repository.DependencyRepository,
	taskRepo repository.TaskRepository) *DependencyService {
	return &DependencyService{taskRepo: taskRepo, depRepo: depRepo, currentId: 0}
}

// Adds a dependency between a task.Scheduable and a list of task.Scheduable
// elements. Returns an error depending on the success of the operation.
//
// Params:
// 		dependentId - the id of the dependent item
// 		dependentOnId - a list of id's for the items that dependent
// 		depends on
//
// Errors:
//		returns errors if either the dependent is not found or if any of the
//		dependentOn are not found. The operation is not done if such error
// 		occurs. Otherwise returns nil.
func (srv *DependencyService) AddDependency(dependentId uint64, dependentOnId ...uint64) error {
	builder := dependency.NewConcreteBuilder()

	// Find dependentId
	// Find a list of all dependentsOn
	dependent, err := srv.taskRepo.Retrieve(dependentId)
	if err != nil {
		return errors.New("could not find item for dependentId")
	} else {
		builder.SetDependent(dependent)
	}

	for _, itemId := range dependentOnId {
		tsk, err := srv.taskRepo.Retrieve(itemId)
		if err != nil {
			return errors.New(fmt.Sprintf("could not find item for dependentOnId %d", itemId))
		} else {
			builder.AddDependentOn(tsk)
		}
	}


	dep := builder.GetResult(srv.currentId)
	srv.currentId++
	srv.depRepo.Add(dep)
	return nil
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
