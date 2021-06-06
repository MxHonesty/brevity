package dependency

import (
	"brevity/task"
)

// Common interface for all Dependency Builders
type Builder interface {
	SetDependent(scheduable task.Scheduable)
	AddDependentOn(scheduable task.Scheduable)
	RemoveDependentOn(id uint64)
}

type ConcreteBuilder struct {
	built Dependency
}

func (c *ConcreteBuilder) SetDependent(scheduable task.Scheduable) {
	c.built.dependent = scheduable
}

func (c *ConcreteBuilder) AddDependentOn(scheduable task.Scheduable) {
	dependentOn := c.built.DependentOn()
	dependentOn = append(dependentOn, scheduable)
	c.built.SetDependentOn(dependentOn)
}

// Remove item with the given Id.
// If there is no such item this function does nothing.
func (c *ConcreteBuilder) RemoveDependentOn(id uint64) {
	// Find the index of the item with given id.
	index := -1
	for i, el := range c.built.DependentOn() {
		if el.GetId() == id {
			index = i
		}
	}

	if index != -1 {  // If element with given id found.
		slice := append(c.built.dependentOn[0:index],
			c.built.dependentOn[index+1:len(c.built.dependentOn)]...)
		c.built.SetDependentOn(slice)
	}
}

// Create a new ConcreteBuilder for a Dependency.
func NewConcreteBuilder() *ConcreteBuilder {
	return &ConcreteBuilder{built: Dependency{}}
}

// Return the instance of Dependency that was built.
// Assigns it's id as the parameter
func (c *ConcreteBuilder) GetResult(id uint64) Dependency {
	c.built.id = id
	defer func() {c.built = Dependency{}}()  // Reset the built Dependency.

	return c.built
}
