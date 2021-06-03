package service

import (
	"brevity/dependency"
	"brevity/repository"
	"brevity/task"
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
func NewDependencyService(factory repository.Factory) *DependencyService {
	taskRepo := factory.CreateTaskRepository()
	depRepo := factory.CreateDependencyRepository()
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
	// Find dependentId
	// Find a list of all dependentsOn
	dependent, err := srv.taskRepo.Retrieve(dependentId)
	if err != nil {
		return errors.New("could not find item for dependentId")
	}

	var tsk task.Scheduable
	slc := make([]task.Scheduable, 0)
	for _, itemId := range dependentOnId {  // Range over ids and add the respective
		// Tasks inside the slice.
		tsk, err = srv.taskRepo.Retrieve(itemId)
		if err != nil {
			return errors.New(fmt.Sprintf("could not find item for dependentOnId %d", itemId))
		} else {
			slc = append(slc, tsk)
		}
	}

	dep := dependency.NewDependency(slc, dependent, srv.currentId)
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
