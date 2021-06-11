// This module implements the commands that the server and the client can use
// between themselves.

package server

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
	execute(session *Session)
}

