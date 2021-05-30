package service

import (
	"brevity/repository"
	"testing"
	"time"
)

// Test case for add functionality.
func TestAddScheduableService(t *testing.T) {
	srv := NewScheduableService(repository.NewInMemoryRepositoryFactory())
	srv.AddContainer(0, time.January, 0, 0, 0,
		0, time.January, 0, 0, 0)

	srv.AddContainer(0, time.January, 0, 0, 0,
		0, time.January, 0, 0, 0)

	srv.AddContainer(0, time.January, 0, 0, 0,
		0, time.January, 0, 0, 0)

	elements := srv.GetAllContainers()
	if len(elements) != 3 {
		t.Errorf("Expected 3, got %d", len(elements))
	}
}

// Test case for remove functionality.
func TestRemoveScheduableService(t *testing.T) {
	srv := NewScheduableService(repository.NewInMemoryRepositoryFactory())
	srv.AddContainer(0, time.January, 0, 0, 0,
		0, time.January, 0, 0, 0)

	srv.AddContainer(0, time.January, 0, 0, 0,
		0, time.January, 0, 0, 0)

	srv.AddContainer(0, time.January, 0, 0, 0,
		0, time.January, 0, 0, 0)

	del := srv.RemoveContainer(4)
	if del {
		t.Error("Removed non existent item")
	}

	del = srv.RemoveContainer(0)
	if !del {
		t.Error("Did not remove container")
	}

	if srv.repo.Size() != 2 {
		t.Errorf("Expectd size 2, got %d", srv.repo.Size())
	}
}
