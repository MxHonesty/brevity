package server

import (
	"brevity/command"
	"brevity/sessions"
	"encoding/gob"
	"fmt"
	"net"
	"strconv"
	"sync"
)

// Contains data about the server.
type Server struct {
	Port          uint64
	Host          string
	repo          sessions.ABSSessionsRepository // Repository for storing Session instances.
	sessionsMutex *sync.Mutex                    // Mutex of modifying sessions data.
	currentId     uint64                         // Keeps track of past sessions ids.
}

// Create a new instance of a tcp server.
func NewServer(host string, port uint64) *Server {
	return &Server{Port: port, Host: host,
		repo: sessions.NewSessionsRepository(), sessionsMutex: &sync.Mutex{},
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
func (srv *Server) initSession() sessions.Session {
	ses := sessions.NewSession(srv.currentId) // Create new session
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

	for session.running {
		var com command.Command
		err := gob.NewDecoder(c).Decode(&com)
		if err == nil {
			response := com.Execute(&session)  // Execute the Command
			encodeErr := gob.NewEncoder(c).Encode(response)  // Send the Response.
			if encodeErr != nil {
				panic("response encoding error")
			}
		} else {
			panic(err.Error())
		}
	}

	srv.removeSession(session.id)  // Close the session.
	_ = c.Close()
}
