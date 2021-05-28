package dependency

import "brevity/task"

// A structr that models a dependency relation
// between the task and the slice of tasks.
// dependent is dependent on dependentOn.
// Visually this relationship is represented
// this way
// dependentOn ---> dependent
type Dependency struct {
	dependentOn []task.Task
	dependent task.Task
}

// Creates a Dependency with the given tasks.
// The relationship being mapped between the two
// members is this:
// dependentOn ---> dependent
// As an order of execution.
func NewDependency(dependentOn []task.Task, dependent task.Task) Dependency {
	return Dependency{dependentOn: dependentOn, dependent: dependent}
}

// Return the dependent task.Task.
func (d *Dependency) Dependent() task.Task {
	return d.dependent
}

// Sets the dependent task.Task.
func (d *Dependency) SetDependent(dependent task.Task) {
	d.dependent = dependent
}

// Return a copy of the slice of task.Task that is depended on.
func (d *Dependency) DependentOn() []task.Task {
	temp := make([]task.Task, len(d.dependentOn))
	copy(temp, d.dependentOn)
	return temp
}

// Sets the task.Task that is depended on.
// The elements are copied.
func (d *Dependency) SetDependentOn(dependentOn []task.Task) {
	temp := make([]task.Task, len(dependentOn))
	copy(temp, dependentOn)
	d.dependentOn = temp
}
