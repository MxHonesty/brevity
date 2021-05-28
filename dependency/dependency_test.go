package dependency

import (
	"brevity/task"
	"testing"
)

// Test case for making sure the tasks are copied.
func TestCreateDependency(t *testing.T) {
	tsk1 := task.NewTask("A", "A")
	tsk2 := task.NewTask("B", "B")
	tsk3 := task.NewTask("C", "C")

	dep := NewDependency([]task.Task{tsk1, tsk2}, tsk3)

	// Test that the Tasks inside the struct are copies of the
	// originals.
	tsk3.SetDescription("AAA")

	depTask := dep.Dependent()
	if depTask.GetDescription() != "C" {
		t.Errorf("Expected description C, got %s", depTask.GetDescription())
	}

	// Testing the same for an element of the slice.
	tsk1.SetDescription("AAA")  // Set description of tsk1 to AAA
	// We expect the task inside the dep not to change.
	depOnSlice := dep.DependentOn()
	if depOnSlice[0].GetDescription() != "A" {
		t.Errorf("Expected A, got %s", depOnSlice[0].GetDescription())
	}
}

// Test case for making sure the setters
// are making copies of the tasks.
func TestSettersDepOnDependency(t *testing.T) {
	dep := NewDependency([]task.Task{task.NewTask("A", "A")},
	task.NewTask("C", "C"))

	newSlice := []task.Task{task.NewTask("AA", "AA"),
		task.NewTask("BB", "BB")}

	dep.SetDependentOn(newSlice)

	// Modify newSlice, we expect the slice inside the struct not to change.
	newSlice[0].SetDescription("BAB")

	depSlice := dep.DependentOn()
	if depSlice[0].GetDescription() != "AA" {
		t.Errorf("Expected 'AA, got %s", depSlice[0].GetDescription())
	}
}

// Test case for making sure that the dependent
// setter makes a copy of the task.Task
func TestSetterDepDependency(t *testing.T) {
	tsk := task.NewTask("A", "A")
	dep := NewDependency(nil, tsk)

	tsk.SetDescription("AAA")
	depTsk := dep.Dependent()

	if depTsk.GetDescription() != "A" {
		t.Errorf("Expected 'A, got %s", depTsk.GetDescription())
	}

	// Set new tsk.
	dep.SetDependent(tsk)  // Expect to have description AAA
	tsk.SetDescription("BBB")
	depTsk = dep.Dependent()

	if depTsk.GetDescription() != "AAA" {
		t.Errorf("Expected 'AAA', got %s", depTsk.GetDescription())
	}
}
