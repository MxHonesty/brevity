package command

import (
	"brevity/dependency"
	"brevity/task"
)

// Common interface for all Command Factories. A Command factory takes the
// necessary arguments for creating the Command
type AbstractCommandFactory interface {
	AddContainer(container task.Container) Command
	RemoveContainer(id uint64) Command
	GetAllContainers() Command
	AddDependency(dependency dependency.Dependency) Command
	RemoveDependency(id uint64) Command
	GetAllDependency() Command
	GetStop() Command
}

// Factory responsible for creating instances of Command. Has a separate method
// for each type of Command. Implementation of the AbstractCommandFactory
// interface.
type Factory struct {
	
}

// Returns a AddContainerCommand for the given task.Container.
func (f Factory) AddContainer(container task.Container) Command {
	com := newAddContainerCommand(container)
	return com
}

// Returns a RemoveContainerCommand for the given id.
func (f Factory) RemoveContainer(id uint64) Command {
	com := newRemoveContainerCommand(id)
	return com
}

// Returns a GetAllContainersCommand.
func (f Factory) GetAllContainers() Command {
	com := newGetAllContainersCommand()
	return com
}

// Returns a AddDependencyCommand for a given dependency.Dependency.
func (f Factory) AddDependency(dependency dependency.Dependency) Command {
	com := newAddDependencyCommand(dependency)
	return com
}

// Returns a RemoveDependencyCommand for a given id.
func (f Factory) RemoveDependency(id uint64) Command {
	com := newRemoveDependencyCommand(id)
	return com
}

// Returns a GetAllDependencyCommand.
func (f Factory) GetAllDependency() Command {
	com := newGetAllDependencyCommand()
	return com
}

// Returns a StopCommand.
func (f Factory) GetStop() Command {
	com := &StopCommand{}
	return com
}
