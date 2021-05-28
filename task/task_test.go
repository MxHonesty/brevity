package task

import (
	"testing"
)

// Test case for creating a Task.
func TestCreateTask(t *testing.T) {
	task := NewTask("title", "description")
	if task.GetTitle() != "title" {
		t.Errorf("Expected 'title', got %s", task.GetTitle())
	}
	if task.GetDescription() != "description" {
		t.Errorf("Expected 'description', got %s", task.GetDescription())
	}
 }

 func TestSettersTask(t *testing.T) {
 	task := NewTask("title", "description")
 	task.SetTitle("A")
 	task.SetDescription("A")

 	if task.GetDescription() != "A" {
 		t.Errorf("Expected 'A', got %s", task.GetDescription())
	}

	if task.GetTitle() != "A" {
		t.Errorf("Expected 'A', got %s", task.GetTitle())
	}
 }
