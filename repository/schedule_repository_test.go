package repository

import (
	"brevity/task"
	"testing"
	"time"
)

// Test for adding Scheduables to the list.
func TestAddScheduleRepository(t *testing.T) {
	repo := NewScheduleRepository()
	a := task.NewTask(1, "A", "A", time.Now(), time.Now())
	b := task.NewTask(2, "B", "B", time.Now(), time.Now())
	c := task.NewTask(3, "C", "C", time.Now(), time.Now())

	repo.Add(&a)
	repo.Add(&b)
	repo.Add(&c)

	vec := repo.GetAll()
	if len(vec) != 3 {
		t.Errorf("Expected size of 3, got %d", len(vec))
	}

	if !repo.Find(1) {
		t.Errorf("Element with id 1 is not found")
	}

	if !repo.Find(2) {
		t.Errorf("Element with id 2 is not found")
	}

	if !repo.Find(3) {
		t.Errorf("Element with id 3 is not found")
	}
}

func TestRemoveAllScheduleRepository(t *testing.T) {
	repo := NewScheduleRepository()
	a := task.NewTask(1, "A", "A", time.Now(), time.Now())
	b := task.NewTask(2, "B", "B", time.Now(), time.Now())
	c := task.NewTask(3, "C", "C", time.Now(), time.Now())

	repo.Add(&a)
	repo.Add(&b)
	repo.Add(&c)
	repo.RemoveAll()

	if repo.Find(1) || repo.Find(2) || repo.Find(3) {
		t.Errorf("Found item, but no items in repo")
	}

	vec := repo.GetAll()
	if len(vec) != 0 {
		t.Errorf("Expected size 0, got %d", len(vec))
	}
}

func TestRetrieveScheduleRepository(t *testing.T) {
	repo := NewScheduleRepository()
	a := task.NewTask(1, "A", "A", time.Now(), time.Now())
	b := task.NewTask(2, "B", "B", time.Now(), time.Now())
	c := task.NewTask(3, "C", "C", time.Now(), time.Now())

	repo.Add(&a)
	repo.Add(&b)
	repo.Add(&c)

	_, found := repo.Retrieve(1)
	if found != nil {
		t.Errorf("Element not found.")
	}
}
