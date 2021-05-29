package repository

import (
	"brevity/dependency"
	"github.com/emirpasic/gods/lists/arraylist"
)

// Common interface for all dependency.Dependency
// repositories.
type DependencyRepository interface {
	Retrieve(id uint64) (dependency.Dependency, error)
	Add(dep dependency.Dependency)
	Remove(id uint64) error
	Find(id uint64) bool
	RemoveAll()
	GetAll() *arraylist.List
	Size() int
	Update(id uint64, newDep dependency.Dependency) error
}
