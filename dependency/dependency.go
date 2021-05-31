package dependency

import "brevity/task"

// A structr that models a dependency relation
// between the task and the slice of tasks.
// dependent is dependent on dependentOn.
// Visually this relationship is represented
// this way
// dependentOn ---> dependent
type Dependency struct {
	dependentOn []task.Scheduable
	dependent task.Scheduable
	id uint64  // Unique id of this Dependency
}

// Creates a Dependency with the given task.Scheduable.
// The relationship being mapped between the two
// members is represented as such:
// dependentOn ---> dependent
// As an order of execution.
func NewDependency(dependentOn []task.Scheduable, dependent task.Scheduable,
	id uint64) Dependency {
	dependentOnCopy := make([]task.Scheduable, 0)
	for _, el := range dependentOn {
		dependentOnCopy = append(dependentOnCopy, el.Copy())
	}

	return Dependency{dependentOn: dependentOnCopy, dependent: dependent.Copy(), id: id}
}

// Return the dependent task.Scheduable.
func (d *Dependency) Dependent() task.Scheduable {
	return d.dependent
}

// Sets the dependent task.Scheduable. Creates a copy.
func (d *Dependency) SetDependent(dependent task.Scheduable) {
	d.dependent = dependent.Copy()
}

// Return the slice task.Scheduable that is depended on.
// This does not make a new copy.
func (d *Dependency) DependentOn() []task.Scheduable {
	return d.dependentOn
}

// Sets the task.Scheduable that is depended on.
// The elements are copied.
func (d *Dependency) SetDependentOn(dependentOn []task.Scheduable) {
	temp := make([]task.Scheduable, 0)
	for _, el := range dependentOn {
		temp = append(temp, el.Copy())
	}
	d.dependentOn = temp
}

// Return the id of the Dependency
func (d *Dependency) GetId() uint64 {
	return d.id
}

// Returns a copy of this Dependency.
func (d *Dependency) Copy() *Dependency {
	sliceCopy := make([]task.Scheduable, 0)
	for _, el := range d.dependentOn {
		sliceCopy = append(sliceCopy, el.Copy())
	}
	newEl := NewDependency(sliceCopy, d.Dependent().Copy(), d.id)
	return &newEl
}
