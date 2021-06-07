package server

import (
	"testing"
	"time"
)

// Test case for creating a new Session.
// Also tests that the created Services interact successfully.
func TestNewSession(t *testing.T) {
	ses := NewSession(0)
	ses.scheduableSrv.AddContainer(0, time.January, 0, 0, 0,
		0, time.January, 0,0,0)

	size := len(ses.scheduableSrv.GetAllContainers())
	if size != 1 {
		t.Errorf("Expected to have size 1, got %d", size)
	}

	err := ses.depSrv.AddDependency(0)
	if err != nil {
		t.Error("Expected to add Dependency Successfully")
	}

	sizeDep := len(ses.depSrv.GetAllDependencies())
	if sizeDep != 1 {
		t.Errorf("Expected size 1, got %d", sizeDep)
	}
}

// Test case for creating a new Server.
func TestNewServer(t *testing.T) {
	server := NewServer("", 8000)
	if server.Port != 8000 || server.Host != "" {
		t.Error("Server did not construct properly.")
	}

	if len(server.repo.GetAll()) != 0 {
		t.Error("Expected session list to be empty")
	}
}

func TestServerInitSession(t *testing.T) {
	server := NewServer("", 8000)
	server.initSession()
	server.initSession()
	server.initSession()

	size := len(server.repo.GetAll())
	if size != 3 {
		t.Errorf("Expected size 3, got %d", size)
	}
}

func TestServerRemoveSession(t *testing.T) {
	server := NewServer("", 8000)
	server.initSession()
	server.initSession()
	server.initSession()

	server.removeSession(1)
	size := len(server.repo.GetAll())
	if size != 2 {
		t.Errorf("Expected size 2, got %d", size)
	}

	// Check the correct item was deleted.
	firstId := server.repo.GetAll()[0].GetId()
	secondId := server.repo.GetAll()[1].GetId()

	if firstId != 0 {
		t.Errorf("Expected first element to have id 0, got %d", firstId)
	}
	if secondId != 2 {
		t.Errorf("Expected first element to have id 2, got %d", secondId)
	}
}
