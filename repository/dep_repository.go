package repository

import (
	"brevity/dependency"
	"errors"
	"github.com/emirpasic/gods/lists/arraylist"
)

type DepRepository struct {
	elements arraylist.List
}

// Create a new Repository
func NewDepRepository() *DepRepository {
	return &DepRepository{elements: *arraylist.New()}
}

// Retrieve a copy of the element with the given id.
func (d *DepRepository) Retrieve(id uint64) (dependency.Dependency, error) {
	index, el := d.elements.Find(func(index int, value interface{}) bool {
		corValue := value.(dependency.Dependency)
		return corValue.GetId() == id
	})
	if index == -1 {
		return dependency.Dependency{}, errors.New("element with id not found")
	} else {
		return el.(dependency.Dependency), nil
	}
}

// Add a copy of the element to the list.
func (d *DepRepository) Add(dep dependency.Dependency) {
	copyVal := dep
	d.elements.Add(copyVal)
}

// Remove the element with the given id. Returns an error if the item cannot be
// found
func (d *DepRepository) Remove(id uint64) error {
	index, _ := d.elements.Find(func(index int, value interface{}) bool {
		corValue := value.(dependency.Dependency)
		return corValue.GetId() == id
	})

	if index == -1 {
		return errors.New("element not found")
	} else {
		d.elements.Remove(index)
	}
	return nil
}

// Return true if an item with this id exists in the list.
func (d *DepRepository) Find(id uint64) bool {
	index, _ := d.elements.Find(func(index int, value interface{}) bool {
		corValue := value.(dependency.Dependency)
		return corValue.GetId() == id
	})

	return index != -1
}

// Clears all the items.
func (d *DepRepository) RemoveAll() {
	d.elements.Clear()
}

// Return the list of all the items.
func (d *DepRepository) GetAll() *arraylist.List {
	return &d.elements
}

// Returns the size of the items.
func (d *DepRepository) Size() int {
	return d.elements.Size()
}

// Update an dependency.Dependency. It DOES NOT update the ID.
func (d *DepRepository) Update(id uint64, newDep dependency.Dependency) error {
	index, el := d.elements.Find(func(index int, value interface{}) bool {
		corValue := value.(dependency.Dependency)
		return corValue.GetId() == id
	})
	if index == -1 {
		return errors.New("element not found")
	} else {
		element := el.(dependency.Dependency)
		element.SetDependent(newDep.Dependent())
		element.SetDependentOn(newDep.DependentOn())
		d.elements.Set(index, element)
		return nil
	}
}
