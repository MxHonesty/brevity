package task

import (
	"testing"
	"time"
)

// De testat
// Copy
// Getters
// Setters
// Get Tasks
// SetTasks

// Test case for the Copy interface operation.
func TestCopyContainer(t *testing.T) {
	tm := time.Now()
	c := NewContainer(0, tm, tm)
	d := c.Copy()

	if c.GetId() != d.GetId() {
		t.Error("Different ids")
	}

	if c.GetEndTime() != d.GetEndTime() {
		t.Error("Different end times")
	}

	if c.GetStartTime() != d.GetStartTime() {
		t.Error("Different start times")
	}

	// Adding a task in c to make sure they don't share
	// the same slice.
	tsk := NewTask("A", "A")
	c.SetTasks([]Task{tsk})

	if len(d.GetTasks()) != 0 {
		t.Errorf("Expected 0 tasks, got %d", len(d.GetTasks()))
	}
}

func TestSettersContainer(t *testing.T) {
	tm := time.Now()
	newTm := tm.Add(10)
	a := NewContainer(0, tm, tm)

	a.SetEndTime(newTm)
	if a.GetDuration() != newTm.Sub(tm) {
		t.Errorf("Expected duration %d, got %d", newTm.Sub(tm), a.GetDuration())
	}

	a.SetStartTime(newTm)  // Set the start time equal to end time
	if a.GetDuration() != 0 {
		t.Errorf("Expected duration 0, got %d", a.GetDuration())
	}
}

func TestTasksContainer(t *testing.T) {
	a := NewContainer(0, time.Now(), time.Now())
	tsk1 := NewTask("A", "A")
	tsk2 := NewTask("B", "B")
	tsk3 := NewTask("C", "C")


	slc := []Task{tsk1, tsk2}
	a.SetTasks(slc)
	slc = append(slc, tsk3)

	vec := a.GetTasks()
	if len(vec) != 2 {
		t.Errorf("Expected 2, got %d", len(vec))
	}
}
