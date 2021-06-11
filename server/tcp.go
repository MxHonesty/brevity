package server

import (
	"fmt"
	"net"
	"strconv"
	"sync"
)

// Contains data about the server.
type Server struct {
	Port uint64
	Host string
	repo ABSSessionsRepository  // Repository for storing Session instances.
	sessionsMutex *sync.Mutex  // Mutex of modifying sessions data.
	currentId uint64  // Keeps track of past sessions ids.
}

// Create a new instance of a tcp server.
func NewServer(host string, port uint64) *Server {
	return &Server{Port: port, Host: host,
		repo: NewSessionsRepository(), sessionsMutex: &sync.Mutex{},
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
	srv.repo.Add(ses)
	srv.sessionsMutex.Unlock()

	return *ses
}

// Removes the session with the given id. If no item with the given id is found
// does nothing. Returns true if the item has been removed. false if the item
// could not be found.
func (srv *Server) removeSession(id uint64) bool {
	srv.sessionsMutex.Lock()
	wasRemoved := srv.repo.Remove(id)
	srv.sessionsMutex.Unlock()
	return wasRemoved
}

// Handler for a connection.
func (srv *Server) handleServerConnection(c net.Conn) {
	session := srv.initSession()  // Initialize the session.


	srv.removeSession(session.id)  // Close the session.
	_ = c.Close()
}
