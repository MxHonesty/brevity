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

// Creates a Dependency with the given tasks.
// The relationship being mapped between the two
// members is this:
// dependentOn ---> dependent
// As an order of execution.
func NewDependency(dependentOn []task.Scheduable, dependent task.Scheduable,
	id uint64) Dependency {
	return Dependency{dependentOn: dependentOn, dependent: dependent, id: id}
}

// Return the dependent task.Scheduable.
func (d *Dependency) Dependent() task.Scheduable {
	return d.dependent
}

// Sets the dependent task.Scheduable.
func (d *Dependency) SetDependent(dependent task.Scheduable) {
	d.dependent = dependent
}

// Return a copy of the slice of task.Scheduable that is depended on.
func (d *Dependency) DependentOn() []task.Scheduable {
	temp := make([]task.Scheduable, len(d.dependentOn))
	copy(temp, d.dependentOn)
	return temp
}

// Sets the task.Scheduable that is depended on.
// The elements are copied.
func (d *Dependency) SetDependentOn(dependentOn []task.Scheduable) {
	temp := make([]task.Scheduable, len(dependentOn))
	copy(temp, dependentOn)
	d.dependentOn = temp
}

// Return the id of the Dependency
func (d *Dependency) GetId() uint64 {
	return d.id
}

// Returns a copy of this Dependency.
func (d *Dependency) Copy() Dependency {
	sliceCopy := make([]task.Scheduable, len(d.DependentOn()))
	copy(sliceCopy, d.DependentOn())  // Copied dependent slice.
	newEl := NewDependency(sliceCopy, d.Dependent(), d.id)
	return newEl
}
