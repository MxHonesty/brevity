// This module implements the commands that the server and the client can use
// between themselves.

package command

import (
	"brevity/response"
	"brevity/sessions"
	"brevity/task"
)

// need to implement a system by which the server knows when and how to send a
// response.

// Common interface for all Command instances. A command instance stores the
// necessary data for executing it's action.
//
// The execute() method:
// 		Command implements an execute() method.
//		The execute method must connect to a service
// 		once it reaches the back end. A pointer to
//		that service will be provided as an argument.
//		In this case we don't know which service will be used
//		so we provide the whole Session as an argument.
//		Execute() returns a server.Response instance.
//
// More on encoding an interface:
// https://golang.org/src/encoding/gob/example_interface_test.go
type Command interface {
	Execute(session *sessions.Session) response.Response
}

// Command for adding a given task.Container into the Repository.
type AddContainerCommand struct {
	container task.Container
}

// Create a new instance of AddContainerCommand
func newAddContainerCommand(container task.Container) *AddContainerCommand {
	return &AddContainerCommand{container: container}
}

// Forwards the request to add the task.Container to the service on the server
// Returns an empty server.Response.
func (com *AddContainerCommand) Execute(session *sessions.Session) response.Response {
	startTime := com.container.GetStartTime()
	endTime := com.container.GetEndTime()

	session.ScheduableSrv.AddContainer(startTime.Year(), startTime.Month(),
		startTime.Day(), startTime.Hour(), startTime.Minute(), endTime.Year(),
		endTime.Month(), endTime.Day(), endTime.Hour(), endTime.Minute())

	return response.Response{}
}


// Command for removing a task.Container with a give id.
type RemoveContainerCommand struct {
	id uint64
}

// Create a new RemoveContainerCommand. It takes the id of the task.Container as
// an argument.
func newRemoveContainerCommand(id uint64) *RemoveContainerCommand {
	return &RemoveContainerCommand{id: id}
}

// Forwards the request to remove a task.Container to the service on the server.
// Returns a bool server.Response.
func (com *RemoveContainerCommand) Execute(session *sessions.Session) response.Response {
	removed := session.ScheduableSrv.RemoveContainer(com.id)
	return response.NewResponse(removed)
}


// Command for getting all Containers.
type GetAllContainersCommand struct {}

// Create new GetAllContainersCommand.
func newGetAllContainersCommand() *GetAllContainersCommand {
	return &GetAllContainersCommand{}
}

// Forwards the request to get all containers.
// Returns a server.Response with a slice of task.Container.
func (com *GetAllContainersCommand) Execute(session *sessions.Session) response.Response {
	containers := session.ScheduableSrv.GetAllContainers()
	return response.NewResponse(containers)
}


// Command for Adding a new dependency.Dependency
type AddDependencyCommand struct {
	dependentId uint64
	dependentOnId []uint64
}

// Returns a server.Response with an error.
func (com *AddDependencyCommand) Execute(session *sessions.Session) response.Response {
	err := session.DepSrv.AddDependency(com.dependentId, com.dependentOnId...)
	return response.NewResponse(err)
}

// Create a new AddDependencyCommand Command.
func newAddDependencyCommand(dependentId uint64, dependentOnId ...uint64) *AddDependencyCommand {
	return &AddDependencyCommand{dependentOnId: dependentOnId, dependentId: dependentId}
}


// Command for Removing a dependency.Dependency by id.
type RemoveDependencyCommand struct {
	id uint64
}

// Create a new RemoveDependencyCommand.
func newRemoveDependencyCommand(id uint64) *RemoveDependencyCommand {
	return &RemoveDependencyCommand{id: id}
}

// Returns a bool server.Response.
func (com *RemoveDependencyCommand) Execute(session *sessions.Session) response.Response {
	removed := session.DepSrv.RemoveDependency(com.id)
	return response.NewResponse(removed)
}


// Command for Getting all dependency.Dependency.
type GetAllDependencyCommand struct {}


// Returns a server.Response with a slice of dependency.Dependency.
func (com *GetAllDependencyCommand) Execute(session *sessions.Session) response.Response {
	dependencies := session.DepSrv.GetAllDependencies()
	return response.NewResponse(dependencies)
}

// Create a new GetAllDependencyCommand.
func newGetAllDependencyCommand() *GetAllDependencyCommand {
	return &GetAllDependencyCommand{}
}


// Command that stops the connection.
type StopCommand struct {

}

// Forwards request to stop the connection. Returns an empty server.Response.
func (com *StopCommand) Execute(session *sessions.Session) response.Response {
	session.Stop()
	return response.Response{}
}
