package service

import (
	"brevity/repository"
	"brevity/task"
	"testing"
	"time"
)

func createInMemoryDepRepo() repository.DependencyRepository {
	factory := repository.NewInMemoryRepositoryFactory()
	return factory.CreateDependencyRepository()
}

func createInMemorySchRepo() repository.TaskRepository {
	factory := repository.NewInMemoryRepositoryFactory()
	return factory.CreateTaskRepository()
}

// Test case for making sure the Service is created successfully.
func TestNewDependencyService(t *testing.T) {
	srv := NewDependencyService(createInMemoryDepRepo(), createInMemorySchRepo())
	if srv.currentId != 0 {
		t.Errorf("Expected currentID of 0, got %d", srv.currentId)
	}
}

func TestDependencyService_AddDependency(t *testing.T) {
	srv := NewDependencyService(createInMemoryDepRepo(), createInMemorySchRepo())

	// Add a few tasks
	srv.taskRepo.Add(task.NewContainer(0, time.Now(), time.Now()))
	srv.taskRepo.Add(task.NewContainer(1, time.Now(), time.Now()))
	srv.taskRepo.Add(task.NewContainer(2, time.Now(), time.Now()))

	// Create dependency between first two and last.
	err := srv.AddDependency(0, 1, 2)
	if err != nil {
		t.Error("Expected to find the given ids")
	}

	dep, depErr := srv.depRepo.Retrieve(0)
	if depErr != nil {
		t.Error("Expected to find the newly created dependency")
	} else {
		foundId := dep.Dependent().GetId()
		if foundId != 0 {
			t.Errorf("Expected id of dependent to be 0, got %d", foundId)
		}
	}
}

func TestDependencyService_RemoveDependency(t *testing.T) {
	srv := NewDependencyService(createInMemoryDepRepo(), createInMemorySchRepo())

	// Add a few tasks
	srv.taskRepo.Add(task.NewContainer(0, time.Now(), time.Now()))
	srv.taskRepo.Add(task.NewContainer(1, time.Now(), time.Now()))
	srv.taskRepo.Add(task.NewContainer(2, time.Now(), time.Now()))

	// Create dependency between first two and last.
	err := srv.AddDependency(0, 1, 2)
	if err != nil {
		t.Error("Expected to find the given ids")
	}

	err = srv.AddDependency(1, 2)  // Made 1 dependent on 2
	if err != nil {
		t.Error("Expected to find the given ids")
	}

	removed := srv.RemoveDependency(0)  // We Remove the first dependency
	if !removed {
		t.Error("Expected to remove the item")
	}

	size := srv.depRepo.Size()
	if size != 1{
		t.Errorf("Expected size of depRepo 1, got %d", size)
	}

	found := srv.depRepo.Find(1)
	if !found {
		t.Error("Expected to find the element")
	}

	found = srv.depRepo.Find(0)  // find not existing
	if found {
		t.Error("Found not existing item")
	}
}

func TestDependencyService_RemoveNotExisting(t *testing.T) {
	srv := NewDependencyService(createInMemoryDepRepo(), createInMemorySchRepo())

	// Try adding non existing dependent item.
	err := srv.AddDependency(0, 1, 2 ,3)
	if err == nil {
		t.Error("Added not existing dependentId")
	}

	// Add a valid task.Scheduable and try adding invalid depdentOns
	srv.taskRepo.Add(task.NewContainer(0, time.Now(), time.Now()))

	// Create dependency between first two and last.
	err = srv.AddDependency(0, 1, 2)
	if err == nil {
		t.Error("Added non existing dependentOnIds")
	}

	deleted := srv.RemoveDependency(0)
	if deleted {
		t.Error("Deleted non existing item")
	}
}

func TestDependencyService_GetAllDependencies(t *testing.T) {
	srv := NewDependencyService(createInMemoryDepRepo(), createInMemorySchRepo())

	oldTime := time.Now()
	newTime := time.Now().Add(10)
	tsk := task.NewContainer(0, oldTime, oldTime)
	// Add a few tasks
	srv.taskRepo.Add(tsk)
	srv.taskRepo.Add(task.NewContainer(1, time.Now(), time.Now()))
	srv.taskRepo.Add(task.NewContainer(2, time.Now(), time.Now()))

	// Create dependency between first two and last.
	_ = srv.AddDependency(0, 1, 2)

	// Modify the original one
	// Get the copy of dependencies.
	// Expect the one inside the copy to not be changed
	tsk.SetStartTime(newTime)
	deps := srv.GetAllDependencies()
	if deps[0].Dependent().GetStartTime() == tsk.GetStartTime() {
		t.Error("Expected Start Time of the task not to change.")
	}
}
