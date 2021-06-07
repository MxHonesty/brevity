package server

import "testing"

// Test case for adding a Session to the Repository.
func TestSessionsRepository_Add(t *testing.T) {
	repo := NewSessionsRepository()
	repo.Add(NewSession(0))
	repo.Add(NewSession(1))
	if repo.Size() != 2 {
		t.Errorf("Expected size 2, got %d", repo.Size())
	}

	if repo.GetAll()[0].GetId() != 0 {
		t.Errorf("Expected to have id 0, got %d", repo.GetAll()[0].GetId())
	}

	if repo.GetAll()[1].GetId() != 1 {
		t.Errorf("Expected to have id 1, got %d", repo.GetAll()[1].GetId())
	}
}

// Test case for Getting a single Session from the Repository.
func TestSessionsRepository_Get(t *testing.T) {
	repo := NewSessionsRepository()
	repo.Add(NewSession(0))
	repo.Add(NewSession(1))

	el, found := repo.Get(1)
	if !found{
		t.Error("Expected to find the Session with id 1")
	}

	if el.GetId() != 1 {
		t.Errorf("Expected id to be 1, got %d", el.GetId())
	}

	el, found = repo.Get(10)  // Expected to find nothing
	if found {
		t.Error("Expected not to find with id 10")
	}
}

func TestSessionsRepository_Remove(t *testing.T) {
	repo := NewSessionsRepository()
	repo.Add(NewSession(0))
	repo.Add(NewSession(1))
	repo.Add(NewSession(2))

	deleted := repo.Remove(1)
	if !deleted {
		t.Error("Expected to successfully remove item with id 1")
	}

	if repo.Size() != 2 {
		t.Errorf("Expected to have size 2 after removing an item, got %d", repo.Size())
	}

	// Try to remove item that does not exist.
	deleted = repo.Remove(10)
	if deleted {
		t.Error("Expected not to find item with id 10")
	}

	if repo.Size() != 2 {
		t.Errorf("Expected to have size 2 after not removing an item, got %d", repo.Size())
	}
}
