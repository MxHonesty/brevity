package server

import (
	"encoding/gob"
	"fmt"
	"net"
	"strconv"
)

// TODO: Add a list of current sessions. Add them in the handlerMethod
// Contains data about the server.
type Server struct {
	Port uint64
	Host string
}

// Create a new instance of a tcp server.
func NewServer(host string, port uint64) *Server {
	return &Server{Port: port, Host: host}
}

// Method for starting the tcp server.
// Creates the listener and runs the main server loop.
func (srv *Server) StartServer() {
	port := strconv.FormatUint(srv.Port, 10)  // Convert port to string
	host := srv.Host
	ln, err := net.Listen("tcp", host + ":" + port) // Create Server
	if err != nil {
		fmt.Println(err)
		return
	}

	for {  // Server loop
		c, err := ln.Accept()  // Waits for request
		if err != nil {
			fmt.Println(err)
			continue
		}
		go srv.handleServerConnection(c)  // New execution thread for response
	}
}

// Handler for a connection.
func (srv *Server) handleServerConnection(c net.Conn) {
	_ = c.Close()
}

// Client for testing.
func client() {
	c, err := net.Dial("tcp", "127.0.0.1:9999")  // Establish Server Connection
	if err != nil {
		fmt.Println(err)
		return
	}

	msg := "Hello, World"
	fmt.Println("Sending", msg)
	err = gob.NewEncoder(c).Encode(msg)  // Encode and send the message
	if err != nil {
		fmt.Println(err)
	}
	_ = c.Close()
}
