package task

import (
	"testing"
)

// Test case for creating a Task.
func TestCreateTask(t *testing.T) {
	task := NewTask("titlu", "descriere")
	if task.GetTitle() != "titlu" {
		t.Errorf("Expected titlu, got %s", task.GetTitle())
	}
	if task.GetDescription() != "descriere" {
		t.Errorf("Expected descriere, got %s", task.GetDescription())
	}
 }
