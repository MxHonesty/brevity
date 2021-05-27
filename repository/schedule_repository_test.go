package repository

import (
	"brevity/task"
	"testing"
	"time"
)

// Test for adding task.Scheduable to the list.
func TestAddScheduleRepository(t *testing.T) {
	repo := NewScheduleRepository()
	now1 := time.Now()
	now2 := now1.Add(1)
	now3 := now2.Add(1)
	a := task.NewContainer(0, now1, now1)
	b := task.NewContainer(1, now2, now2)
	c := task.NewContainer(2, now3, now3)

	repo.Add(a)
	repo.Add(b)
	repo.Add(c)

	current := uint64(0)

	vec := repo.GetAll()
	it := vec.Iterator()
	for it.Next() {
		elem := it.Value().(*task.Container)
		if elem.GetId() != current {
			t.Errorf("Expected id %d, got %d", current, elem.GetId())
		}
		current++
	}
}

// Test case for removing all the items in ScheduleRepository.
func TestRemoveAllScheduleRepository(t *testing.T) {
	repo := NewScheduleRepository()
	now1 := time.Now()
	now2 := now1.Add(1)
	now3 := now2.Add(1)
	a := task.NewContainer(0, now1, now1)
	b := task.NewContainer(1, now2, now2)
	c := task.NewContainer(2, now3, now3)

	repo.Add(a)
	repo.Add(b)
	repo.Add(c)
	repo.RemoveAll()

	if repo.Find(1) || repo.Find(2) || repo.Find(3) {
		t.Errorf("Found item, but no items in repo")
	}

	vec := repo.GetAll()
	if vec.Size() != 0 {
		t.Errorf("Expected size 0, got %d", vec.Size())
	}
}

// Test case for removing an item from the ScheduleRepository.
func TestRemoveScheduleRepository(t *testing.T) {
	repo := NewScheduleRepository()
	now1 := time.Now()
	now2 := now1.Add(1)
	now3 := now2.Add(1)
	a := task.NewContainer(0, now1, now1)
	b := task.NewContainer(1, now2, now2)
	c := task.NewContainer(2, now3, now3)

	repo.Add(a)
	repo.Add(b)
	repo.Add(c)

	err := repo.Remove(0)
	if err != nil {
		t.Error("Expected nil, got error")
	}

	if repo.Size() != 2 {
		t.Errorf("Expected size 2, got %d", repo.Size())
	}

	found := repo.Find(0)
	if found {
		t.Error("Found a deleted item.")
	}

	found = repo.Find(1)
	if !found{
		t.Error("Did not find existing item.")
	}

	found = repo.Find(2)
	if !found{
		t.Error("Did not find existing item.")
	}
}

func TestRetrieveScheduleRepository(t *testing.T) {
	repo := NewScheduleRepository()
	a := task.NewContainer(0, time.Now(), time.Now())
	b := task.NewContainer(1, time.Now(), time.Now())
	c := task.NewContainer(2, time.Now(), time.Now())

	repo.Add(a)
	repo.Add(b)
	repo.Add(c)

	element, err := repo.Retrieve(0)
	if err != nil {
		t.Error("Element not found")
	} else {
		if element.GetId() != 0 {
			t.Errorf("id expected 0, got %d", element.GetId())
		}
	}

	element, err = repo.Retrieve( 500)
	if err == nil {
		t.Error("Expected error, got nil")
	}
}

// Test case for removing an item that doesn't
// exist from ScheduleRepository.
func TestRemoveNonExistingScheduleRepository(t *testing.T) {
	repo := NewScheduleRepository()

	err := repo.Remove(10)
	if err == nil {
		t.Error("Expected error, got nil")
	}

	repo.Add(task.NewContainer(0, time.Now(), time.Now()))
	err = repo.Remove(0)
	if err != nil {
		t.Error("Expected nil, got error")
	}

	err = repo.Remove(0)
	if err == nil {
		t.Error("Expected error, got nil")
	}
}

// Test case for making sure that the added item is
// Added by value and not by reference.
func TestAddByValueScheduleRepository(t *testing.T) {
	repo := NewScheduleRepository()
	a := task.NewContainer(0, time.Now(), time.Now())

	repo.Add(a)
	repo.Add(task.NewContainer(1, time.Now(), time.Now()))
	repo.Add(task.NewContainer(2, time.Now(), time.Now()))

	tsk := task.NewTask("A", "A")
	a.SetTasks([]task.Task{tsk})

	if a.GetTasks()[0] != tsk {
		t.Error("Task not assigned successfully")
	}

	// We updated a and we expect that the copy
	// of a we inserted will not be updated as well.

	el, _ := repo.Retrieve(0)
	if len(el.GetTasks()) != 0 {
		t.Error("The value inside the repo was updated.")
	}
}
