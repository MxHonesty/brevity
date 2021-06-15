// This module implements the commands that the server and the client can use
// between themselves.

package command

import (
	"brevity/server"
	"brevity/task"
)

// TODO: coordinate responses from Execute method.
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
//
// More on encoding an interface:
// https://golang.org/src/encoding/gob/example_interface_test.go
type Command interface {
	Execute(session *server.Session)
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
func (com *AddContainerCommand) Execute(session *server.Session) {
	startTime := com.container.GetStartTime()
	endTime := com.container.GetEndTime()

	session.ScheduableSrv.AddContainer(startTime.Year(), startTime.Month(),
		startTime.Day(), startTime.Hour(), startTime.Minute(), endTime.Year(),
		endTime.Month(), endTime.Day(), endTime.Hour(), endTime.Minute())
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
func (com *RemoveContainerCommand) Execute(session *server.Session) {
	session.ScheduableSrv.RemoveContainer(com.id)
}


// Command for getting all Containers.
type GetAllContainersCommand struct {}

// Create new GetAllContainersCommand.
func NewGetAllContainersCommand() *GetAllContainersCommand {
	return &GetAllContainersCommand{}
}

func (com *GetAllContainersCommand) Execute(session *server.Session) {
	session.ScheduableSrv.GetAllContainers()
}


