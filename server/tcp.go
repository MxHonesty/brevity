package server

import (
	"encoding/gob"
	"fmt"
	"net"
	"strconv"
	"sync"
)

// Contains data about the server.
// TODO: Testing for Session management.
type Server struct {
	Port uint64
	Host string
	currentSessions []Session  // List of current sessions
	sessionsMutex *sync.Mutex  // Mutex of modifying sessions data.
	currentId uint64  // Keeps track of past sessions ids.
}

// Create a new instance of a tcp server.
func NewServer(host string, port uint64) *Server {
	return &Server{Port: port, Host: host,
		currentSessions: nil, sessionsMutex: &sync.Mutex{},
		currentId: 0}
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

// Initialises a new Session.
// Returns the created Session.
func (srv *Server) initSession() Session {
	ses := NewSession(srv.currentId)  // Create new session
	srv.currentId++

	srv.sessionsMutex.Lock()
	srv.currentSessions = append(srv.currentSessions, *ses)  // Append the item
	// to the list inside the critical zone.
	srv.sessionsMutex.Unlock()

	return *ses
}

// Removes the session with the given id.
// If no item with the given id is found does nothing.
// Returns true if the item has been removed.
// false if the item could not be found.
func (srv *Server) removeSession(id uint64) bool {
	srv.sessionsMutex.Lock()
	// Find the index of the item to remove.
	index := -1
	for i, el := range srv.currentSessions {
		if el.id == id {
			index = i
		}
	}
	if index == -1 {
		return false
	}

	srv.currentSessions = append(srv.currentSessions[:index],
		srv.currentSessions[index+1:]...)  // Remove the found item
	srv.sessionsMutex.Unlock()
	return true
}

// Handler for a connection.
func (srv *Server) handleServerConnection(c net.Conn) {
	session := srv.initSession()  // Initialize the session.


	srv.removeSession(session.id)  // Close the session.
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
