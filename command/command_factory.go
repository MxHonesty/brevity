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
}

// Factory responsible for creating instances of Command. Has a separate method
// for each type of Command. Implementation of the AbstractCommandFactory
// interface.
type Factory struct {
	
}

func (f Factory) AddContainer(container task.Container) Command {
	com := newAddContainerCommand(container)
	return com
}

func (f Factory) RemoveContainer(id uint64) Command {
	com := newRemoveContainerCommand(id)
	return com
}

func (f Factory) GetAllContainers() Command {
	panic("implement me")
}

func (f Factory) AddDependency(dependency dependency.Dependency) Command {
	panic("implement me")
}

func (f Factory) RemoveDependency(id uint64) Command {
	panic("implement me")
}

func (f Factory) GetAllDependency() Command {
	panic("implement me")
}
