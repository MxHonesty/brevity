package task

import (
	"testing"
	"time"
)

// Test case for creating a Task.
func TestCreateTask(t *testing.T) {
	tm := time.Now()  // Test time
	task := NewTask(1, "titlu", "descriere", tm, tm)
	if task.GetTitle() != "titlu" {
		t.Errorf("Expected titlu, got %s", task.GetTitle())
	}
	if task.GetDescription() != "descriere" {
		t.Errorf("Expected descriere, got %s", task.GetDescription())
	}
	if task.GetStartTime() != tm || task.GetEndTime() != tm {
		t.Error("Unexpected time")
	}
	if task.GetId() != 1 {
		t.Errorf("Expected 1, got %d", task.GetId())
	}
	if task.GetDuration() != 0 {
		t.Errorf("Expected 0, got %d", task.GetDuration())
	}
 }
