package repository

import (
	"brevity/dependency"
	"brevity/task"
	"testing"
)

// Test case for adding dependency.Dependency.
func TestAddDepRepository(t *testing.T) {
	repo := NewDepRepository()

	dep1 := dependency.NewDependency([]task.Task{task.NewTask("A", "A")},
	task.NewTask("A", "A"), 0)
	dep2 := dependency.NewDependency([]task.Task{task.NewTask("A", "A")},
	task.NewTask("B", "B"), 1)

	repo.Add(dep1)
	repo.Add(dep2)

	size := repo.Size()
	if size != 2 {
		t.Errorf("Expected size 2, got %d", size)
	}

	dep1.SetDependent(task.NewTask("AA", "AA"))
	// We expect the repo item not to change
	item, err := repo.Retrieve(0)
	if err != nil {
		t.Error("Existing item not found")
	} else {
		itemDep := item.Dependent()
		if itemDep.GetDescription() != "A" {
			t.Errorf("Expected 'A', got %s", itemDep.GetDescription())
		}
	}
}

// Test case for removing dependency.Dependency.
func TestRemoveDepRepository(t *testing.T) {
	repo := NewDepRepository()

	dep1 := dependency.NewDependency([]task.Task{task.NewTask("A", "A")},
		task.NewTask("A", "A"), 0)
	dep2 := dependency.NewDependency([]task.Task{task.NewTask("A", "A")},
		task.NewTask("B", "B"), 1)
	dep3 := dependency.NewDependency([]task.Task{task.NewTask("A", "A")},
		task.NewTask("C", "C"), 2)

	repo.Add(dep1)
	repo.Add(dep2)
	repo.Add(dep3)

	deleted := repo.Remove(1)
	if deleted != nil {
		t.Error("Item could not be deleted.")
	}

	size := repo.Size()
	if size != 2 {
		t.Errorf("Exptected size 2, got %d", size)
	}

	if repo.Find(1) {
		t.Error("Found deleted item")
	}

	if !repo.Find(0) || !repo.Find(2) {
		t.Error("Could not find existing item")
	}

	repo.RemoveAll()
	size = repo.Size()
	if size != 0 {
		t.Errorf("Expected size 0, got %d", size)
	}

}

// Test case for taking the list of elements.
func TestGetAllRepository(t *testing.T) {
	repo := NewDepRepository()

	dep1 := dependency.NewDependency([]task.Task{task.NewTask("A", "A")},
		task.NewTask("A", "A"), 0)
	dep2 := dependency.NewDependency([]task.Task{task.NewTask("A", "A")},
		task.NewTask("B", "B"), 1)
	dep3 := dependency.NewDependency([]task.Task{task.NewTask("A", "A")},
		task.NewTask("C", "C"), 2)

	repo.Add(dep1)
	repo.Add(dep2)
	repo.Add(dep3)

	list := repo.GetAll()
	if list.Size() != 3 {
		t.Errorf("Expected size 3, got %d", list.Size())
	}
}

// Test the update operation on DepRepository.
func TestUpdateDepRepository(t *testing.T) {
	repo := NewDepRepository()

	dep1 := dependency.NewDependency([]task.Task{task.NewTask("A", "A")},
		task.NewTask("A", "A"), 0)
	dep2 := dependency.NewDependency([]task.Task{task.NewTask("A", "A")},
		task.NewTask("B", "B"), 1)
	dep3 := dependency.NewDependency([]task.Task{task.NewTask("A", "A")},
		task.NewTask("C", "C"), 2)

	repo.Add(dep1)
	repo.Add(dep2)

	replaced := repo.Update(1, dep3)
	if replaced != nil {
		t.Error("Item existed, but not found")
	}

	found := repo.Find(1)  // we expect not to find id 1
	if !found {
		t.Error("Expected true, got false")
	}

	found = repo.Find(2)
	if found {
		t.Error("Expected false, got true")
	}

	// Testing is copied.
	dep3.SetDependent(task.NewTask("AAA", "AAA"))
	ret, err := repo.Retrieve(1)
	if err != nil {
		t.Error("Error in finding id 2.")
	} else {
		retDep := ret.Dependent()
		if retDep.GetDescription() != "C" {
			t.Errorf("Expected description 'C', got %s", retDep.GetDescription())
		}
	}
}

func TestUpdateNotExistingDepRepository(t *testing.T) {
	repo := NewDepRepository()

	dep1 := dependency.NewDependency([]task.Task{task.NewTask("A", "A")},
		task.NewTask("A", "A"), 0)
	dep2 := dependency.NewDependency([]task.Task{task.NewTask("A", "A")},
		task.NewTask("B", "B"), 1)

	repo.Add(dep1)
	repo.Add(dep2)

	err := repo.Update(2, dependency.NewDependency([]task.Task{task.NewTask("A", "A")},
		task.NewTask("A", "A"), 2))
	if err == nil {
		t.Error("Expected error, got none.")
	}
}
